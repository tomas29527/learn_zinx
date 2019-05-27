package znet

import (
	"errors"
	"fmt"
	"io"
	"learn_zinx/ziface"
	"net"
)

type Connection struct {
	//服务对象
	Server ziface.IServer
	//连接
	Conn *net.TCPConn
	//连接id
	Connid uint32
	//是否关闭
	isClosed bool
	//处理数据函数
	Handlers ziface.IMsgHandler
	//退出的channle
	ExitChan chan bool
	//无缓冲d管道，用于读、写Goroutine之间的消息通信
	msgChan chan []byte
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
		_, err := io.ReadFull(c.Conn, headbuf)
		if err != nil {
			fmt.Println("读取消息头信息出错:", err)
			return
		}
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

		//判断是否初始化任务池
		if c.Handlers.GetTaskPoolSize() > 0 {
			c.Handlers.PushToTaskQueue(r)
		} else {
			//调用当前链接业务(这里执行的是当前conn的绑定的handle方法)
			go c.Handlers.DoMsgHandler(r)
		}

	}
}

//写数据
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("[conn Writer exit!]", c.Conn.RemoteAddr().String())
	//不断的阻塞的等待channel的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err)
				return
			}
		case <-c.ExitChan:
			//代表Reader已经退出，此时Writer也要推出
			return
		}
	}
}

//启动链接
func (c *Connection) Start() {
	fmt.Printf("connc is read connid:%d\n", c.Connid)
	//读数据
	go c.readData()
	//写数据
	go c.StartWriter()
}

//停止链接
func (c *Connection) Stop() {
	fmt.Printf("connc is stop connid:%d", c.Connid)
	if c.isClosed == true {
		return
	}

	c.isClosed = true
	//关闭socket链接
	c.Conn.Close()
	//告知Writer关闭
	c.ExitChan <- true

	//把当前链接移除
	c.Server.GetConnManager().Remover(c)

	close(c.ExitChan)
	close(c.msgChan)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	//将data进行封包 MsgDataLen|MsgID|Data
	dp := NewDataPack()

	//MsgDataLen|MsgID|Data
	binaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg")
	}
	c.msgChan <- binaryMsg
	return nil
}

func NewConnection(conn *net.TCPConn, cid uint32, handlers ziface.IMsgHandler, server ziface.IServer) (c *Connection) {
	c = &Connection{
		Conn:     conn,
		Connid:   cid,
		isClosed: false,
		Handlers: handlers,
		ExitChan: make(chan bool, 1),
		msgChan:  make(chan []byte),
		Server:   server,
	}
	//把链接纳入链接管理模块
	c.Server.GetConnManager().Add(c)

	return
}
