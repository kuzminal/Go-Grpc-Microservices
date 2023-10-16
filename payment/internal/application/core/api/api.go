package api

import (
	"context"
	"github.com/kuzminal/microservices/payment/internal/application/core/domain"
	"github.com/kuzminal/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) SendPayment(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}

func (a Application) GetPayment(ctx context.Context, id int64) (domain.Payment, error) {
	payment, err := a.db.Get(ctx, id)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
