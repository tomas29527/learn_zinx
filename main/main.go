package main

import "learn_zinx/znet"

func main() {
	newS := znet.NewServer("[zinx V0.1]")
	newS.Serve()
}
