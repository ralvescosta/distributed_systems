package interfaces

import (
	"context"
	"webapi/pkg/infra/grpc_clients/proto"
)

type IIventoryClient interface {
	GetProductById(ctx context.Context, id string) (*proto.ProductResponse, error)
}
