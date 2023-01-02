package service

import (
	"qkeruen/dto"
	"qkeruen/models"
	"qkeruen/repository"
)

type OfferUserService interface {
	Create(id int, offer dto.OfferRequest) error
	MyOffer(id int) ([]*dto.OfferResponseUser, error)
	FindAllOffers() ([]*dto.OfferResponseUser, error)
	Search(to, from string) ([]*models.OfferUserModel, error)
	DeleteOffer(offerId int) error
}

type offerUserService struct {
	db repository.OfferUserDB
}

func NewOfferuserService(ds repository.OfferUserDB) *offerUserService {
	return &offerUserService{
		db: ds,
	}
}

func (s *offerUserService) Create(id int, offer dto.OfferRequest) error {
	return s.db.Create(id, offer)
}

func (s *offerUserService) MyOffer(id int) ([]*dto.OfferResponseUser, error) {
	res, err := s.db.MyOffer(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) FindAllOffers() ([]*dto.OfferResponseUser, error) {
	res, err := s.db.FindAllOffers()

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) Search(to, from string) ([]*models.OfferUserModel, error) {
	res, err := s.db.Search(to, from)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) DeleteOffer(offerId int) error {
	return s.db.Delete(offerId)
}
