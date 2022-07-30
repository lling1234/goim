package nsq

import (
	"errors"

	"strconv"
	"strings"
	"sync"

	"log"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/dsd"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/inter"

	"github.com/nsqio/go-nsq"
)

var (
	once sync.Once
	ins  *MessageProducer
)

type MessageProducer struct {
	producer []*nsq.Producer
	index    int
}

func NewMessageProducer() (*MessageProducer, error) {
	var mp = MessageProducer{
		producer: make([]*nsq.Producer, dsd.ProducerNum),
		index:    -1,
	}

	for i := range mp.producer {
		config := nsq.NewConfig()
		p, err := nsq.NewProducer(dsd.NsqdNodeAddr, config)
		if err != nil {
			return nil, err
		}

		mp.producer[i] = p
	}

	return &mp, nil
}

func ProducerFactory() (inter.Producer, error) {
	var err error
	once.Do(func() {
		ins, err = NewMessageProducer()
	})

	return ins, err
}

func (mp *MessageProducer) Publish(chn int, topic string, body []byte) error {
	if len(strings.TrimSpace(topic)) == 0 {
		return errors.New("topic is null")
	}
	if chn < 0 {
		chn = -1
	}
	// 临时主题 (消息不做缓存)
	topic += strconv.Itoa(chn) + "#ephemeral"

	mp.index++
	mp.index %= len(mp.producer)

	err := mp.producer[mp.index].Publish(topic, body)
	if err != nil {
		log.Fatalf("nsq publish msg failed: ", err.Error())
	}

	return err
}

func (mp *MessageProducer) Stop() {
	for k := range mp.producer {
		if mp.producer[k] != nil {
			mp.producer[k].Stop()
		}
	}
}
