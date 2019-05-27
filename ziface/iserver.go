package ziface

//定义服务器接口
type IServer interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启业务服务方法
	Serve()
	//添加路由
	AddRouter(msgID uint32, router IRouter)
	//获取链接管理对象
	GetConnManager() IConnManager
}
