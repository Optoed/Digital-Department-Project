package repository

import (
	"backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(order *models.Order) error {

}

func (r *OrderRepo) GetNearestAndFree(courierID uint) (*models.Order, error) {

}
