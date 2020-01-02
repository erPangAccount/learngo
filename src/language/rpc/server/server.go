package main

import (
	rpcdemo "language/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	registerErr := rpc.Register(rpcdemo.DemoServer{})
	if registerErr != nil {
		panic(registerErr)
	}

	listener, listenerErr := net.Listen("tcp", ":1234")
	if listenerErr != nil {
		panic(listenerErr)
	}

	for {
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Printf("accept error:%v", connErr)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
