package worker

import (
	"context"
	"encoding/json"
	"log"
	"worker/17k/parser"
	"worker/engine"
	"worker/fetcher"
	"worker/rpc/worker_rpc/proto"
)

type Work struct {
}

/*func (w *Work) Worker(request engine.Request) (engine.ParseResult, error) {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher error, [url: %s], [err: %s]", request.Url, err)
		return engine.ParseResult{}, err
	}
	parseResult := request.ParserFunc(content)
	return parseResult, nil
}*/

func (w Work) Worker(ctx context.Context, in *proto.RequestInfo) (*proto.ResponseInfo, error) {
	// map存储所有解析方法
	parses := make(map[string]func([]byte) engine.SerializationParseResult)
	parses["ParseBook"] = parser.ParseBook
	parses["ParseBookKindList"] = parser.ParseBookKindList
	parses["ParseBookList"] = parser.ParseBookList
	var request engine.SerializationRequest
	err := json.Unmarshal([]byte(in.Request), &request)
	if err != nil {
		return nil, err
	}
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher error, [url: %s], [err: %s]", request.Url, err)
		return nil, err
	}
	parseResult := parses[request.ParserFunc](content)
	result, err := json.Marshal(parseResult)
	if err != nil {
		return nil, err
	}
	return &proto.ResponseInfo{
		Response: string(result),
	}, nil
}
