package clients

import (
	"context"
	"os"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/infra/grpc_clients/proto"
	"webapi/pkg/infra/telemetry"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type inventoryClient struct {
	logger    interfaces.ILogger
	telemetry telemetry.ITelemetry
}

func (pst inventoryClient) GetProductById(ctx context.Context, id string) (dtos.ProductDto, error) {
	gRPCConfigs := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.DialContext(ctx, os.Getenv("INVENTORY_MS_URI"), gRPCConfigs...)
	if err != nil {
		return dtos.ProductDto{}, errors.NewInternalError("error whiling connect in inventory gRPC")
	}
	defer conn.Close()

	span, spanCtx := pst.telemetry.InstrumentGRPCClient(ctx, "Inventory Client")
	defer span.Finish()

	client := proto.NewInventoryClient(conn)

	result, err := client.GetProductById(spanCtx, &proto.GetByIdRequest{
		Id: id,
	})

	if err != nil {
		span.SetTag("error", true)
		err = mapErrorToHttp(err)
	}

	return toBookDto(result), err
}

func toBookDto(response *proto.ProductResponse) dtos.ProductDto {
	if response == nil {
		return dtos.ProductDto{}
	}

	return dtos.ProductDto{}
}

func mapErrorToHttp(grpcError error) error {
	errStatus, _ := status.FromError(grpcError)

	switch errStatus.Code() {
	case codes.NotFound:
		return errors.NewNotFoundError("product not found")
	default:
		return errors.NewInternalError("some error occur in grpc client request")
	}
}

func NewInventoryClient(logger interfaces.ILogger, telemetry telemetry.ITelemetry) interfaces.IIventoryClient {
	return inventoryClient{
		logger,
		telemetry,
	}
}
