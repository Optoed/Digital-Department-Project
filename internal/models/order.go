package models

import "time"

// Status constants for Order status
const (
	Created   = "created"
	InTransit = "in_transit"
	Delivered = "delivered"
	Canceled  = "canceled"
)

type Order struct {
	ID              int        `json:"id"`
	Status          string     `json:"status"`
	SenderID        int        `json:"sender_id"`
	Receiver        string     `json:"receiver"`
	Cost            int        `json:"cost"`
	CourierID       int        `json:"courier_id"`
	AssignedCourier *Courier   `json:"assigned_courier,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}
