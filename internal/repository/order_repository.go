package repository

import (
	"backend/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(order *models.Order) (uint, error) {
	query := `
		INSERT INTO orders (receiver_name, receiver_phone, address_from, address_to, cost)
		VALUES (:receiver_name, :receiver_phone, :address_from, :address_to, :cost)
		RETURNING id
	`
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return 0, err
	}

	err = stmt.Get(&order.ID, order)
	return order.ID, err
}

func (r *OrderRepo) GetNearestAndFree(courierID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Get(&order, `
		SELECT * FROM orders
		WHERE courier_id IS NULL AND status = 'created'
		ORDER BY created_at
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}

	log.Println("Stub route used")

	return &order, nil
}
