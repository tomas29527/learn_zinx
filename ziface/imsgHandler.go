package ziface

type IMsgHandler interface {
	//执行handler
	DoMsgHandler(request IRequest)

	//添加router
	AddRouter(msgID uint32, router IRouter)
}
