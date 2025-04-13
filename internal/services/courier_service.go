package services

import "backend/internal/repository"

type CourierService struct {
	repo *repository.CourierRepo
}

func NewCourierService(repo *repository.CourierRepo) *CourierService {
	return &CourierService{repo: repo}
}
