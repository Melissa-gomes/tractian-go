package service

import (
	"context"
	"tractian-go/internal/domain"
	"tractian-go/internal/repository"
)

type Service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s Service) List(ctx context.Context) ([]domain.Products, error) {
	products, err := s.Repo.List(ctx)
	if err != nil {
		return []domain.Products{}, err
	}

	return products, nil
}
