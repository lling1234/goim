package dsd

const (
	// 消息队列名称
	DefaultQueueName = "nsq"

	// NSQ Address
	NsqdNodeAddr   = "ling11.top:4150" // 生产者通讯地址
	NsqLookupdHttp = "ling11.top:4161" // 消费者通讯地址

	//生产者数目
	ProducerNum = 10
)
