package ports

import (
	"context"
	"github.com/kuzminal/microservices/order/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(*domain.Order) error
	UpdateStatus(field string, value interface{}, orderId int64) error
}
