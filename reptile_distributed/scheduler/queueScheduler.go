package scheduler

import (
	"fmt"
	"reptile_distributed/rpc/proto"
	"reptile_distributed/wrap"
)

type QueueScheduler struct {
	RequestChan chan wrap.SerializationRequest
	WorkerPool  []*proto.WorkerClient
	WorkerChan  chan wrap.Trans
}

func (q *QueueScheduler) Submit(request wrap.SerializationRequest) {
	go func() { q.RequestChan <- request }()
}

func (q *QueueScheduler) Ready(client *proto.WorkerClient) {
	q.WorkerPool = append(q.WorkerPool, client)
}

func (q *QueueScheduler) GetWorkerChan() chan wrap.Trans {
	q.WorkerChan = make(chan wrap.Trans)
	return q.WorkerChan
}

func (q *QueueScheduler) Run() {
	q.RequestChan = make(chan wrap.SerializationRequest)
	go func() {
		for {
			// todo delete
			fmt.Println("worker num:", len(q.WorkerPool))
			for _, client := range q.WorkerPool {
				request := <-q.RequestChan
				q.WorkerChan <- wrap.Trans{
					Client:  client,
					Request: request,
				}
			}
		}
	}()
}
