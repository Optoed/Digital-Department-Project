// repository/courier_repository.go
package repository

import (
	"backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type CourierRepo struct {
	db *sqlx.DB
}

func NewCourierRepo(db *sqlx.DB) *CourierRepo {
	return &CourierRepo{db: db}
}

func (r *CourierRepo) Create(courier *models.Courier) error {

}

func (r *CourierRepo) GetByID(id uint) (*models.Courier, error) {

}

func (r *CourierRepo) TakeOrder(courierID, orderID uint) error {

}

// Получить кратчайший маршрут
func (r *CourierRepo) GetDirections(courierID, orderID uint) error {

}

func (r *CourierRepo) FinishDelivery(courierID, orderID uint) error {

}
