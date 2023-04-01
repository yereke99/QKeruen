package service

import (
	"qkeruen/dto"
	"qkeruen/models"
	"qkeruen/repository"
)

type OfferDriverSevvice interface {
	CreateOffer(id int, offer dto.OfferRequest) error
	MyOffer(id int) ([]*dto.OfferResponseDriver, error)
	FindAllOffers() ([]*models.OfferDriverModel, error)
	SearchOffers(to, from string) ([]*models.OfferDriverModel, error)
	Delete(offerId int) error
}

type offerDriverService struct {
	db repository.OfferDriverDB
}

func NewOfferDriverService(ds repository.OfferDriverDB) *offerDriverService {
	return &offerDriverService{
		db: ds,
	}
}

func (s *offerDriverService) CreateOffer(id int, offer dto.OfferRequest) error {
	return s.db.Create(id, offer)
}

func (s *offerDriverService) MyOffer(id int) ([]*dto.OfferResponseDriver, error) {
	res, err := s.db.MyOffer(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *offerDriverService) FindAllOffers() ([]*models.OfferDriverModel, error) {
	res, err := s.db.FindAllOffers()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *offerDriverService) SearchOffers(to, from string) ([]*models.OfferDriverModel, error) {
	res, err := s.db.Search(to, from)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *offerDriverService) Delete(offerId int) error {
	return s.db.Delete(offerId)
}
