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

func main() {

	newS := znet.NewServer("[zinx V0.1]")
	newS.AddRouter(123, new(HelloRouter))
	newS.AddRouter(124, new(PingRouter))
	newS.Serve()
}
