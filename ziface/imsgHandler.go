package ziface

type IMsgHandler interface {
	//执行handler
	DoMsgHandler(request IRequest)

	//添加router
	AddRouter(msgID uint32, router IRouter)

	//初始化任务池
	InitTaskPool()

	//消息放到任务队列中(采用轮训的方式)
	PushToTaskQueue(request IRequest)

	GetTaskPoolSize() uint32
}
