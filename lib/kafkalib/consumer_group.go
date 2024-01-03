package kafkalib

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func AsyncRunConsumerGroup(brokers []string, topic, group string, consumeFunc func(msg kafka.Message) error) {
	go func() {
		config := kafka.ConsumerGroupConfig{
			Brokers: brokers,
			Topics:  []string{topic},
		}
		consumerGroup, err := kafka.NewConsumerGroup(config)
		if err != nil {
			panic(err)
		}
		defer consumerGroup.Close()

		for {
			gen, err := consumerGroup.Next(context.TODO())
			if err != nil {
				panic(err)
			}
			assignments := gen.Assignments[topic]
			for range assignments {
				reader := kafka.NewReader(kafka.ReaderConfig{
					Brokers: brokers,
					Topic:   topic,
					GroupID: group,
				})
				defer reader.Close()

				gen.Start(func(ctx context.Context) {
					for {
						msg, err := reader.FetchMessage(ctx)
						if err != nil {
							return
						}

						if err = consumeFunc(msg); err != nil {
							return
						}

						err = reader.CommitMessages(ctx, msg)
						if err != nil {
							return
						}
					}
				})
			}
		}
	}()
}
