package main

import (
	"log"

	"github.com/airlangga-hub/ecommerce-microservices/order/config"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/adapters/db"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/adapters/grpc"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/adapters/payment"
	"github.com/airlangga-hub/ecommerce-microservices/order/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grcpAdapter := grpc.NewAdapter(application, int64(config.GetApplicationPort()))
	
	grcpAdapter.Run()
}
