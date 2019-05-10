package main

import (
	"fmt"
	"learn_zinx/ziface"
	"learn_zinx/znet"
)

type Router struct {
	znet.BaseRouter
}

//前置处理
func (b *Router) PreHandle(r ziface.IRequest) {
	fmt.Println("======do PreHandle=======")
}
func (b *Router) Handle(r ziface.IRequest) {
	fmt.Println("======do PreHandle=======")
	message := znet.NewMessage(222, []byte("ni hao a"))
	pack := znet.NewDataPack()
	bytes, e := pack.Pack(message)
	if e != nil {
		fmt.Println("pack is err:", e)
	}
	r.GetConnection().GetTcpConnection().Write(bytes)

}

//后置处理
func (b *Router) PostHandle(r ziface.IRequest) {
	fmt.Println("======do PreHandle=======")
}
func main() {
	newS := znet.NewServer("[zinx V0.1]")
	newS.AddRouter(new(Router))
	newS.Serve()
}
