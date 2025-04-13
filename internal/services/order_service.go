package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepo
}

func NewOrderService(repo *repository.OrderRepo) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order *models.Order) (uint, error) {
	return s.repo.Create(order)
}

func (s *OrderService) GetNearestAndFree(courierID uint) (*models.Order, error) {
	return s.repo.GetNearestAndFree(courierID)
}
