// Пока не нужно

package models

type CourierOrder struct {
	ID        uint        `db:"id"`
	CourierID uint        `db:"courier_id"`
	OrderID   uint        `db:"order_id"`
	Status    OrderStatus `db:"status"`
}
