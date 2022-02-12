package clients

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"reptile_distributed/rpc/proto"
	"time"
)

func GetSaverClient() (*proto.SaverClient, error) {
	conn, err := grpc.Dial("127.0.0.1:20000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(keepalive.ClientParameters{Timeout: time.Second * 9999}))
	if err != nil {
		return nil, err
	}
	// defer conn.Close()
	client := proto.NewSaverClient(conn)
	return &client, nil
}

func Save(client proto.SaverClient, item string) error {
	_, err := client.Save(context.Background(), &proto.Request{Item: item})
	if err != nil {
		return err
	}
	return nil
}
