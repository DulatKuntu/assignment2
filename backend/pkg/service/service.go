package service

import (
	"bonus/pkg/repository"
)

type All interface {
}

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
