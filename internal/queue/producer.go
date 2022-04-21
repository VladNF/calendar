package queue

import (
	"context"
	"fmt"

	"github.com/VladNF/calendar/internal/common"
	"github.com/streadway/amqp"
)

type Producer struct {
	mqURI      string
	mqExchange string
	mqKey      string
	mqCon      *amqp.Connection
	mqChan     *amqp.Channel
	messages   <-chan string
	log        common.Logger
	done       chan interface{}
}

func NewProducer(mqConfig common.MQConf, log common.Logger, messages <-chan string) common.StartStopper {
	return &Producer{
		mqURI:      mqConfig.URI,
		mqExchange: mqConfig.Exchange,
		mqKey:      mqConfig.Key,
		messages:   messages,
		log:        log,
		done:       make(chan interface{}),
	}
}

func (p *Producer) Start() error {
	var err error
	p.log.Infof("dialing %q", p.mqURI)
	p.mqCon, err = amqp.Dial(p.mqURI)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	p.log.Infof("got Connection, getting Channel")
	p.mqChan, err = p.mqCon.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}

	p.log.Infof("got Channel, declaring %q Exchange (%q)", "direct", p.mqExchange)
	if err := p.mqChan.ExchangeDeclare(
		p.mqExchange, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}

	go p.publish()
	return nil
}

func (p *Producer) Stop(ctx context.Context) error {
	close(p.done)
	return p.mqCon.Close()
}

func (p *Producer) publish() {
	for {
		select {
		case m := <-p.messages:
			p.log.Infof("publishing %dB body (%q)", len(m), m)
			if err := p.mqChan.Publish(
				p.mqExchange, // publish to an mqExchange
				p.mqKey,      // routing to 0 or more queues
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					Headers:         amqp.Table{},
					ContentType:     "application/json",
					ContentEncoding: "",
					Body:            []byte(m),
					DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
					Priority:        0,              // 0-9
					// a bunch of application/implementation-specific fields
				},
			); err != nil {
				p.log.Errorf("exchange publish: %s", err)
			}
		case <-p.done:
			return
		}
	}
}
