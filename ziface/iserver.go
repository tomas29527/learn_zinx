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
	//注册OnConnStart 钩子函数的方法
	SetOnConnStart(func(conneciton IConnection))
	//注册OnConnStop钩子函数的方法
	SetOnConnStop(func(conneciton IConnection))
	//调用OnConnStart钩子函数的方法
	CallOnConnStart(conneciton IConnection)
	//调用OnConnStop钩子函数的方法
	CallOnConnStop(conneciton IConnection)
}
