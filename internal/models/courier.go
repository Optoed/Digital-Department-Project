package models

const (
	OnFoot  = "on_foot"
	Scooter = "scooter"
	Car     = "car"
)

type Courier struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic string  `json:"patronymic"` // отчество
	Transport  string  `json:"transport"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Rating     float32 `json:"rating"`
	Orders     []Order `json:"orders"`
}
