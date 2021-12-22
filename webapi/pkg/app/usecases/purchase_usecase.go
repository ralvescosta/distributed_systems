package usecases

import (
	"context"
	"os"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type purchaseUseCase struct {
	messageBroker interfaces.IMessageBroker
}

func (pst purchaseUseCase) Perform(ctx context.Context, dto dtos.CreatePurchaseDto) error {
	queue := os.Getenv("AMQP_PURCHASE_QUEUE")
	exchange := os.Getenv("AMQP_PURCHASE_EXCHANGE")
	kind := os.Getenv("AMQP_PURCHASE_EXCHANGE_KIND")
	routingKey := os.Getenv("AMQP_PURCHASE_ROUTING_KEY")
	deadLetterExchange := os.Getenv("DEAD_LETTER_EXCHANGE")
	deadLetterRoutingKey := os.Getenv("DEAD_LETTER_ROUTING_KEY")

	return pst.messageBroker.Publisher(ctx, exchange, kind, queue, routingKey, deadLetterExchange, deadLetterRoutingKey, dto, nil)
}

func NewPruchaseUseCase(messageBroker interfaces.IMessageBroker) usecases.IPurchaseUseCase {
	return purchaseUseCase{messageBroker}
}
