package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	listener, err := net.Listen("tcp", "0.0.0.0:7878")
	if err != nil {
		fmt.Println("listen is err", err)
		return
	}
	go func() {
		for {
			//获取连接
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept is err", err)
				return
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					// 1第一次从conn读， 把包的head读出来
					headData := make([]byte, dp.GetHeadLen())
					n, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head is err", err)
						return
					}
					fmt.Println("============RemoteAddr", conn.RemoteAddr().String())
					fmt.Println("is read head n=:", n)
					//fmt.Println("is read head data:",string(headData))
					msg, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("UnPack is err", err)
						return
					}
					//第二次读出数据
					if msg.GetMsgLen() > 0 {
						dataBuf := make([]byte, msg.GetMsgLen())
						a, err := io.ReadFull(conn, dataBuf)
						if err != nil {
							fmt.Println("read dataBuf is err", err)
							return
						}
						fmt.Println("is dataBuf  a=:", a)
						fmt.Println("is dataBuf  =:", string(dataBuf))
						msg.SetData(dataBuf)
						//完整的一个消息已经读取完毕
						fmt.Println("---> Recv MsgID: ", msg.GetMsgId(), ", datalen = ", msg.GetMsgLen(), "data = ", string(msg.GetData()))
					}
				}

			}(conn)
		}
	}()

	//客户端连接
	//conn, err := net.Dial("tcp", "0.0.0.0:7878")
	//if err !=nil {
	//	fmt.Println("open conn is err")
	//	return
	//}
	////封包
	//dp := NewDataPack()
	//msg1 := &Message{
	//	Mid:      123,
	//	DataLen: 4,
	//	Data:    []byte{'z', 'i', 'n'},
	//}
	//bytes1, err := dp.Pack(msg1)
	//if err !=nil{
	//	fmt.Println("Pack1  is err",err)
	//}
	//msg2 := &Message{
	//	Mid:      1234,
	//	DataLen: 5,
	//	Data:    []byte{'n', 'i', 'h', 'a','o'},
	//}
	//_, er := dp.Pack(msg2)
	//if er !=nil{
	//	fmt.Println("Pack2  is err",err)
	//}
	////bytes1 =append(bytes1,bytes2...)
	//conn.Write(bytes1)
	//time.Sleep(time.Second)
	//conn.Write([]byte{'x'})
	//time.Sleep(time.Millisecond*10)
	//conn.Write(bytes2)
	//客户端阻塞
	select {}
}
