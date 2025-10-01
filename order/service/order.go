package service

import "github.com/saleh-ghazimoradi/MicroMarket/order/repository"

type OrderService interface{}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}
