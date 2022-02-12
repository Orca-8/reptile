package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"worker"
	"worker/rpc/worker_rpc/proto"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterWorkerServer(server, worker.Work{})
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		panic(err)
	}
	for {
		_, err := listen.Accept()
		fmt.Println("worker accepted")
		if err != nil {
			log.Printf("faild to accept, err: %s", err)
			continue
		}
		go server.Serve(listen)
	}
}
