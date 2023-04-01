package repository

import (
	"context"
	"qkeruen/dto"
	"qkeruen/models"
)

type OfferUserDB struct {
	DB PgxIface
}

func NewOfferUserRepository(ds PgxIface) OfferUserDB {
	return OfferUserDB{DB: ds}
}

func (pool OfferUserDB) Create(id int, offer dto.OfferRequest) error {
	q := `INSERT INTO offer_user(
		    comment,
		    locationFrom,
		    locationTo,
		    type,
		    customer
	)VALUES($1,$2,$3,$4,$5);`

	_, err := pool.DB.Exec(context.Background(), q, offer.Comment, offer.From, offer.To, offer.Type, id)

	if err != nil {
		return err
	}

	return nil
}

func (pool OfferUserDB) MyOffer(id int) ([]*dto.OfferResponseUser, error) {
	q := `Select * From offer_user WHERE customer=$1`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*dto.OfferResponseUser
	for rows.Next() {
		offer := new(dto.OfferResponseUser)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Type,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

// here ou must to change!
func (pool OfferUserDB) FindAllOffers() ([]*dto.OfferResponseUser, error) {
	q := `Select * From offer_user`
	rows, err := pool.DB.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return nil, nil
}

func (pool OfferUserDB) Search(to, from string) ([]*models.OfferUserModel, error) {
	q := `Select * From offer_driver WHERE locationFrom=$1 AND locationTo=$2;`

	rows, err := pool.DB.Query(context.Background(), q, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var offers []*models.OfferUserModel
	for rows.Next() {
		offer := new(models.OfferUserModel)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Type,
			&offer.Driver,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

func (pool OfferUserDB) Delete(offerId int) error {
	q := `Delete From offer_user WHERE Id = $1`

	_, err := pool.DB.Exec(context.Background(), q, offerId)

	if err != nil {
		return err
	}

	return nil
}
