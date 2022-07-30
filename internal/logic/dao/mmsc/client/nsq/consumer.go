package nsq

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/dsd"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/inter"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/utils/uuid"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	*nsq.Consumer
	channel string
}

type MessageConsumer struct {
	topic  string           // 消费者主题
	cons   map[int]Consumer // 主题:消费者
	msgchn chan []byte      // 消息传输通道
	closed bool
	count  int
}

func (mc *MessageConsumer) NewConsumer(chn int) error {
	// 临时主题和通道 (消息不做缓存)
	uid, _ := uuid.NewV4()
	topic := mc.topic + strconv.Itoa(chn) + "#ephemeral"
	channel := uid.String() + "#ephemeral"

	config := nsq.NewConfig()
	config.LookupdPollInterval = 1 * time.Second

	cons, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}
	// LogLevel>3不打印nsq日志
	cons.SetLoggerLevel(11)

	cons.AddHandler(mc)
	err = cons.ConnectToNSQLookupd(dsd.NsqLookupdHttp)
	if err != nil {
		return err
	}

	mc.cons[chn] = Consumer{
		Consumer: cons,
		channel:  uid.String(),
	}
	return nil
}

func ConsumerFactory(topic string, chns []int) (inter.Consumer, error) {
	if len(strings.TrimSpace(topic)) == 0 {
		return nil, errors.New("topic is null")
	}

	if len(chns) == 0 {
		chns = []int{-1}
	}

	mc := MessageConsumer{
		topic:  topic,
		cons:   make(map[int]Consumer),
		msgchn: make(chan []byte, 30),
		closed: false,
		count:  0,
	}

	for _, chn := range chns {
		if err := mc.NewConsumer(chn); err != nil {
			return nil, err
		}
	}

	return &mc, nil
}

func (mc *MessageConsumer) Add(chns []int) error {
	for _, chn := range chns {
		if err := mc.NewConsumer(chn); err != nil {
			return err
		}
	}
	return nil
}

func (mc *MessageConsumer) Consume() ([]byte, bool) {
	data, ok := <-mc.msgchn
	return data, ok
}

func (mc *MessageConsumer) Cancel(chns []int) error {
	node, err := GetNsqNodes()
	if err != nil {
		return err
	}

	for _, prod := range node.Producers {
		for _, chn := range chns {
			if cons, ok := mc.cons[chn]; ok {
				cons.Stop()
				topic := mc.topic + strconv.Itoa(chn)
				DeleteChannel(prod, topic, cons.channel)
				delete(mc.cons, chn)
			}
		}
	}

	return nil
}

func (mc *MessageConsumer) Close() error {
	if mc.closed {
		return nil
	}
	node, err := GetNsqNodes()
	if err != nil {
		goto end
	}
	for _, prod := range node.Producers {
		for key, val := range mc.cons {
			mc.cons[key].Stop()
			topic := mc.topic + strconv.Itoa(key)
			DeleteChannel(prod, topic, val.channel)
		}
	}

end:
	mc.cons = nil
	mc.closed = true
	for {
		if mc.count == 0 {
			close(mc.msgchn)
			return nil
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func (mc *MessageConsumer) HandleMessage(message *nsq.Message) error {
	if !mc.closed {
		mc.count++
		mc.msgchn <- message.Body
		mc.count--
	}

	return nil
}
