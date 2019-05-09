package main

import (
	"bufio"
	"fmt"
	"io"
	"learn_zinx/ziface"
	"learn_zinx/znet"
	"os"
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
	//newS := znet.NewServer("[zinx V0.1]")
	//newS.AddRouter(new(Router))
	//newS.Serve()
	var buf []byte = make([]byte, 512)
	file, err := os.Open("E:\\ccc.txt")
	if err != nil {
		fmt.Printf("open file is err", err)
	}
	reader := bufio.NewReader(file)
	//reader.Read(buf)

	n, err := io.ReadFull(reader, buf)
	if err != nil {
		fmt.Println("read is err:", err)
	}
	fmt.Println(string(buf))
	fmt.Println("====n=:", n)
}
