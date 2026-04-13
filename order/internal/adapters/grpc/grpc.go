package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/airlangga-hub/ecommerce-microservices-proto/golang/order"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/application/domain"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int64
	order.UnimplementedOrderServiceServer
}

func NewAdapter(api ports.APIPort, port int64) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	order.RegisterOrderServiceServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem

	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)

	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
