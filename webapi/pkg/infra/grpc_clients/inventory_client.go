package clients

import (
	"context"
	"os"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/grpc_clients/proto"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
)

type inventoryClient struct {
	logger    interfaces.ILogger
	telemetry interfaces.ITelemetry
}

func (pst inventoryClient) GetProductById(ctx context.Context, id string) (*proto.ProductResponse, error) {
	gRPCConfigs := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(pst.telemetry.GetTracer())),
	}
	conn, err := grpc.DialContext(ctx, os.Getenv("INVENTORY_MS_URI"), gRPCConfigs...)
	if err != nil {
		panic("inventory client is down")
	}
	defer conn.Close()

	// span, spanCtx := pst.telemetry.InstrumentGRPCClient(ctx, "Inventory Client")
	// defer span.Finish()

	client := proto.NewInventoryClient(conn)

	return client.GetProductById(ctx, &proto.GetByIdRequest{
		Id: id,
	})
}

func NewInventoryClient(logger interfaces.ILogger, telemetry interfaces.ITelemetry) interfaces.IIventoryClient {
	return inventoryClient{
		logger,
		telemetry,
	}
}
