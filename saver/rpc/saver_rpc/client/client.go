package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"saver/model"
	"saver/rpc/saver_rpc/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:20000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewSaverClient(conn)
	var book model.Book
	book.Name = "java进阶"
	book.Url = "https://www.17k.com/book/golang.html"
	book.Author = "未知"
	book.Kind = "职业教育"
	book.Platform = "17k"
	book.Info = "java进阶"
	item, err := json.Marshal(&book)
	if err != nil {
		log.Printf("Faild to marshal object: %v, err: %s", book, err)
		return
	}
	client.Save(context.Background(), &proto.Request{Item: string(item)})
}
