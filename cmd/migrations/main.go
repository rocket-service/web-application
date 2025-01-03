package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	defaultDSN            = "postgres://postgres:test-password@localhost:5432/rocket_db?sslmode=disable"
	defaultMigrationsPath = "./migrations"
)

func main() {
	dsn := flag.String("dsn", defaultDSN, "PostgreSQL data source name")
	up := flag.Bool("up", false, "Run migrations up")
	down := flag.Bool("down", false, "Run migrations down")

	flag.Parse()

	if *up && *down {
		log.Fatalf("Cannot specify both --up and --down at the same time")
	}
	if !*up && !*down {
		log.Fatalf("Specify at least one action: --up or --down")
	}

	db, cleanup, err := setupDatabase(*dsn)
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}
	defer cleanup()

	if err := runMigrations(db, *up, *down); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations applied successfully!")
}

func setupDatabase(dsn string) (*sql.DB, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	cleanup := func() {
		pool.Close()
	}

	db := stdlib.OpenDBFromPool(pool)

	if err := goose.SetDialect("postgres"); err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("failed to set goose dialect: %w", err)
	}

	return db, cleanup, nil
}

func runMigrations(db *sql.DB, up, down bool) error {
	migrationsPath := defaultMigrationsPath

	if up {
		if err := goose.Up(db, migrationsPath); err != nil {
			return fmt.Errorf("failed to apply up migrations: %w", err)
		}
	}

	if down {
		if err := goose.Down(db, migrationsPath); err != nil {
			return fmt.Errorf("failed to apply down migrations: %w", err)
		}
	}

	return nil
}
