package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(service interface{}, host string) error {
	registerErr := rpc.Register(service)
	if registerErr != nil {
		return registerErr
	}

	listener, listenerErr := net.Listen("tcp", host)
	if listenerErr != nil {
		return listenerErr
	}
	log.Printf("Service listening %s", host)

	for {
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Printf("accept error:%v", connErr)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	log.Printf("Client connected %s", host)

	return jsonrpc.NewClient(conn), nil
}
