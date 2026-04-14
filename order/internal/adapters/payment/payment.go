package payment

import (
	"context"

	"github.com/airlangga-hub/ecommerce-microservices-proto/golang/payment"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/application/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	paymentClient payment.PaymentServiceClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	paymentCC, err := grpc.NewClient(paymentServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer paymentCC.Close()

	paymentClient := payment.NewPaymentServiceClient(paymentCC)

	return &Adapter{paymentClient: paymentClient}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.paymentClient.CreatePayment(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}
