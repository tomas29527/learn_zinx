package ziface

type IDataPack interface {
	//装包
	Pack(m IMessage) ([]byte, error)
	//拆包
	UnPack(buf []byte) (IMessage, error)
}
