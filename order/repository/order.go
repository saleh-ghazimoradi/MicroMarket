package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/saleh-ghazimoradi/MicroMarket/order/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrdersForAccount(ctx context.Context, accountId string) ([]*domain.Order, error)
	Close() error
}

type orderRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func (o *orderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	tx, err := o.dbWrite.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	// Insert order
	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO "order"(id, created_at, account_id, total_price) VALUES($1, $2, $3, $4)`,
		order.Id,
		order.CreatedAt,
		order.AccountId,
		order.TotalPrice,
	)
	if err != nil {
		return err
	}
	// Insert order catalog
	stmt, _ := tx.PrepareContext(ctx, pq.CopyIn("order_catalog", "order_id", "catalog_id", "quantity"))
	for _, c := range order.Catalogs {
		_, err = stmt.ExecContext(ctx, order.Id, c.Id, c.Quantity)
		if err != nil {
			return err
		}
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	stmt.Close()
	return err
}

func (o *orderRepository) GetOrdersForAccount(ctx context.Context, accountId string) ([]*domain.Order, error) {
	rows, err := o.dbRead.QueryContext(
		ctx,
		`SELECT
		      o.id,
		      o.created_at,
		      o.account_id,
		      o.total_price::money::numeric::float8,
		      oc.catalog_id,
		      oc.quantity,
		      c.name,
		      c.description,
		      c.price::money::numeric::float8
		    FROM "order" o 
		    JOIN order_catalog oc ON (o.id = oc.order_id)
		    JOIN catalogs c ON (oc.catalog_id = c.id)
		    WHERE o.account_id = $1
		    ORDER BY o.id`,
		accountId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var orders []*domain.Order
	var order domain.Order
	lastOrder := domain.Order{}
	var orderedCatalog domain.OrderedCatalog
	var catalogs []domain.OrderedCatalog

	for rows.Next() {
		if err = rows.Scan(
			&order.Id,
			&order.CreatedAt,
			&order.AccountId,
			&order.TotalPrice,
			&orderedCatalog.Id,
			&orderedCatalog.Quantity,
			&orderedCatalog.Name,
			&orderedCatalog.Description,
			&orderedCatalog.Price,
		); err != nil {
			return nil, err
		}
		// Check for new order
		if lastOrder.Id != "" && lastOrder.Id != order.Id {
			newOrder := &domain.Order{
				Id:         lastOrder.Id,
				AccountId:  lastOrder.AccountId,
				CreatedAt:  lastOrder.CreatedAt,
				TotalPrice: lastOrder.TotalPrice,
				Catalogs:   catalogs,
			}
			orders = append(orders, newOrder)
			catalogs = []domain.OrderedCatalog{}
		}
		// Add catalog
		catalogs = append(catalogs, orderedCatalog)
		lastOrder = order
	}
	// Add last order
	if lastOrder.Id != "" {
		newOrder := &domain.Order{
			Id:         lastOrder.Id,
			AccountId:  lastOrder.AccountId,
			CreatedAt:  lastOrder.CreatedAt,
			TotalPrice: lastOrder.TotalPrice,
			Catalogs:   catalogs,
		}
		orders = append(orders, newOrder)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderRepository) Close() error {
	if err := o.dbWrite.Close(); err != nil {
		return err
	}
	if err := o.dbRead.Close(); err != nil {
		return err
	}
	return nil
}

func NewOrderRepository(dbWrite, dbRead *sql.DB) OrderRepository {
	return &orderRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
