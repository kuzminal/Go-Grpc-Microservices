package api

import (
	"context"
	"github.com/kuzminal/microservices/order/internal/adapters/payment"
	"github.com/kuzminal/microservices/order/internal/application/core/domain"
	"github.com/kuzminal/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment *payment.Adapter
}

func NewApplication(db ports.DBPort, paymentAdapter *payment.Adapter) *Application {
	return &Application{
		db:      db,
		payment: paymentAdapter,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.SendPayment(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}
	err = a.db.UpdateStatus("status", "Success", order.ID)
	if err != nil {
		return domain.Order{}, nil
	}
	return order, nil
}

func (a Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.db.Get(ctx, id)
}
