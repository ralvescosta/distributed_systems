package clients

import (
	"context"
	"os"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/grpc_clients/proto"
	"webapi/pkg/infra/telemetry"

	"google.golang.org/grpc"
)

type inventoryClient struct {
	logger    interfaces.ILogger
	telemetry telemetry.ITelemetry
}

func (pst inventoryClient) GetProductById(ctx context.Context, id string) (*proto.ProductResponse, error) {
	gRPCConfigs := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.DialContext(ctx, os.Getenv("INVENTORY_MS_URI"), gRPCConfigs...)
	if err != nil {
		panic("inventory client is down")
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
	}

	return result, err
}

func NewInventoryClient(logger interfaces.ILogger, telemetry telemetry.ITelemetry) interfaces.IIventoryClient {
	return inventoryClient{
		logger,
		telemetry,
	}
}
