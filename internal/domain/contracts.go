package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ServicePorts interface {
	List(ctx context.Context) ([]Products, error)
	Create()
	Update()
	GetById()
	Delete()
}

type RepositoryPorts interface {
	List(ctx context.Context) ([]Products, error)
	Create()
	Update()
	GetById()
	Delete()
}

type Products struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Seller    string    `json:"seller"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
