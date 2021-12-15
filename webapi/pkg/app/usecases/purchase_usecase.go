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

	return pst.messageBroker.Publisher(ctx, exchange, kind, queue, routingKey, dto, nil)
}

func NewPruchaseUseCase() usecases.IPurchaseUseCase {
	return purchaseUseCase{}
}
