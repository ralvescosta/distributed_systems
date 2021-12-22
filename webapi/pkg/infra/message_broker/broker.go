package messagebroker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/telemetry"

	"github.com/streadway/amqp"
)

type messageBroker struct {
	telemetry telemetry.ITelemetry
}

func connect() (*amqp.Channel, error) {
	host := os.Getenv("AMQP_BROKER_HOST")
	port := os.Getenv("AMQP_BROKER_PORT")
	user := os.Getenv("AMQP_BROKER_USER")
	password := os.Getenv("AMQP_BROKER_PASS")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port))
	if err != nil {
		return nil, err
	}

	return conn.Channel()
}

func asserts(ch *amqp.Channel, exchangeName, queueName, exchangeType, routingKey, deadLetterExchange, deadLetterRoutingKey string) error {
	err := ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		return errors.NewInternalError("amqp assert exchange!")
	}

	_, err = ch.QueueDeclare(queueName, true, false, false, false, amqp.Table{
		"x-dead-letter-exchange":    deadLetterExchange,
		"x-dead-letter-routing-key": deadLetterRoutingKey,
	})
	if err != nil {
		return errors.NewInternalError("amqp assert queue!")
	}

	err = ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		return errors.NewInternalError("amqp bind queue!")
	}

	return nil
}

func (pst messageBroker) Publisher(
	ctx context.Context,
	exchangeName, exchangeType, queueName, routingKey, deadLetterExchange, deadLetterRoutingKey string,
	body interface{},
	headers map[string]interface{},
) error {
	span, spanCtx := pst.telemetry.InstrumentAMQPPublisher(ctx, exchangeName, queueName)
	defer span.Finish()

	if headers == nil {
		headers = make(map[string]interface{})
	}

	ch, err := connect()
	if err != nil {
		return errors.NewInternalError("amqp connection error!")
	}

	if err := asserts(ch, exchangeName, queueName, exchangeType, routingKey, deadLetterExchange, deadLetterRoutingKey); err != nil {
		return err
	}

	amqpBody, err := json.Marshal(body)
	if err != nil {
		return errors.NewInternalError("body convert")
	}

	pst.telemetry.InjectAMQPHeader(headers, spanCtx)

	err = ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        amqpBody,
		Headers:     headers,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewMessageBroker(telemetry telemetry.ITelemetry) interfaces.IMessageBroker {
	return messageBroker{telemetry}
}
