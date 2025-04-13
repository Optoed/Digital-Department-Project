// models/courier.go
package models

import "time"

type TransportType string

const (
	OnFoot  TransportType = "on_foot"
	Scooter TransportType = "scooter"
	Car     TransportType = "car"
)

type Courier struct {
	ID             uint          `db:"id" json:"id"`
	Name           string        `db:"name" json:"name"`
	Surname        string        `db:"surname" json:"surname"`
	Transport      TransportType `db:"transport" json:"transport"`
	Email          string        `db:"email" json:"email"`
	Phone          string        `db:"phone" json:"phone"`
	Rating         float32       `db:"rating" json:"rating"`
	CurrentOrderID *uint         `db:"current_order_id" json:"current_order_id,omitempty"`
	IsAvailable    bool          `db:"is_available" json:"is_available"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at"`
	Password       string        `db:"password" json:"password"`
}
