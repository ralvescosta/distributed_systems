package clients

import (
	"context"
	"errors"
	"fmt"
	"os"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/grpc_clients/proto"

	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type inventoryClient struct {
	logger    interfaces.ILogger
	telemetry interfaces.ITelemetry
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
	asdf, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		return nil, errors.New("")
	}

	ctxWithHeaders := metadata.NewOutgoingContext(
		spanCtx,
		metadata.Pairs("traceparent", fmt.Sprintf("00-%s-%s-01", asdf.ParentID(), asdf.SpanID())),
	)

	client := proto.NewInventoryClient(conn)

	result, err := client.GetProductById(ctxWithHeaders, &proto.GetByIdRequest{
		Id: id,
	})
	if err != nil {
		span.SetTag("error", true)
	}

	return result, err
}

func NewInventoryClient(logger interfaces.ILogger, telemetry interfaces.ITelemetry) interfaces.IIventoryClient {
	return inventoryClient{
		logger,
		telemetry,
	}
}
