package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"strconv"
	"time"
)

const (
	topic            = "greetings"
	bootstrapServer1 = "localhost:9092"
	bootstrapServer2 = "localhost:9093"
	bootstrapServer3 = "localhost:9094"
)

func produce(ctx context.Context) {
	i := 0

	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "greetings",
		RequiredAcks: -1,
		MaxAttempts:  1,
	}

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("hello-" + strconv.Itoa(i) + os.Args[1]),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}

}

func main() {
	ctx := context.Background()
	produce(ctx)
}
