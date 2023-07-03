package app

import "github.com/Shopify/sarama"

type Context struct {
	Consumer sarama.Consumer

	TopicHandlers []ITopicHandler
}

func (ctx *Context) RegisterTopicHandler(topicHandler ITopicHandler) {
	ctx.TopicHandlers = append(ctx.TopicHandlers, topicHandler)
}
