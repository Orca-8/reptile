package saver

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	elastic "github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"saver/rpc/saver_rpc/proto"
	"strings"
)

type Saver struct {
}

func (s *Saver) Save(ctx context.Context, in *proto.Request) (*emptypb.Empty, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get es client, err: %s\n", err))
		return &empty.Empty{}, err
	}
	uuid := uuid.NewV4().String()
	uuid = strings.ReplaceAll(uuid, "-", "")
	item := in.Item
	_, err = client.Index().Index("books").Id(uuid).BodyJson(item).Do(context.Background())
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
