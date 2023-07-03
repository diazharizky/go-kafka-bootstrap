package app

import "os"

type ITopicHandler interface {
	RegisterConsumer(handlerConsumer IHandlerConsumer)
	Listen(signterm <-chan os.Signal)
}

type IHandlerConsumer interface {
	Consume(value string)
}
