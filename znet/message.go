package znet

//消息模块
type Message struct {
	//消息id
	Mid uint32
	//消息长度
	DataLen uint32
	//消息类容
	Data []byte
}

func NewMessage(id uint32, data []byte) *Message {
	msg := &Message{
		Mid:     id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	return msg
}

func (m *Message) GetMsgId() uint32 {
	return m.Mid
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(msgId uint32) {
	m.Mid = msgId
}

func (m *Message) SetData(buf []byte) {
	m.Data = buf
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
