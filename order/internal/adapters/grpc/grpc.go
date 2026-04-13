package grpc

import (
	"github.com/airlangga-hub/ecommerce-microservices-proto/golang/order"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/ports"
)

type Adapter struct {
	api ports.APIPort
	port int64
	order.UnimplementedOrderServiceServer
}