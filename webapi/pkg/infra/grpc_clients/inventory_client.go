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
	conn, err := connectToGrpcServer(ctx)
	if err != nil {
		return dtos.ProductDto{}, err
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

	return toProduct(result), err
}

func (pst inventoryClient) RegisterProduct(ctx context.Context, product dtos.ProductDto) (dtos.ProductDto, error) {
	conn, err := connectToGrpcServer(ctx)
	if err != nil {
		return dtos.ProductDto{}, err
	}
	defer conn.Close()

	span, spanCtx := pst.telemetry.InstrumentGRPCClient(ctx, "Inventory Client")
	defer span.Finish()

	client := proto.NewInventoryClient(conn)

	result, err := client.CreateProduct(spanCtx, toCreateProductRequest(product))

	if err != nil {
		span.SetTag("error", true)
		err = mapErrorToHttp(err)
	}

	return toProduct(result), err
}

func connectToGrpcServer(ctx context.Context) (*grpc.ClientConn, error) {
	gRPCConfigs := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.DialContext(ctx, os.Getenv("INVENTORY_MS_URI"), gRPCConfigs...)
	if err != nil {
		return nil, errors.NewInternalError("error whiling connect in inventory gRPC")
	}
	return conn, nil
}

func toCreateProductRequest(product dtos.ProductDto) *proto.CreateProductRequest {
	return &proto.CreateProductRequest{
		Tag:           product.Tag,
		Title:         product.Title,
		Subtitle:      product.Subtitle,
		Authors:       product.Authors,
		AmountInStock: int64(product.AmountInStock),
		NumPages:      int64(product.NumPages),
		Tags:          product.Tags,
	}
}

func toProduct(response *proto.ProductResponse) dtos.ProductDto {
	if response == nil {
		return dtos.ProductDto{}
	}

	return dtos.ProductDto{
		Id:              response.Id,
		ProductCategory: response.ProductCategory,
		Tag:             response.Tag,
		Title:           response.Title,
		Subtitle:        response.Subtitle,
		Authors:         response.Authors,
		AmountInStock:   int(response.AmountInStock),
		NumPages:        int(response.NumPages),
		Tags:            response.Tags,
		CreatedAt:       response.CreatedAt,
		UpdatedAt:       response.UpdatedAt,
	}
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
