package repository

import (
	"context"
	"qkeruen/dto"
)

type OrderDB struct {
	DB PgxIface
}

func NewOrderRepository(ds PgxIface) OrderDB {
	return OrderDB{DB: ds}
}

func (pool OrderDB) CreateOrder(userId int, order dto.OrderRequest) error {
	q := `INSERT INTO order_process(
		userId,
		latitudeFrom,
		longitudeFrom,
		latitudeTo,
		longitudeTo,
		comments,
		price,
		type,
		orderStatus
	)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);`

	_, err := pool.DB.Query(
		context.Background(),
		q,
		order.UserId,
		order.LatitudeFrom,
		order.LongitudeFrom,
		order.LatitudeTo,
		order.LongitudeTo,
		order.Comments,
		order.Price,
		order.Type,
		order.OrderStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pool OrderDB) GetDriverType(driverId int) (string, error) {
	q := `Select carType From driver WHERE id=$1`

	row := pool.DB.QueryRow(context.Background(), q, driverId)

	var Type string

	err := row.Scan(&Type)

	if err != nil {
		return "", err
	}

	return Type, nil
}

func (pool OrderDB) GetOrders(driverId int) ([]*dto.OrderResponse, error) {
	q := `Select * From order_process WHERE orderStatus=$1`

	rows, err := pool.DB.Query(context.Background(), q, 0)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	carType, err := pool.GetDriverType(driverId)

	if err != nil {
		return nil, err
	}

	var orders []*dto.OrderResponse

	for rows.Next() {
		order := new(dto.OrderResponse)

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.LatitudeFrom,
			&order.LongitudeFrom,
			&order.LatitudeTo,
			&order.LongitudeTo,
			&order.Comments,
			&order.Price,
			&order.Type,
			&order.OrderStatus,
		)

		if err != nil {
			return nil, err
		}
		if order.Type == carType {
			orders = append(orders, order)
		}

	}

	return orders, nil
}

func (pool OrderDB) GetMyOrders(id int) ([]*dto.OrderResponse, error) {
	q := `Select * From order_process WHERE userId=$1`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*dto.OrderResponse

	for rows.Next() {
		order := new(dto.OrderResponse)

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.LatitudeFrom,
			&order.LongitudeFrom,
			&order.LatitudeTo,
			&order.LongitudeTo,
			&order.Comments,
			&order.Price,
			&order.Type,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (pool OrderDB) DeleteOrder(orderId int) error {
	q := `Delete From order_process WHERE Id=$1`

	_, err := pool.DB.Query(context.Background(), q, orderId)

	if err != nil {
		return err
	}

	return nil
}
