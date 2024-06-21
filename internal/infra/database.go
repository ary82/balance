package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSQLDB(dburl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return db, nil
}
