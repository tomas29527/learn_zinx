package ziface

type IConnManager interface {
	//添加链接
	Add(connection IConnection)
	//移除链接
	Remover(connection IConnection)
	//根据链接id，获取当前链接
	Get(connId uint32) (IConnection, error)
	//获取总连接数
	Len() int
	//清除链接
	ClearConn()

	//向其他链接发送下线消息
	SentMsgToOtherDown()
}
