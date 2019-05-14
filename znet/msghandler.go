package znet

import (
	"fmt"
	"learn_zinx/ziface"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
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
		Apis: make(map[uint32]ziface.IRouter),
	}
	return
}
