package ports

import (
	"context"
	"github.com/kuzminal/microservices/payment/internal/application/core/domain"
)

type APIPort interface {
	SendPayment(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	GetPayment(ctx context.Context, id int64) (domain.Payment, error)
}
