package service

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

	return jsonrpc.NewClient(conn), nil
}
