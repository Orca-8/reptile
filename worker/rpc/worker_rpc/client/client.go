package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"worker/engine"
	"worker/rpc/worker_rpc/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewWorkerClient(conn)
	request := engine.SerializationRequest{
		Url:        "https://www.17k.com/all",
		ParserFunc: "ParseBookKindList",
	}
	marshal, err := json.Marshal(request)
	fmt.Println(string(marshal))
	if err != nil {
		panic(err)
	}
	worker, err := client.Worker(context.Background(), &proto.RequestInfo{
		Request: string(marshal),
	})
	if err != nil {
		panic(err)
	}
	var resp engine.SerializationParseResult
	err = json.Unmarshal([]byte(worker.Response), &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(resp.Requests))
}
