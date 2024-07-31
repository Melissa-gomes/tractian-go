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

func (r Repository) Create(ctx context.Context, product domain.Products) error {
	stmt, err := r.DB.Prepare(`
		INSERT INTO products (
			"name",
			"price",
			"quantity"
		)
		VALUES (?, ?, ?)
		RETURNING id;
	`)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := stmt.Exec(product.Name, product.Price, product.Seller); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r Repository) Update(ctx context.Context, productToUpdate domain.Products) error {
	result, err := r.DB.ExecContext(ctx, `
		UPDATE 
			products 
		SET 
			name = ?,
			price = ?
		WHERE 
			id = ?;
	`, productToUpdate.Name, productToUpdate.Price, productToUpdate.Id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	return nil
}

func (r Repository) GetById() error {
	return nil
}

func (r Repository) Delete() error {
	return nil
}
