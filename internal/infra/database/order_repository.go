package database

import (
	"database/sql"

	"github.com.caiommdev/go-intensivo/internal/entities"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRespository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entities.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}