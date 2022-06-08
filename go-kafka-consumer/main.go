package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	topic            = "greetings"
	bootstrapServer1 = "localhost:9092"
	bootstrapServer2 = "localhost:9093"
	bootstrapServer3 = "localhost:9094"
)

func consume(ctx context.Context) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{bootstrapServer1, bootstrapServer2, bootstrapServer3},
		Topic:   topic,
		GroupID: "group1",
	})

	for {
		message, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Printf(" Key: %v\n Value: %v\n Partition: %v\n Offfset: %v \n\n\n", string(message.Key), string(message.Value), message.Partition, message.Offset)

		time.Sleep(time.Second * 3)

	}

}

func main() {
	ctx := context.Background()
	consume(ctx)
}
