package main

import (
	"fmt"
	"learn_zinx/ziface"
	"learn_zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

//处理
func (b *HelloRouter) Handle(r ziface.IRequest) {
	fmt.Println("======do PreHandle=======")
	r.GetConnection().SendMsg(341, []byte("ni hao a"))
}

type PingRouter struct {
	znet.BaseRouter
}

//处理
func (b *PingRouter) Handle(r ziface.IRequest) {
	fmt.Println("======do PreHandle=======")
	r.GetConnection().SendMsg(342, []byte("ping..ping"))
}

func onStart(conneciton ziface.IConnection) {
	fmt.Println("on conncetion START============")
}

func onStop(conn ziface.IConnection) {
	fmt.Println("on conncetion stop============")
	//通知其他链接
	server := conn.GetTcpServer()
	//总连接数
	totalConn := server.GetConnManager().Len()
	if totalConn > 0 {
		server.GetConnManager().SentMsgToOtherDown()
	}
}

func main() {

	newS := znet.NewServer("[zinx V0.1]")

	newS.SetOnConnStart(onStart)
	newS.SetOnConnStop(onStop)

	newS.AddRouter(123, new(HelloRouter))
	newS.AddRouter(124, new(PingRouter))
	newS.Serve()
}
