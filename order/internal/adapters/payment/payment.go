package payment

import (
	"github.com/airlangga-hub/ecommerce-microservices-proto/golang/payment"
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
