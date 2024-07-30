package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func NewConnDB(env Env) *sql.DB {
	config := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		env.DB_USERNAME,
		env.DB_PASSWORD,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_DATABASE,
		env.DB_SSLMODE,
	)

	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
