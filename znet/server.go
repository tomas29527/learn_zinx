package znet

import (
	"fmt"
	"learn_zinx/util"
	"learn_zinx/ziface"
	"net"
	"time"
)

//iServer 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	//当前server的消息管理模块， 用来绑定MsgID和对应的处理业务API关系
	MsgHandler ziface.IMsgHandler
	//链接管理属性
	ConnManager ziface.IConnManager
	// 创建链接时候调用的函数
	OnConnStart func(ziface.IConnection)
	//断开链接的时候调用的函数
	OnConnStop func(ziface.IConnection)
}

//添加路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Succ!!")
}

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//先初始化工作池
	s.MsgHandler.InitTaskPool()

	//开启一个go去做服务端Linster业务
	go func() {
		//解析地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//监听服务地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		//已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		//3 启动server网络连接业务
		var cid uint32 = 1000
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnManager.Len() >= util.Global.MaxConn {
				fmt.Println("conn is too manny ")
				conn.Close()
				continue
			}
			//3.3 TODO Server.Start() 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			c := NewConnection(conn, cid, s.MsgHandler, s)
			go c.Start()
			cid++
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnManager.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

//获取链接管理对象
func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

//注册OnConnStart 钩子函数的方法
func (s *Server) SetOnConnStart(hookfunc func(conneciton ziface.IConnection)) {
	s.OnConnStart = hookfunc
}

//注册OnConnStop钩子函数的方法
func (s *Server) SetOnConnStop(hookfunc func(conneciton ziface.IConnection)) {
	s.OnConnStop = hookfunc
}

//调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(conneciton ziface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(conneciton)
	}
}

//调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(conneciton ziface.IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conneciton)
	}
}

func NewServer(name string) (newS ziface.IServer) {
	newS = &Server{
		Name:        name,
		IPVersion:   "tcp4",
		IP:          "0.0.0.0",
		Port:        7777,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
	return
}
