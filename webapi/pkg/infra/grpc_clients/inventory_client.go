package clients

import (
	"context"
	"fmt"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/grpc_clients/proto"

	"google.golang.org/grpc"
)

type inventoryClient struct {
	conn   grpc.ClientConn
	logger interfaces.ILogger
}

func (pst inventoryClient) GetInventoryById() {
	client := proto.NewInventoryClient(&pst.conn)

	response, err := client.GetProductById(context.Background(), &proto.GetByIdRequest{
		Id: "1",
	})

	pst.logger.Debug("[InventoryClient::GetInventoryById]")
	pst.logger.Debug(fmt.Sprintf("%v", response))
	pst.logger.Debug(err.Error())
}

func NewInventoryClient(conn grpc.ClientConn, logger interfaces.ILogger) interfaces.IIventoryClient {
	return inventoryClient{
		conn,
		logger,
	}
}
