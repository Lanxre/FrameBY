package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type UniversityDepartmentRepository struct {
	db *pgxpool.Pool
}

func NewUniversityDepartmentRepository(db *pgxpool.Pool) *UniversityDepartmentRepository {
	return &UniversityDepartmentRepository{db: db}
}

func (r *UniversityDepartmentRepository) GetAll(ctx context.Context) ([]db.UniversityDepartmentEntity, error) {
	query := `
		SELECT id, university_name, department_name, address, created_at
		FROM university_departments
		ORDER BY university_name ASC, department_name ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []db.UniversityDepartmentEntity
	for rows.Next() {
		var d db.UniversityDepartmentEntity
		if err := rows.Scan(&d.ID, &d.UniversityName, &d.DepartmentName, &d.Address, &d.CreatedAt); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	return departments, nil
}

func (r *UniversityDepartmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.UniversityDepartmentEntity, error) {
	query := `
		SELECT id, university_name, department_name, address, created_at
		FROM university_departments
		WHERE id = $1`

	var d db.UniversityDepartmentEntity
	err := r.db.QueryRow(ctx, query, id).Scan(&d.ID, &d.UniversityName, &d.DepartmentName, &d.Address, &d.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &d, nil
}
