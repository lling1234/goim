package main

import (
	"fmt"
	"testing"

	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/inter"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/nsq"

	"time"
)

func TestMain11(t *testing.T) {
	queue := inter.NewQueue(nsq.ProducerFactory, nsq.ConsumerFactory)

	consumer, _ := queue.Subscribe(client.IVAAttribute, []int{2, 3})
	go func() {
		for {
			data, ok := consumer.Consume()
			if !ok {
				return
			}

			fmt.Println("..............9999 ", string(data), ok)
		}
	}()
	datastr := "hello11"
	queue.Publish(2, client.IVAAttribute, datastr)

	time.Sleep(time.Second * 3)

	consumer.Cancel([]int{2})
	fmt.Println("111")
}
