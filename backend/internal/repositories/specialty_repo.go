package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type SpecialtyRepository struct {
	db *pgxpool.Pool
}

func NewSpecialtyRepository(db *pgxpool.Pool) *SpecialtyRepository {
	return &SpecialtyRepository{db: db}
}

func (r *SpecialtyRepository) GetByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]db.SpecialtyEntity, error) {
	query := `
		SELECT id, name, department_id, created_at
		FROM specialties
		WHERE department_id = $1
		ORDER BY name ASC`

	rows, err := r.db.Query(ctx, query, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []db.SpecialtyEntity
	for rows.Next() {
		var s db.SpecialtyEntity
		if err := rows.Scan(&s.ID, &s.Name, &s.DepartmentID, &s.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (r *SpecialtyRepository) GetByUniversityDepartmentID(ctx context.Context, universityDepartmentID uuid.UUID) ([]db.SpecialtyEntity, error) {
	query := `
		SELECT s.id, s.name, s.department_id, s.created_at
		FROM specialties s
		JOIN university_departments ud ON ud.department_id = s.department_id
		WHERE ud.id = $1
		ORDER BY s.name ASC`

	rows, err := r.db.Query(ctx, query, universityDepartmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []db.SpecialtyEntity
	for rows.Next() {
		var s db.SpecialtyEntity
		if err := rows.Scan(&s.ID, &s.Name, &s.DepartmentID, &s.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}
