package znet

import (
	"errors"
	"fmt"
	"learn_zinx/ziface"
	"sync"
)

//链接管理模块
type ConnManager struct {
	connections map[uint32]ziface.IConnection
	Rwsync      sync.RWMutex
}

//添加链接
func (c *ConnManager) Add(conn ziface.IConnection) {
	c.Rwsync.Lock()
	defer c.Rwsync.Unlock()
	c.connections[conn.GetConnId()] = conn
	fmt.Println("connID = ", conn.GetConnId(), " add to ConnManager successfully: conn num = ", c.Len())
}

//移除链接
func (c *ConnManager) Remover(conn ziface.IConnection) {
	c.Rwsync.Lock()
	defer c.Rwsync.Unlock()
	delete(c.connections, conn.GetConnId())

	fmt.Println("connID = ", conn.GetConnId(), " Remover to ConnManager successfully: conn num = ", c.Len())
}

//根据链接id，获取当前链接
func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.Rwsync.RLock()
	defer c.Rwsync.RUnlock()
	if c, ok := c.connections[connId]; ok {
		return c, nil
	} else {
		return nil, errors.New("===未找到此链接")
	}
}

//获取总连接数
func (c *ConnManager) Len() int {
	return len(c.connections)
}

//清除链接
func (c *ConnManager) ClearConn() {
	c.Rwsync.Lock()
	defer c.Rwsync.Unlock()

	for connId, conn := range c.connections {
		//停止链接
		conn.Stop()
		//移除链接
		delete(c.connections, connId)
	}
	fmt.Println("Clear All connections succ!  conn num = ", c.Len())
}

//向其他链接发送下线消息
func (c *ConnManager) SentMsgToOtherDown() {
	c.Rwsync.Lock()
	defer c.Rwsync.Unlock()
	//向其他用户发送下线消息
	for _, conn := range c.connections {
		fmt.Println("=========SentMsgToOtherDown=======")
		msg := fmt.Sprintf("用户下线了")
		//停止链接
		conn.SendMsg(456, []byte(msg))
	}
}

func NewConnManager() *ConnManager {
	connManager := &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
	return connManager
}
