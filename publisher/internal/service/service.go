package service

import (
	"fmt"
	"service/internal/config"
)

type Service interface {
	Echo(number int) string
}

type service struct{}

func NewService(cfg config.Service) (Service, error) {
	return &service{}, nil
}

func (s *service) Echo(number int) string {
	return fmt.Sprintf("%d", number)
}
