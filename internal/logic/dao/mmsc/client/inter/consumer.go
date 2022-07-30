package inter

// 消费者接口
type Consumer interface {
	Consume() ([]byte, bool) // 订阅消息
	Add([]int) error         // 添加指定通道
	Cancel([]int) error      // 取消指定通道
	Close() error            // 关闭所有通道
}

type ConsumerFactory func(string, []int) (Consumer, error)
