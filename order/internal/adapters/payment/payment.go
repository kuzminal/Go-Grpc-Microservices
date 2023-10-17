package payment

import (
	"context"
	"github.com/kuzminal/microservices/order/internal/adapters/grpc/generated/payments"
	"github.com/kuzminal/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payments.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	//defer conn.Close()
	client := payments.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) SendPayment(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(),
		&payments.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice(),
		})
	return err
}
