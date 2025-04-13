// models/order.go
package models

import "time"

type OrderStatus string

const (
	Created   OrderStatus = "created"
	InTransit OrderStatus = "in_transit"
	Delivered OrderStatus = "delivered"
	Canceled  OrderStatus = "canceled"
)

type Order struct {
	ID            uint        `db:"id" json:"id"`
	CourierID     *uint       `db:"courier_id" json:"courier_id"`
	Status        OrderStatus `db:"status" json:"status"`
	ReceiverName  string      `db:"receiver_name" json:"receiver_name"`
	ReceiverPhone string      `db:"receiver_phone" json:"receiver_phone"`
	AddressFrom   string      `db:"address_from" json:"address_from"`
	AddressTo     string      `db:"address_to" json:"address_to"`
	Cost          int         `db:"cost" json:"cost"`
	CreatedAt     time.Time   `db:"created_at" json:"created_at"`
	AssignedAt    *time.Time  `db:"assigned_at" json:"assigned_at"`
	FinishedAt    *time.Time  `db:"finished_at" json:"finished_at"`
}
