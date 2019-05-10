package znet

import (
	"fmt"
	"io"
	"learn_zinx/ziface"
	"net"
)

type Connection struct {
	//连接
	Conn *net.TCPConn
	//连接id
	Connid uint32
	//是否关闭
	isClosed bool
	//处理数据函数
	Handle ziface.IRouter
	//退出的channle
	ExitChan chan bool
}

func (c *Connection) readData() {
	fmt.Println(c.GetRemoteAddr().String(), "is begin read ")
	defer fmt.Println(c.GetRemoteAddr().String(), " conn reader exit!")
	defer c.Stop()
	var headbuf []byte
	for {
		//进行拆包 读取包头
		dp := NewDataPack()
		headbuf = make([]byte, dp.GetHeadLen())
		io.ReadFull(c.Conn, headbuf)
		//头信息拆包
		message, e := dp.UnPack(headbuf)
		if e != nil {
			fmt.Println("UnPack is err", e)
		}
		//第二次读出数据
		if message.GetMsgLen() > 0 {
			dataBuf := make([]byte, message.GetMsgLen())
			_, err := io.ReadFull(c.Conn, dataBuf)
			if err != nil {
				fmt.Println("read dataBuf is err", err)
				return
			}
			message.SetData(dataBuf)
		}
		r := NewRequest(c, message)
		fmt.Println("---> Recv MsgID: ", message.GetMsgId(), ", datalen = ", message.GetMsgLen(), "data = ", string(message.GetData()))

		//调用当前链接业务(这里执行的是当前conn的绑定的handle方法)
		go func(req ziface.IRequest) {
			c.Handle.PreHandle(req)
			c.Handle.Handle(req)
			c.Handle.PostHandle(req)
		}(r)
	}
}

//启动链接
func (c *Connection) Start() {
	fmt.Printf("connc is read connid:%d\n", c.Connid)
	//读数据
	go c.readData()
	select {
	case <-c.ExitChan:
		c.Stop()
	}
}

//停止链接
func (c *Connection) Stop() {
	fmt.Printf("connc is stop connid:%d", c.Connid)
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	c.Conn.Close()
	c.ExitChan <- true
	close(c.ExitChan)
}

//获取当前连接
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

//得到连接id
func (c *Connection) GetConnId() uint32 {
	return c.Connid
}

//客户端连接地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, cid uint32, handle ziface.IRouter) (c *Connection) {
	c = &Connection{
		Conn:     conn,
		Connid:   cid,
		isClosed: false,
		Handle:   handle,
		ExitChan: make(chan bool, 1),
	}
	return
}
