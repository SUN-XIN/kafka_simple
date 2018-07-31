package main

import (
	"log"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func main() {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Return.Errors = true
	config.Version = sarama.MaxVersion

	// init consumer
	brokers := []string{"localhost:9092"}
	topics := []string{"test1"}
	consumer, err := cluster.NewConsumer(brokers, "g1", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	go func(c *cluster.Consumer) {
		errors := c.Errors()
		notif := c.Notifications()
		for {
			select {
			case err := <-errors:
				log.Printf("Errors chan: %+v", err)
			case not := <-notif:
				log.Printf("Notifications chan: %+v", not)
			}
		}
	}(consumer)

	for msg := range consumer.Messages() {
		log.Printf("%s/%d/%d\t%s", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		consumer.MarkOffset(msg, "")
	}
}
