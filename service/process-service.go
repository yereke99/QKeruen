package service

import (
	"qkeruen/models"
	"qkeruen/repository"
)

type ProcessService interface {
	AcceptOrder(driverId, orderid int) (interface{}, error)
	CancellOrder(orderid int) (interface{}, error)
	GetOrdersInProcessDriver(driverId int) ([]*models.ProcessModel, error)
	GetOrdersInProcessUser(userId int) ([]*models.ProcessModel, error)
	FinishOrder(driverId, orderId int) (interface{}, error)
}

type processService struct {
	db repository.ProcessDB
}

func NewProcessService(ds repository.ProcessDB) *processService {
	return &processService{db: ds}
}

func (s *processService) AcceptOrder(driverId, orderid int) (interface{}, error) {
	res, err := s.db.AcceptOrder(driverId, orderid)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *processService) CancellOrder(orderId int) (interface{}, error) {
	res, err := s.db.CancellOrder(orderId)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *processService) GetOrdersInProcessDriver(driverId int) ([]*models.ProcessModel, error) {
	res, err := s.db.GetOrdersInProcessDriver(driverId)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *processService) GetOrdersInProcessUser(userId int) ([]*models.ProcessModel, error) {
	res, err := s.db.GetOrdersInProcessUser(userId)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *processService) FinishOrder(driverId, orderId int) (interface{}, error) {
	res, err := s.db.FinishOrder(driverId, orderId)

	if err != nil {
		return nil, err
	}

	return res, nil
}
