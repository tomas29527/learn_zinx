package znet

import (
	"fmt"
	"learn_zinx/util"
	"learn_zinx/ziface"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
	//任务队列
	TaskQueue []chan ziface.IRequest
	//任务池数量大小
	TaskPoolSize uint32
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//1 从Request中找到msgID
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is NOT FOUND! Need Register!")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgID]; ok {
		//id已经注册了
		panic("repeat api , msgID = " + strconv.Itoa(int(msgID)))
	}
	m.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, " succ!")
}

func NewMsgHandler() (handler *MsgHandler) {
	handler = &MsgHandler{
		Apis:         make(map[uint32]ziface.IRouter),
		TaskQueue:    make([]chan ziface.IRequest, util.Global.TaskPoolSize),
		TaskPoolSize: util.Global.TaskPoolSize,
	}
	return
}

//初始化任务池
func (m *MsgHandler) InitTaskPool() {

	//初始化任务队列
	//根据TaskPoolSize 分别开启TaskQueue池，每个TaskQueue用一个go来承载
	for i := 0; i < int(util.Global.TaskPoolSize); i++ {
		// 1 当前的TaskQueue对应的channel消息队列 开辟空间 第0个worker 就用第0个channel ...
		m.TaskQueue[i] = make(chan ziface.IRequest, util.Global.MaxTaskChanLen)
		//2 启动当前的TaskQueue， 阻塞等待消息从channel传递进来
		go m.initTaskChan(i, m.TaskQueue[i])
	}

}

//
func (m *MsgHandler) initTaskChan(taskId int, taskQueue chan ziface.IRequest) {
	fmt.Println("task ID = ", taskId, " is started ...")
	//不断的阻塞等待对应消息队列的消息
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

//消息放到任务队列中(采用轮训的方式)
func (m *MsgHandler) PushToTaskQueue(request ziface.IRequest) {
	//取模，决定放入哪个队列
	queueId := request.GetConnection().GetConnId() % m.TaskPoolSize

	fmt.Println("Add ConnID = ", request.GetConnection().GetConnId(),
		" reqeust MsgID = ", request.GetMsgID(),
		" to queueId = ", queueId)

	m.TaskQueue[queueId] <- request
}

func (m *MsgHandler) GetTaskPoolSize() uint32 {
	return m.TaskPoolSize
}
