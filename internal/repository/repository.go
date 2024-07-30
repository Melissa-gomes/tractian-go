package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tractian-go/internal/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepo(database *sql.DB) Repository {
	return Repository{
		DB: database,
	}
}

func (r Repository) List(ctx context.Context) ([]domain.Products, error) {
	var products []domain.Products
	rows, err := r.DB.QueryContext(ctx, "SELECT id, nome FROM products")
	if err != nil {
		return []domain.Products{}, fmt.Errorf("err to find in DB: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product)
		if err != nil {
			return []domain.Products{}, fmt.Errorf("err to scan result line: %s", err.Error())
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Erro durante a iteração das linhas: %v", err)
	}

	return products, nil
}

// Id
// Name
// Price
// Seller
// CreatedAt
// UpdatedAt

// func (r Repository) Create(ctx context.Context, Product domain.Products) error {
// 	rows, err := r.DB.QueryContext(ctx, `
// 		INSERT INTO products (
// 			name,
// 			price,
// 			seller
// 		)
// 		VALUES ( ?, ?, ?)
// 		RETURNING id;
// 	`, Product.Name, Product.Price, Product.Seller)
// 	return nil
// }

func (r Repository) Update() error {
	return nil
}

func (r Repository) GetById() error {
	return nil
}

func (r Repository) Delete() error {
	return nil
}
