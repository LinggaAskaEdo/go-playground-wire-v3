package rabbit

import (
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	usersvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/service"
	pubsub "github.com/samber/go-amqp-pubsub"
	"github.com/samber/mo"
)

type RabbitHandlerImpl struct {
	usersvc.UserService
}

func NewRabbitHandler(
	usersvc usersvc.UserService,
) *RabbitHandlerImpl {
	return &RabbitHandlerImpl{
		UserService: usersvc,
	}
}

func (r *RabbitHandlerImpl) Rabbit(conn *pubsub.Connection) []*pubsub.Consumer {
	var consumers []*pubsub.Consumer

	if config.Get().Module.User.Pubsub.UserCreatedEnable {
		consumerUserCreated := pubsub.NewConsumer(conn, "consumer-user", pubsub.ConsumerOptions{
			Queue: pubsub.ConsumerOptionsQueue{
				Name: "user.queue",
			},
			Bindings: []pubsub.ConsumerOptionsBinding{
				{ExchangeName: "event.go.play", RoutingKey: "user.created"},
			},
			Message: pubsub.ConsumerOptionsMessage{
				PrefetchCount: mo.Some(100),
			},
			EnableDeadLetter: mo.Some(false),
		})

		r.ConsumeUserMessages(consumerUserCreated)

		consumers = append(consumers, consumerUserCreated)

		return consumers
	}

	return nil
}
