package inter

// 生产者接口
type Producer interface {
	Publish(int, string, []byte) error
	Stop()
}

type ProducerFactory func() (Producer, error)
