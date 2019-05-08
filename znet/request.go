package znet

import "learn_zinx/ziface"

//当前请求对象
type Request struct {
	//当前链接
	Conn ziface.IConnection
	//当前数据
	Data []byte
}

//获取当前链接对象
func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

//获取数据
func (r *Request) GetData() []byte {
	return r.Data
}

func NewRequest(conn ziface.IConnection, buf []byte) (r *Request) {
	r = &Request{
		Conn: conn,
		Data: buf,
	}
	return
}
