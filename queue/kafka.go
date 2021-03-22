package queue

import (
	"context"
	"github.com/jdy879526487/log"
	"github.com/segmentio/kafka-go"
	"time"
)

type kafkaConsumer struct {
	r        *kafka.Reader
	recv     func(msg []byte) error
	DeadLine time.Duration
	sig      chan struct{}
}

func NewKafkaConsumer(opts *Options) *kafkaConsumer {
	c := &kafkaConsumer{}
	c.r = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  opts.brokers,
		GroupID:  opts.group,
		Topic:    opts.topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	c.DeadLine = opts.DeadLine
	c.sig = make(chan struct{})
	return c
}

func (kc *kafkaConsumer) WithConsumer(recv func(msg []byte) error) {
	kc.recv = recv
	go kc.run()
}

func (kc *kafkaConsumer) run() {
	for {
		select {
		case <-kc.sig:
			_ = kc.r.Close()
			return
		default:
			err := kc.readMsg()
			if err != nil {
				log.Errorf("Failed deal msg ")
			}
		}
	}
}

func (kc *kafkaConsumer) readMsg() error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(kc.DeadLine))
	defer cancel()
	msg, err := kc.r.FetchMessage(ctx)
	if err != nil {
		log.Errorf("failed reade msg ", err)
		return err
	}
	err = kc.recv(msg.Value)
	if err != nil {
		log.Errorf("failed deal received msg ", err)
		return err
	}
	err = kc.r.CommitMessages(ctx, msg)
	if err != nil {
		log.Errorf("FAILED commit msg ", err)
		return err
	}
	return nil
}

func (kc *kafkaConsumer) Close() {
	kc.sig <- struct{}{}
}
