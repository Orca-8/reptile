package main

import (
	"reptile_distributed/engine"
	"reptile_distributed/scheduler"
	"reptile_distributed/wrap"
)

func main() {
	concurrent := &engine.QueueConcurrent{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		WorkersAddr: []string{
			"127.0.0.1:30000",
		},
	}
	concurrent.Run(wrap.SerializationRequest{
		Url:        "https://www.17k.com/all",
		ParserFunc: "ParseBookKindList",
	})
}
