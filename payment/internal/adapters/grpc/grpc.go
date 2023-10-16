package grpc

import (
	"context"
	"fmt"
	payment "github.com/kuzminal/microservices/payment/internal/adapters/grpc/generated"
	"github.com/kuzminal/microservices/payment/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {

	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.SendPayment(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to send payment. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *payment.GetPaymentRequest) (*payment.GetPaymentResponse, error) {
	result, err := a.api.GetPayment(ctx, request.PaymentId)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to get payment. %v ", err)).Err()
	}
	return &payment.GetPaymentResponse{
		UserId:     result.CustomerID,
		PaymentId:  result.ID,
		OrderId:    result.OrderId,
		TotalPrice: result.TotalPrice,
		Status:     result.Status,
	}, nil
}
