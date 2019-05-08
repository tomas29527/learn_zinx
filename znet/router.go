package znet

import "learn_zinx/ziface"

//路由模块基类,需要怎么处理用户自行实现
type BaseRouter struct {
}

//前置处理
func (b *BaseRouter) PreHandle(r ziface.IRequest) {}
func (b *BaseRouter) Handle(r ziface.IRequest)    {}

//后置处理
func (b *BaseRouter) PostHandle(r ziface.IRequest) {}
