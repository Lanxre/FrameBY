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
	
	if err := runSeeds(pool); err != nil {
		return nil, fmt.Errorf("seed failed: %w", err)
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

func runSeeds(pool *pgxpool.Pool) error {
	const seedSQL = `INSERT INTO public.users (id, email, login, password_hash, role, created_at, updated_at, avatar)
VALUES (
    'a3204e6e-3300-4f3b-839c-abb42f65a417'::uuid,
    'admin@example.com',
    'admin',
    '$2a$10$vlPm8u0NJUZta7vhBNs/mek9QiX5cK7lMPBH9k9vN1I45zT1ZwZYm',
    'admin',
    NOW(),
    NOW(),
    ''
)
ON CONFLICT (id) DO UPDATE
SET 
    email = EXCLUDED.email,
    login = EXCLUDED.login,
    password_hash = EXCLUDED.password_hash,
    role = EXCLUDED.role,
    updated_at = NOW(),
    avatar = EXCLUDED.avatar;`;

	_, err := pool.Exec(context.Background(), seedSQL)
	if err != nil {
		return fmt.Errorf("failed to execute seed SQL: %w", err)
	}

	return nil
} 