package topichandlers

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/diazharizky/go-kafka-bootstrap/internal/app"
	"github.com/diazharizky/go-kafka-bootstrap/internal/enum"
)

type userRegistrationsHandler struct {
	topic            string
	consumer         sarama.Consumer
	handlerConsumers []app.IHandlerConsumer
}

func NewUserRegistrationsHandler(consumer sarama.Consumer) userRegistrationsHandler {
	return userRegistrationsHandler{
		topic:    enum.TopicUserRegistrations.String(),
		consumer: consumer,
	}
}

func (handler *userRegistrationsHandler) RegisterConsumer(handlerConsumer app.IHandlerConsumer) {
	handler.handlerConsumers = append(handler.handlerConsumers, handlerConsumer)
}

func (handler userRegistrationsHandler) Listen(sigterm <-chan os.Signal) {
	partitions, err := handler.consumer.Partitions(handler.topic)
	if err != nil {
		log.Printf("Error unable to get partitions: %v\n", err)
		return
	}

	msgChan := make(chan *sarama.ConsumerMessage, 256)
	for part := range partitions {
		go handler.transportMessage(int32(part), msgChan)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("UserRegistrationsHandler is listening!")

	for {
		select {
		case <-sigterm:
			log.Println("Terminating: Via signal")
			return
		case <-ctx.Done():
			log.Println("Terminating: Context canceled")
			return
		case msg := <-msgChan:
			for _, hc := range handler.handlerConsumers {
				go hc.Consume(string(msg.Value))
			}
		}
	}
}

func (handler userRegistrationsHandler) transportMessage(
	partition int32,
	msgChan chan *sarama.ConsumerMessage,
) {
	pc, err := handler.consumer.ConsumePartition(handler.topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Error unable to create consumer for partition %d: %v\n", partition, err)
		return
	}
	defer pc.AsyncClose()

	for {
		msg := <-pc.Messages()
		msgChan <- msg
	}
}
