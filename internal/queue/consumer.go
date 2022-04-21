package queue

import (
	"context"
	"fmt"

	"github.com/VladNF/calendar/internal/common"
	"github.com/streadway/amqp"
)

type Consumer struct {
	mqURI       string
	mqExchange  string
	mqKey       string
	mqQueue     string
	mqCon       *amqp.Connection
	mqChan      *amqp.Channel
	consumeFunc func(string)
	log         common.Logger
	done        chan interface{}
}

func NewConsumer(mqConfig common.MQConf, log common.Logger, consumeFunc func(string)) common.StartStopper {
	return &Consumer{
		mqURI:       mqConfig.URI,
		mqExchange:  mqConfig.Exchange,
		mqKey:       mqConfig.Key,
		mqQueue:     mqConfig.Queue,
		consumeFunc: consumeFunc,
		log:         log,
		done:        make(chan interface{}),
	}
}

func (c *Consumer) Start() error {
	var err error
	c.log.Infof("dialing %q", c.mqURI)
	c.mqCon, err = amqp.Dial(c.mqURI)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	c.log.Infof("got Connection, getting Channel")
	c.mqChan, err = c.mqCon.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}

	c.log.Infof("got Channel, declaring %q Exchange (%q)", "direct", c.mqExchange)
	if err := c.mqChan.ExchangeDeclare(
		c.mqExchange, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}

	c.log.Infof("declared Exchange, declaring Queue %q", c.mqQueue)
	queue, err := c.mqChan.QueueDeclare(
		c.mqQueue, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("queue declare: %w", err)
	}

	c.log.Infof("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, c.mqKey)

	if err = c.mqChan.QueueBind(
		queue.Name,   // name of the queue
		c.mqKey,      // bindingKey
		c.mqExchange, // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("queue bind: %w", err)
	}

	c.log.Info("Queue bound to Exchange, starting Consume")
	deliveries, err := c.mqChan.Consume(
		queue.Name, // name
		"",         // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("queue consume: %w", err)
	}

	go c.consume(deliveries)
	return nil
}

func (c *Consumer) Stop(ctx context.Context) error {
	// will close() the deliveries channel
	if err := c.mqChan.Cancel("", true); err != nil {
		return fmt.Errorf("consumer cancel failed: %w", err)
	}

	if err := c.mqCon.Close(); err != nil {
		return fmt.Errorf("connection close: %w", err)
	}

	defer c.log.Info("AMQP shutdown OK")

	close(c.done)
	return nil
}

func (c *Consumer) consume(deliveries <-chan amqp.Delivery) {
	for {
		select {
		case d := <-deliveries:
			c.log.Infof("got calendar alert: %q", d.Body)
			c.consumeFunc(string(d.Body))
			d.Ack(false)
		case <-c.done:
			return
		}
	}
}
