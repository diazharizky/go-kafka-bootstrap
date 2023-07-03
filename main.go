package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/diazharizky/go-kafka-bootstrap/internal/app"
	"github.com/diazharizky/go-kafka-bootstrap/internal/services"
	"github.com/diazharizky/go-kafka-bootstrap/internal/topichandlers"
)

var appCtx *app.Context

func main() {
	defer func() {
		if err := appCtx.Consumer.Close(); err != nil {
			log.Fatalf("Error unable to close consumer: %v\n", err)
		}
	}()

	initApp()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	for _, h := range appCtx.TopicHandlers {
		go h.Listen(sigterm)
	}

	<-sigterm
}

func initApp() {
	appCtx = &app.Context{}

	addrs := []string{"localhost:19092"}

	var err error

	appCtx.Consumer, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		log.Fatalf("Error unable to get consumer: %v\n", err)
	}

	userRegistrationsHandler := topichandlers.NewUserRegistrationsHandler(appCtx.Consumer)
	appCtx.RegisterTopicHandler(&userRegistrationsHandler)

	userRegistrationsHandler.RegisterConsumer(services.NewSendUserConfirmationEmailService())

	userProfileUpdatesHandler := topichandlers.NewUserProfileUpdatesHandler(appCtx.Consumer)
	appCtx.RegisterTopicHandler(&userProfileUpdatesHandler)

	userProfileUpdatesHandler.RegisterConsumer(services.NewNotifyFollowersOnProfileUpdatedService())
}
