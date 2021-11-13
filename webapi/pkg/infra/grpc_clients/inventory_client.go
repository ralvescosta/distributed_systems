package clients

import (
	"context"
	"fmt"
	"os"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/grpc_clients/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type inventoryClient struct {
	logger    interfaces.ILogger
	telemetry interfaces.ITelemetry
}

func (pst inventoryClient) GetProductById(ctx context.Context, id string) (*proto.ProductResponse, error) {
	conn, err := grpc.DialContext(ctx, os.Getenv("INVENTORY_MS_URI"), []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		panic("inventory client is down")
	}
	defer conn.Close()

	span, spanCtx := pst.telemetry.InstrumentGRPCClient(ctx, "Inventory Client")
	defer span.Finish()

	headersIn, _ := metadata.FromIncomingContext(spanCtx)
	fmt.Printf("\nheadersIn: %s", headersIn)

	client := proto.NewInventoryClient(conn)

	return client.GetProductById(spanCtx, &proto.GetByIdRequest{
		Id: id,
	})
}

func NewInventoryClient(logger interfaces.ILogger, telemetry interfaces.ITelemetry) interfaces.IIventoryClient {
	return inventoryClient{
		logger,
		telemetry,
	}
}
