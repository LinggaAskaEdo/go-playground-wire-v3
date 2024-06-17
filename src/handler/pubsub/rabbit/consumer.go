package rabbit

import (
	"github.com/rs/zerolog/log"
	pubsub "github.com/samber/go-amqp-pubsub"
)

func (s *RabbitHandlerImpl) ConsumeUserMessages(consumerUser *pubsub.Consumer) {
	// var forever chan struct{}
	stopChan := make(chan bool)

	channel := consumerUser.Consume()
	go func() {
		for msg := range channel {
			log.Info().Msgf("Consumed message [EX=%s, RK=%s] %s", msg.Exchange, msg.RoutingKey, string(msg.Body))
			msg.Ack(false)
		}
	}()

	// <-forever
	<-stopChan
}
