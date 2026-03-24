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

func (r *UniversityDepartmentRepository) GetAll(ctx context.Context) ([]db.UniversityDepartmentWithDetails, error) {
	query := `
		SELECT ud.id, u.name as university_name, d.name as department_name, ud.address
		FROM university_departments ud
		JOIN universities u ON ud.university_id = u.id
		JOIN departments d ON ud.department_id = d.id
		ORDER BY u.name ASC, d.name ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []db.UniversityDepartmentWithDetails
	for rows.Next() {
		var d db.UniversityDepartmentWithDetails
		if err := rows.Scan(&d.ID, &d.UniversityName, &d.DepartmentName, &d.Address); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	return departments, nil
}

func (r *UniversityDepartmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.UniversityDepartmentEntity, error) {
	query := `
		SELECT ud.id, ud.university_id, ud.department_id, ud.address, ud.created_at,
			   u.id, u.name, u.created_at,
			   d.id, d.name, d.created_at
		FROM university_departments ud
		JOIN universities u ON ud.university_id = u.id
		JOIN departments d ON ud.department_id = d.id
		WHERE ud.id = $1`

	var d db.UniversityDepartmentEntity
	var uni db.University
	var dept db.Department

	err := r.db.QueryRow(ctx, query, id).Scan(
		&d.ID, &d.UniversityID, &d.DepartmentID, &d.Address, &d.CreatedAt,
		&uni.ID, &uni.Name, &uni.CreatedAt,
		&dept.ID, &dept.Name, &dept.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	d.University = &uni
	d.Department = &dept

	return &d, nil
}

func (r *UniversityDepartmentRepository) Create(ctx context.Context, universityName, departmentName string, address *string) (*db.UniversityDepartmentWithDetails, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var universityID, departmentID uuid.UUID

	err = tx.QueryRow(ctx, `
		INSERT INTO universities (name) VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
		RETURNING id
	`, universityName).Scan(&universityID)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, `
		INSERT INTO departments (name) VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
		RETURNING id
	`, departmentName).Scan(&departmentID)
	if err != nil {
		return nil, err
	}

	var udID uuid.UUID
	err = tx.QueryRow(ctx, `
		INSERT INTO university_departments (university_id, department_id, address)
		VALUES ($1, $2, $3)
		ON CONFLICT (university_id, department_id) DO UPDATE SET address = EXCLUDED.address
		RETURNING id
	`, universityID, departmentID, address).Scan(&udID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &db.UniversityDepartmentWithDetails{
		ID:             udID,
		UniversityName: universityName,
		DepartmentName: departmentName,
		Address:        address,
	}, nil
}
