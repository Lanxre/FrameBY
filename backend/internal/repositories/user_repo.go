package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/types"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*db.UserEntity, error) {
	return r.findOne(ctx, "email", email)
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*db.UserEntity, error) {
	return r.findOne(ctx, "login", login)
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.UserEntity, error) {
	return r.findOne(ctx, "id", id)
}

func (r *UserRepository) findOne(ctx context.Context, field string, value any) (*db.UserEntity, error) {
	u := &db.UserEntity{}
	
	query := `
		SELECT 
			u.id, u.email, u.login, u.password_hash, u.role, u.created_at, u.updated_at,
			COALESCE(bp.full_name, sp.full_name, cp.full_name, '') as full_name,
			COALESCE(bp.subrole, '') as subrole,
			cp.enterprise_id
		FROM users u
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE u.` + field + ` = $1`

	err := r.db.QueryRow(ctx, query, value).Scan(
		&u.ID,
		&u.Email,
		&u.Login,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.FullName,
		&u.Subrole,
		&u.EnterpriseID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Create(ctx context.Context, email, login, passwordHash string) (uuid.UUID, error) {
	var id uuid.UUID
	query := `
		INSERT INTO users (email, login, password_hash) 
		VALUES ($1, $2, $3) 
		RETURNING id`

	err := r.db.QueryRow(ctx, query, email, login, passwordHash).Scan(&id)
	return id, err
}

func (r *UserRepository) AssignProfile(ctx context.Context, userID uuid.UUID, role string, fullName string, subrole string, enterpriseID *uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE users SET role = $1 WHERE id = $2", role, userID)
	if err != nil {
		return err
	}

	switch role {
	case string(types.RoleBRSM):
		_, err = tx.Exec(ctx, 
			"INSERT INTO brsm_profiles (user_id, full_name, subrole) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET full_name = EXCLUDED.full_name, subrole = EXCLUDED.subrole", 
			userID, fullName, subrole)
	case string(types.RoleStudent):
		_, err = tx.Exec(ctx, 
			"INSERT INTO student_profiles (user_id, full_name) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET full_name = EXCLUDED.full_name", 
			userID, fullName)
	case string(types.RoleUniversity):
		_, err = tx.Exec(ctx, 
			"INSERT INTO university_profiles (user_id, full_name) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET full_name = EXCLUDED.full_name", 
			userID, fullName)
	case string(types.RoleCustomer):
		_, err = tx.Exec(ctx, 
			"INSERT INTO customer_profiles (user_id, full_name, enterprise_id) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET full_name = EXCLUDED.full_name, enterprise_id = EXCLUDED.enterprise_id", 
			userID, fullName, enterpriseID)
	default:
		return errors.New("invalid role")
	}

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *UserRepository) UpdatePassword(ctx context.Context, userID uuid.UUID, newHash string) error {
	query := `UPDATE users SET password_hash = $1 WHERE id = $2`
	_, err := r.db.Exec(ctx, query, newHash, userID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}