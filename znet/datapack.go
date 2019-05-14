package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"learn_zinx/util"
	"learn_zinx/ziface"
)

type DataPack struct {
}

//拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包的头的长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//ID uint32(4字节)+Datalen uint32（4字节）
	return 8
}

func (dp *DataPack) Pack(m ziface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, m.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, m.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, m.GetData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

/**
* 请求解码器
* <pre>
* 数据包格式
        包头=消息id + 长度=8
* +——----——+——-----——+——----——+——----——+——-----——+
*  | 消息id       |  长度        |   数据       |
* +——----——+——-----——+——----——+——----——+——-----——+
* </pre>
* 包头4字节
* 模块号2字节short
* 命令号2字节short
* 长度4字节(描述数据部分字节长度)
*
*
*/
func (dp *DataPack) UnPack(buf []byte) (ziface.IMessage, error) {

	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(buf)
	//只解压head信息，得到datalen和MsgID
	msg := &Message{}
	//读MsgID
	binary.Read(dataBuff, binary.LittleEndian, &msg.Mid)
	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//判断datalen是否已经超出了我们允许的最大包长度
	if util.Global.MaxPackageSize > 0 && msg.DataLen > util.Global.MaxPackageSize {
		return nil, errors.New("too Large msg data recv!")
	}
	return msg, nil
}
