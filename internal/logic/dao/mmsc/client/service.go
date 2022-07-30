package client

import (
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/inter"
	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/nsq"
)

// topic
const (
	IVATrack     = "IVATrack" // 智能分析 目标跟踪
	IVAAttribute = "goimtest" // 智能分析 目标属性
)

var nsqQueue *inter.Queue

// 组件自启动
func init() {
	nsqQueue = inter.NewQueue(nsq.ProducerFactory, nsq.ConsumerFactory)
}

// 若推送消息无所属通道号,chn需小于0; topic定义见dsd包
func Publish(topic string, chn int, msg interface{}, queueName ...string) error {
	return nsqQueue.Publish(chn, topic, msg)
}

// 若订阅消息无所属通道号,chns传空数组; 可同时订阅topic下多个通道
func Subscribe(topic string, chns []int, queueName ...string) (inter.Consumer, error) {
	return nsqQueue.Subscribe(topic, chns)
}

// TODO: 未实现
func Stop(queueName ...string) {
	nsqQueue.Stop()
}
