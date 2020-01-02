package main

import (
	rpcdemo "language/rpc"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoServer.Div", rpcdemo.Args{10, 3}, &result)
	log.Printf("result: %v; err: %v", result, err)

	result = 0
	err = client.Call("DemoServer.Div", rpcdemo.Args{10, 0}, &result)
	log.Printf("result: %v; err: %v", result, err)
}
