package ziface

//请求接口
type IRequest interface {
	//获取当前链接对象
	GetConnection() IConnection
	//获取数据
	GetData() []byte
}
