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
	r.GetConnection().GetTcpConnection().Write([]byte("ping... ping ... ping "))
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
