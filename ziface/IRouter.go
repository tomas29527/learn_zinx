package ziface

//路由接口
type IRouter interface {
	//前置处理
	PreHandle(r IRequest)

	Handle(r IRequest)
	//后置处理
	PostHandle(r IRequest)
}
