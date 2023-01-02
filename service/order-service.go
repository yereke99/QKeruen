package service

import (
	"fmt"
	"qkeruen/dto"
	"qkeruen/repository"
)

type OrderService interface {
	CreateOrder(userId int, order dto.OrderRequest) error
	GetOrders(id int) ([]*dto.OrderResponse, error)
	GetMyOrders(id int) ([]*dto.OrderResponse, error)
	DeleteOrder(orderId int) error
}

type orderService struct {
	db repository.OrderDB
}

func NewOrderService(order repository.OrderDB) *orderService {
	return &orderService{db: order}
}

func (s *orderService) CreateOrder(userId int, order dto.OrderRequest) error {
	return s.db.CreateOrder(userId, order)
}

func (s *orderService) GetOrders(driverId int) ([]*dto.OrderResponse, error) {
	res, err := s.db.GetOrders(driverId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *orderService) GetMyOrders(id int) ([]*dto.OrderResponse, error) {
	res, err := s.db.GetMyOrders(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *orderService) DeleteOrder(orderId int) error {
	return s.db.DeleteOrder(orderId)
}
