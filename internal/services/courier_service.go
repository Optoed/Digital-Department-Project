package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type CourierService struct {
	repo *repository.CourierRepo
}

func NewCourierService(repo *repository.CourierRepo) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) Register(courier *models.Courier) (uint, error) {
	return s.repo.Create(courier)
}

func (s *CourierService) AuthenticateCourier(email, password string) (string, error) {
	courier, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(courier.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	return utils.GenerateJWT(courier.ID)
}

func (s *CourierService) TakeOrder(courierID, orderID uint) error {
	return s.repo.TakeOrder(courierID, orderID)
}

func (s *CourierService) FinishDelivery(courierID, orderID uint) error {
	return s.repo.FinishDelivery(courierID, orderID)
}

func (s *CourierService) GetDirections(courierID, orderID uint) (string, error) {
	// Заглушка или вызов внешнего API
	return "Mocked shortest route between A and B", nil
}

func (s *CourierService) GetByID(id uint) (*models.Courier, error) {
	return s.repo.GetByID(id)
}

func (s *CourierService) RateCourier(courierID uint, newRating float64) error {
	// Можно реализовать обновление средней оценки
	return nil
}
