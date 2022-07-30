package inter

import (
	"encoding/json"
	"log"

	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/dsd"
)

// 消息队列
type Queue struct {
	name            string
	producerFactory ProducerFactory
	consumerFactory ConsumerFactory
}

func NewQueue(producer ProducerFactory, consumer ConsumerFactory, name ...string) *Queue {
	var queueName = dsd.DefaultQueueName
	if len(name) > 0 {
		queueName = name[0]
	}
	return &Queue{
		producerFactory: producer,
		consumerFactory: consumer,
		name:            queueName,
	}
}

// 修改消息队列名
func (q *Queue) SetName(name string) {
	q.name = name
}

// 发布消息
func (q *Queue) Publish(chn int, topic string, msg interface{}) error {
	producer, err := q.producerFactory()
	if err != nil {
		log.Fatalf("get producer failed: %v", err)
		return nil
	}

	data, _ := json.Marshal(msg)
	return producer.Publish(chn, topic, data)
}

// 消费者订阅主题
func (q *Queue) Subscribe(topic string, chns []int) (Consumer, error) {
	consumer, err := q.consumerFactory(topic, chns)
	if err != nil {
		log.Fatalf("create consumer failed: %v", err)
		return nil, err
	}

	return consumer, nil
}

// 停止
func (q *Queue) Stop() {
}
