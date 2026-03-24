package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type EnterpriseRepository struct {
	db *pgxpool.Pool
}

func NewEnterpriseRepository(db *pgxpool.Pool) *EnterpriseRepository {
	return &EnterpriseRepository{db: db}
}

func (r *EnterpriseRepository) GetAll(ctx context.Context) ([]db.EnterpriseEntity, error) {
	query := `
		SELECT id, name, address, created_at
		FROM enterprises
		ORDER BY name ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enterprises []db.EnterpriseEntity
	for rows.Next() {
		var e db.EnterpriseEntity
		if err := rows.Scan(&e.ID, &e.Name, &e.Address, &e.CreatedAt); err != nil {
			return nil, err
		}
		enterprises = append(enterprises, e)
	}

	return enterprises, nil
}

func (r *EnterpriseRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.EnterpriseEntity, error) {
	query := `
		SELECT id, name, address, created_at
		FROM enterprises
		WHERE id = $1`

	var e db.EnterpriseEntity
	err := r.db.QueryRow(ctx, query, id).Scan(&e.ID, &e.Name, &e.Address, &e.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &e, nil
}
