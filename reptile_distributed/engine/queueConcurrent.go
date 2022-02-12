package engine

import (
	"fmt"
	"log"
	"reptile_distributed/rpc/clients"
	"reptile_distributed/rpc/proto"
	"reptile_distributed/wrap"
)

type QueueConcurrent struct {
	Scheduler   QueueScheduler
	WorkerCount int
	WorkersAddr []string
}

type QueueScheduler interface {
	Submit(request wrap.SerializationRequest)
	Ready(client *proto.WorkerClient)
	GetWorkerChan() chan wrap.Trans
	Run()
}

var count = 0

func (q *QueueConcurrent) Run(seeds ...wrap.SerializationRequest) {
	q.CreateWorker()
	q.Scheduler.Run()
	out := make(chan wrap.SerializationParseResult)

	for _, seed := range seeds {
		q.Scheduler.Submit(seed)
	}

	go q.Work(out)
	saverClient, err := clients.GetSaverClient()
	if err != nil {
		log.Printf("Faild to get saver client, err: %s", err)
		return
	}

	for {
		select {
		case resp := <-out:
			{
				for _, request := range resp.Requests {
					q.Scheduler.Submit(request)
				}
				for _, item := range resp.Items {
					go q.Save(*saverClient, item)
					count++
					fmt.Println("小说数量：", count)
				}
			}
		}
	}

}

func (q *QueueConcurrent) CreateWorker() {
	count := 0
	for i := 0; i < q.WorkerCount; i++ {
		if count == len(q.WorkersAddr) {
			count = 0
		}
		client, err := clients.GetWorkerClient(q.WorkersAddr[count])
		if err != nil {
			log.Printf("Faild to get worker client, err: %s", err)
			continue
		}
		q.Scheduler.Ready(client)
		count++
	}
}

// Work 通过对应client的Work方法
func (q QueueConcurrent) Work(out chan wrap.SerializationParseResult) {
	workerChan := q.Scheduler.GetWorkerChan()
	for {
		tran := <-workerChan
		go func() {
			work, err := clients.Work(*tran.Client, tran.Request)
			if err != nil {
				log.Printf("Faild to call work, err: %s", err)
				return
			}
			out <- work
		}()
	}
}

// Save 将item发送到es中
func (q QueueConcurrent) Save(client proto.SaverClient, item string) {
	err := clients.Save(client, item)
	if err != nil {
		log.Printf("Faild to send to es, err: %s", err)
		return
	}
}
