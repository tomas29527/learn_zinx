package ziface

type IMessage interface {
	//获取消息ID
	GetMsgId() uint32
	//消息长度
	GetMsgLen() uint32
	// 获取数据
	GetData() []byte
	//设置消息id
	SetMsgId(msgId uint32)
	//设置消息
	SetData(buf []byte)
	//设置消息长度
	SetDataLen(len uint32)
}
