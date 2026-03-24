package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lanxre/frameby/internal/config"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
)

const (
	DefaultTimeout = 5 * time.Second
)

func NewPostgresDB(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	if err := runMigrations(cfg.DatabaseURL); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Closing database pool...")
			pool.Close()
			return nil
		},
	})

	return pool, nil
}

func runMigrations(dbURL string) error {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}
	
	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Database migrations check completed")
	return nil
}