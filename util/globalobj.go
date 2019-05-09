package util

import (
	"learn_zinx/ziface"
)

type GlobalObj struct {
	//server 配置
	TcpServer ziface.IServer //当前Zinx全局的Server对象
	Host      string         //当前服务器主机监听的IP
	TcpPort   int            //当前服务器主机监听的端口号
	Name      string         //当前服务器的名称

	/*
	   Zinx配置
	*/
	Version          string //当前Zinx的版本号
	MaxConn          int    //当前服务器主机允许的最大链接数
	MaxPackageSize   uint32 //当前Zinx框架数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作Worker池的Goroutine数量
	MaxWorkerTaskLen uint32 //Zinx框架允许用户最多开辟多少个Worker(限定条件)
}

var Global *GlobalObj

func init() {
	Global = &GlobalObj{}

	Global.Name = "zinx server"
	Global.Host = "0.0.0.0"
	Global.TcpPort = 8787
	Global.MaxConn = 100
	Global.MaxPackageSize = 512
	Global.Version = "V1.0.1"
}
