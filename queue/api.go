package queue

import "time"

type Options struct {
	topic      string
	group      string
	brokers    []string
	maxRetry   int
	bufferSize int
	maxWorker  int
	DeadLine time.Duration
}

type IConsumer interface {
	// 注册 消息处理函数
	WithConsumer(func(msg []byte) error)
}

type IProducer interface {
	// 发送msg 的函数
	SendMsg(msg string) error
}
