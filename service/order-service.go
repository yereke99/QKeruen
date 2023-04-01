package service

import (
	"fmt"
	"qkeruen/dto"
	"qkeruen/repository"
)

type OrderService interface {
	CreateOrder(order dto.OrderRequest) error
	GetOrders(id int64) ([]*dto.OrderResponse, error)
	GetMyOrders(id int64) ([]*dto.OrderResponse, error)
	DeleteOrder(orderId int64) error
}

type orderService struct {
	db repository.OrderDB
}

func NewOrderService(order repository.OrderDB) *orderService {
	return &orderService{db: order}
}

func (s *orderService) CreateOrder(order dto.OrderRequest) error {
	return s.db.CreateOrder(order)
}

func (s *orderService) GetOrders(driverId int64) ([]*dto.OrderResponse, error) {
	res, err := s.db.GetOrders(driverId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *orderService) GetMyOrders(id int64) ([]*dto.OrderResponse, error) {
	res, err := s.db.GetMyOrders(id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, err
}

func (s *orderService) DeleteOrder(orderId int64) error {
	return s.db.DeleteOrder(orderId)
}
