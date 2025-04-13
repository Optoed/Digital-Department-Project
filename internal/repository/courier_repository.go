// repository/courier_repository.go
package repository

import (
	"backend/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
)

type CourierRepo struct {
	db *sqlx.DB
}

func NewCourierRepo(db *sqlx.DB) *CourierRepo {
	return &CourierRepo{db: db}
}

func (r *CourierRepo) Create(courier *models.Courier) (uint, error) {
	query := `
		INSERT INTO couriers (name, surname, transport, email, phone, rating, is_available)
		VALUES (:name, :surname, :transport, :email, :phone, :rating, :is_available)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return 0, err
	}

	err = stmt.Get(&courier.ID, courier)
	if err != nil {
		return 0, err
	}

	return courier.ID, nil
}

func (r *CourierRepo) GetByID(id uint) (*models.Courier, error) {
	var courier models.Courier
	err := r.db.Get(&courier, `SELECT * FROM couriers WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &courier, nil
}

func (r *CourierRepo) GetByEmail(email string) (*models.Courier, error) {
	var courier models.Courier
	err := r.db.Get(&courier, `SELECT * FROM couriers WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	return &courier, nil
}

func (r *CourierRepo) TakeOrder(courierID, orderID uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем заказ
	_, err = tx.Exec(`
		UPDATE orders
		SET courier_id = $1, status = 'in_transit', assigned_at = NOW()
		WHERE id = $2 AND status = 'created'`, courierID, orderID)
	if err != nil {
		return err
	}

	// Обновляем курьера
	_, err = tx.Exec(`
		UPDATE couriers
		SET current_order_id = $1, is_available = FALSE
		WHERE id = $2`, orderID, courierID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Получить кратчайший маршрут
func (r *CourierRepo) GetDirections(courierID, orderID uint) (models.RouteInfo, error) {
	var route models.RouteInfo

	// Заглушка — подгружаем адреса из базы
	err := r.db.Get(&route.FromAddress, `SELECT address_from FROM orders WHERE id = $1`, orderID)
	if err != nil {
		return route, err
	}
	err = r.db.Get(&route.ToAddress, `SELECT address_to FROM orders WHERE id = $1`, orderID)
	if err != nil {
		return route, err
	}

	// Временные значения
	route.DistanceKm = rand.Float64() * 5     // от 0 до 5 км
	route.DurationMin = route.DistanceKm * 10 // 1 км = 10 минут
	log.Println("Stub route used")

	return route, nil
}

func (r *CourierRepo) FinishDelivery(courierID, orderID uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем заказ
	_, err = tx.Exec(`
		UPDATE orders
		SET status = 'delivered', finished_at = NOW()
		WHERE id = $1 AND courier_id = $2`, orderID, courierID)
	if err != nil {
		return err
	}

	// Обновляем курьера
	_, err = tx.Exec(`
		UPDATE couriers
		SET current_order_id = NULL, is_available = TRUE
		WHERE id = $1`, courierID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
