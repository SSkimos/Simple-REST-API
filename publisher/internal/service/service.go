package service

import (
	"fmt"
	"service/internal/config"
	"service/internal/models"
)

type Service interface {
	Echo(number int) string
	EmployeeAdd(request *models.Request) string
}

type service struct{}

func NewService(cfg config.Service) (Service, error) {
	return &service{}, nil
}

func (s *service) Echo(number int) string {
	return fmt.Sprintf("%d", number)
}

func (s *service) EmployeeAdd(request *models.Request) string {
	return "OK"
}
