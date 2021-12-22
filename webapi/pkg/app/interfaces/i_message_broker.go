package interfaces

import "context"

type IMessageBroker interface {
	Publisher(
		ctx context.Context,
		exchangeName, exchangeType, queueName, routingKey, deadLetterExchange, deadLetterRoutingKey string,
		body interface{},
		header map[string]interface{},
	) error
}
