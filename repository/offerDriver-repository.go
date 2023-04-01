package repository

import (
	"context"
	"qkeruen/dto"
	"qkeruen/models"
)

type OfferDriverDB struct {
	DB PgxIface
}

func NewOfferDriverRepository(ds PgxIface) OfferDriverDB {
	return OfferDriverDB{
		DB: ds,
	}
}

func (pool OfferDriverDB) Create(id int, offer dto.OfferRequest) error {
	q := `INSERT INTO offer_driver(
		comment,
		locationFrom,
		locationTo,
		type,
		driver
	)VALUES($1,$2,$3,$4,$5);`

	_, err := pool.DB.Exec(context.Background(), q, offer.Comment, offer.From, offer.To, offer.Type, id)

	if err != nil {
		return err
	}

	return nil
}

func (pool OfferDriverDB) MyOffer(id int) ([]*dto.OfferResponseDriver, error) {
	q := `Select * From offer_driver WHERE driver=$1`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*dto.OfferResponseDriver
	for rows.Next() {
		offer := new(dto.OfferResponseDriver)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

func (pool OfferDriverDB) FindAllOffers() ([]*models.OfferDriverModel, error) {
	q := `Select * From offer_driver`
	rows, err := pool.DB.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return nil, nil
}

func (pool OfferDriverDB) Search(to, from string) ([]*models.OfferDriverModel, error) {
	q := `Select * From offer_user WHERE locationFrom=$1 AND locationTo=$2;`

	rows, err := pool.DB.Query(context.Background(), q, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var offers []*models.OfferDriverModel
	for rows.Next() {
		offer := new(models.OfferDriverModel)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Type,
			&offer.User,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

func (pool OfferDriverDB) Delete(offerId int) error {
	q := `Delete From offer_driver WHERE Id = $1`

	_, err := pool.DB.Exec(context.Background(), q, offerId)

	if err != nil {
		return err
	}

	return nil
}
