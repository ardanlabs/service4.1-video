package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ardanlabs/service/business/data/dbmigrate"
	database "github.com/ardanlabs/service/business/sys/database/pgx"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	cfg := database.Config{
		User:         "postgres",
		Password:     "postgres",
		Host:         "database-service.sales-system.svc.cluster.local",
		Name:         "postgres",
		MaxIdleConns: 2,
		MaxOpenConns: 0,
		DisableTLS:   true,
	}

	if err := migrate(cfg); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := seed(cfg); err != nil {
		return fmt.Errorf("seed: %w", err)
	}

	return nil
}

func migrate(cfg database.Config) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbmigrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}

func seed(cfg database.Config) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbmigrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
