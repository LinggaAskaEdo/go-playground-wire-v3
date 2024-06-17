package consumer

import (
	pubsub "github.com/samber/go-amqp-pubsub"

	rabbithandler "github.com/linggaaskaedo/go-playground-wire-v3/src/handler/pubsub/rabbit"
)

type PubsubConsumerImpl struct {
	rabbitHandlers *rabbithandler.RabbitHandlerImpl
}

func NewPubsubConsumer(
	rabbitHandlers *rabbithandler.RabbitHandlerImpl,
) *PubsubConsumerImpl {
	return &PubsubConsumerImpl{
		rabbitHandlers: rabbitHandlers,
	}
}

func (t *PubsubConsumerImpl) Handler(conn *pubsub.Connection) []*pubsub.Consumer {
	return t.rabbitHandlers.Rabbit(conn)
}
