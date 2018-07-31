package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	//  config.Producer.RequiredAcks = sarama.WaitForAll
	//  config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.MaxVersion

	brokers := []string{"localhost:9092"}
	topic := "test1"

	p, err := sarama.NewSyncProducer(brokers, config)
	defer p.Close()
	if err != nil {
		log.Printf("Failed NewSyncProducer: %+v", err)
		return
	}

	var msgStr string
	var msg sarama.ProducerMessage
	for {
		msgStr = "sync: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
		msg = sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(msgStr),
		}
		if _, _, err := p.SendMessage(&msg); err != nil {
			log.Printf("Failed SendMessage: %+v", err)
			return
		}
		log.Printf("Send msg (%s) ok", msgStr)

		time.Sleep(2 * time.Second)
	}
}
