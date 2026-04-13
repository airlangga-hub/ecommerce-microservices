package ports

import "github.com/airlangga-hub/ecommerce-microservices/order/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
