package rabbit

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	pubsub "github.com/samber/go-amqp-pubsub"
	"github.com/samber/mo"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/rabbit/consumer"
)

type RabbitImpl struct {
	PubsubConsumer *consumer.PubsubConsumerImpl
	conn           *pubsub.Connection
	producer       *pubsub.Producer
	producers      []*pubsub.Producer
	consumer       *pubsub.Consumer
	consumers      []*pubsub.Consumer
}

func NewRabbit(
	PubsubConsumer *consumer.PubsubConsumerImpl,
) *RabbitImpl {
	log.Info().Msg("Initialize PubSub - Rabbit...")

	return &RabbitImpl{
		PubsubConsumer: PubsubConsumer,
	}
}

func (p *RabbitImpl) setupConsumer(conn *pubsub.Connection) []*pubsub.Consumer {
	return p.PubsubConsumer.Handler(conn)
}

func (p *RabbitImpl) Start() {
	rabbitUser := config.Get().Queue.Rabbit.User
	rabbitPass := config.Get().Queue.Rabbit.Password
	rabbitHost := config.Get().Queue.Rabbit.Host
	rabbitPort := config.Get().Queue.Rabbit.Port

	rabbitUri := fmt.Sprintf("amqp://%s:%s@%s:%d", rabbitUser, rabbitPass, rabbitHost, rabbitPort)

	conn, _ := pubsub.NewConnection("connection-rabbit", pubsub.ConnectionOptions{
		URI:            rabbitUri,
		LazyConnection: mo.Some(true),
	})

	p.setupConsumer(conn)
	log.Info().Msg("PubSub - Rabbit starting...")
	// log.Info().Msg("XXX")
	// // time.Sleep(15 * time.Second)
	// log.Info().Msg("YYY")
	// p.Shutdown(context.Background())
	// log.Info().Msg("ZZZ")
}

func (p *RabbitImpl) Shutdown(ctx context.Context) error {
	log.Info().Msg("AAA")
	for _, val := range p.consumers {
		if err := val.Close(); err != nil {
			return err
		}
	}

	if err := p.consumer.Close(); err != nil {
		return err
	}

	for _, val := range p.producers {
		if err := val.Close(); err != nil {
			return err
		}
	}

	if err := p.producer.Close(); err != nil {
		return err
	}

	if err := p.conn.Close(); err != nil {
		return err
	}
	log.Info().Msg("BBB")
	return nil
}
