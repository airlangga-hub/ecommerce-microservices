package ports

import "github.com/airlangga-hub/ecommerce-microservices/order/internal/application/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}