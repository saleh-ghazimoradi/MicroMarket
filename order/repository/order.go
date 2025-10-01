package repository

import "database/sql"

type OrderRepository interface{}

type orderRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func NewOrderRepository(dbWrite, dbRead *sql.DB) OrderRepository {
	return &orderRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
