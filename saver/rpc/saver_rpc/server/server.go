package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"saver"
	"saver/rpc/saver_rpc/proto"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterSaverServer(server, &saver.Saver{})
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		panic(err)
	}
	for {
		listen.Accept()
		fmt.Println("saver accepted")
		go server.Serve(listen)
	}
}
