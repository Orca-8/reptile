package clients

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"reptile_distributed/rpc/proto"
	"reptile_distributed/wrap"
	"time"
)

func GetWorkerClient(addr string) (*proto.WorkerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(keepalive.ClientParameters{Timeout: time.Second * 9999}))
	if err != nil {
		return nil, err
	}
	// defer conn.Close()
	client := proto.NewWorkerClient(conn)
	return &client, nil
}

func Work(client proto.WorkerClient, request wrap.SerializationRequest) (wrap.SerializationParseResult, error) {
	marshal, err := json.Marshal(request)
	//fmt.Println(string(marshal))
	if err != nil {
		return wrap.SerializationParseResult{}, err
	}
	worker, err := client.Worker(context.Background(), &proto.RequestInfo{
		Request: string(marshal),
	})
	if err != nil {
		return wrap.SerializationParseResult{}, err
	}
	var resp wrap.SerializationParseResult
	err = json.Unmarshal([]byte(worker.Response), &resp)
	if err != nil {
		return wrap.SerializationParseResult{}, err
	}
	// fmt.Println("workerClient: ", len(resp.Requests))
	return resp, nil
}
