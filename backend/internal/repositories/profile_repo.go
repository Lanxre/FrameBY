package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetBrsmProfile(ctx context.Context, userID uuid.UUID) (*db.BrsmProfileEntity, error) {
	p := &db.BrsmProfileEntity{}
	query := `
		SELECT user_id, full_name, subrole, position, phone, updated_at
		FROM brsm_profiles
		WHERE user_id = $1`

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&p.UserID,
		&p.FullName,
		&p.Subrole,
		&p.Position,
		&p.Phone,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) CreateBrsmProfile(ctx context.Context, userID uuid.UUID, fullName, subrole string, position, phone *string) error {
	query := `
		INSERT INTO brsm_profiles (user_id, full_name, subrole, position, phone)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			subrole = EXCLUDED.subrole,
			position = EXCLUDED.position,
			phone = EXCLUDED.phone`

	_, err := r.db.Exec(ctx, query, userID, fullName, subrole, position, phone)
	return err
}

func (r *ProfileRepository) UpdateBrsmProfile(ctx context.Context, userID uuid.UUID, fullName, subrole string, position, phone *string) error {
	query := `
		UPDATE brsm_profiles
		SET full_name = $2, subrole = $3, position = $4, phone = $5
		WHERE user_id = $1`

	result, err := r.db.Exec(ctx, query, userID, fullName, subrole, position, phone)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("profile not found")
	}
	return nil
}

func (r *ProfileRepository) DeleteBrsmProfile(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM brsm_profiles WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *ProfileRepository) GetStudentProfile(ctx context.Context, userID uuid.UUID) (*db.StudentProfileEntity, error) {
	p := &db.StudentProfileEntity{}
	query := `
		SELECT user_id, full_name, position, phone, specialty, grade, university_department_id, updated_at
		FROM student_profiles
		WHERE user_id = $1`

	var specialty sql.NullString
	var grade sql.Null[float64]
	var univDeptID sql.Null[uuid.UUID]

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&p.UserID,
		&p.FullName,
		&p.Position,
		&p.Phone,
		&specialty,
		&grade,
		&univDeptID,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if specialty.Valid {
		p.Specialty = &specialty.String
	}
	if grade.Valid {
		p.Grade = &grade.V
	}
	if univDeptID.Valid {
		p.UniversityDepartmentID = &univDeptID.V
	}
	return p, nil
}

func (r *ProfileRepository) CreateStudentProfile(ctx context.Context, userID uuid.UUID, fullName string, specialty *string, grade *float64, position, phone *string, universityDeptID *uuid.UUID) error {
	query := `
		INSERT INTO student_profiles (user_id, full_name, specialty, grade, position, phone, university_department_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (user_id) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			specialty = EXCLUDED.specialty,
			grade = EXCLUDED.grade,
			position = EXCLUDED.position,
			phone = EXCLUDED.phone,
			university_department_id = EXCLUDED.university_department_id`

	_, err := r.db.Exec(ctx, query, userID, fullName, specialty, grade, position, phone, universityDeptID)
	return err
}

func (r *ProfileRepository) UpdateStudentProfile(ctx context.Context, userID uuid.UUID, fullName string, specialty *string, grade *float64, position, phone *string, universityDeptID *uuid.UUID) error {
	query := `
		UPDATE student_profiles
		SET full_name = $2, specialty = $3, grade = $4, position = $5, phone = $6, university_department_id = $7
		WHERE user_id = $1`

	result, err := r.db.Exec(ctx, query, userID, fullName, specialty, grade, position, phone, universityDeptID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("profile not found")
	}
	return nil
}

func (r *ProfileRepository) DeleteStudentProfile(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM student_profiles WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *ProfileRepository) GetUniversityProfile(ctx context.Context, userID uuid.UUID) (*db.UniversityProfileEntity, error) {
	p := &db.UniversityProfileEntity{}
	query := `
		SELECT user_id, full_name, subrole, position, phone, university_department_id, updated_at
		FROM university_profiles
		WHERE user_id = $1`

	var univDeptID sql.Null[uuid.UUID]

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&p.UserID,
		&p.FullName,
		&p.Subrole,
		&p.Position,
		&p.Phone,
		&univDeptID,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if univDeptID.Valid {
		p.UniversityDepartmentID = &univDeptID.V
	}
	return p, nil
}

func (r *ProfileRepository) CreateUniversityProfile(ctx context.Context, userID uuid.UUID, fullName, subrole string, position, phone *string, universityDeptID *uuid.UUID) error {
	query := `
		INSERT INTO university_profiles (user_id, full_name, subrole, position, phone, university_department_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_id) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			subrole = EXCLUDED.subrole,
			position = EXCLUDED.position,
			phone = EXCLUDED.phone,
			university_department_id = EXCLUDED.university_department_id`

	_, err := r.db.Exec(ctx, query, userID, fullName, subrole, position, phone, universityDeptID)
	return err
}

func (r *ProfileRepository) UpdateUniversityProfile(ctx context.Context, userID uuid.UUID, fullName, subrole string, position, phone *string, universityDeptID *uuid.UUID) error {
	query := `
		UPDATE university_profiles
		SET full_name = $2, subrole = $3, position = $4, phone = $5, university_department_id = $6
		WHERE user_id = $1`

	result, err := r.db.Exec(ctx, query, userID, fullName, subrole, position, phone, universityDeptID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("profile not found")
	}
	return nil
}

func (r *ProfileRepository) DeleteUniversityProfile(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM university_profiles WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *ProfileRepository) GetCustomerProfile(ctx context.Context, userID uuid.UUID) (*db.CustomerProfileEntity, error) {
	p := &db.CustomerProfileEntity{}
	query := `
		SELECT user_id, full_name, subrole, enterprise_id, position, phone, updated_at
		FROM customer_profiles
		WHERE user_id = $1`

	var enterpriseID sql.NullString
	var subrole sql.NullString
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&p.UserID,
		&p.FullName,
		&subrole,
		&enterpriseID,
		&p.Position,
		&p.Phone,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if subrole.Valid {
		p.Subrole = &subrole.String
	}
	if enterpriseID.Valid {
		eid, _ := uuid.Parse(enterpriseID.String)
		p.EnterpriseID = &eid
	}
	return p, nil
}

func (r *ProfileRepository) CreateCustomerProfile(ctx context.Context, userID uuid.UUID, fullName string, enterpriseID *uuid.UUID, position, phone *string) error {
	query := `
		INSERT INTO customer_profiles (user_id, full_name, enterprise_id, position, phone)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			enterprise_id = EXCLUDED.enterprise_id,
			position = EXCLUDED.position,
			phone = EXCLUDED.phone`

	_, err := r.db.Exec(ctx, query, userID, fullName, enterpriseID, position, phone)
	return err
}

func (r *ProfileRepository) UpdateCustomerProfile(ctx context.Context, userID uuid.UUID, fullName string, enterpriseID *uuid.UUID, position, phone *string) error {
	query := `
		UPDATE customer_profiles
		SET full_name = $2, enterprise_id = $3, position = $4, phone = $5
		WHERE user_id = $1`

	result, err := r.db.Exec(ctx, query, userID, fullName, enterpriseID, position, phone)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("profile not found")
	}
	return nil
}

func (r *ProfileRepository) DeleteCustomerProfile(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM customer_profiles WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *ProfileRepository) GetAllProfiles(ctx context.Context, profileType, search string, limit, offset int) ([]db.AllProfilesResult, int, error) {
	results := make([]db.AllProfilesResult, 0)

	baseQuery := `
		SELECT u.id, u.email, u.login, u.role, u.avatar,
			COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as full_name,
			COALESCE(bp.subrole, up.subrole, cp.subrole, NULL) as subrole,
			COALESCE(bp.position, sp.position, up.position, cp.position, NULL) as position,
			COALESCE(bp.phone, sp.phone, up.phone, cp.phone, NULL) as phone,
			sp.specialty,
			sp.grade,
			cp.enterprise_id,
			sp.university_department_id,
			up.university_department_id,
			COALESCE(
				TO_CHAR(bp.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
				TO_CHAR(sp.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
				TO_CHAR(up.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
				TO_CHAR(cp.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
				TO_CHAR(u.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
			) as updated_at
		FROM users u
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE u.role != 'admin'`

	countQuery := `
		SELECT COUNT(*) FROM users u
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE u.role != 'admin'`

	args := []interface{}{}
	argNum := 1

	if profileType != "" && profileType != "all" {
		switch profileType {
		case "brsm":
			baseQuery += ` AND (bp.user_id IS NOT NULL OR u.role = 'brsm')`
			countQuery += ` AND (bp.user_id IS NOT NULL OR u.role = 'brsm')`
		case "student":
			baseQuery += ` AND (sp.user_id IS NOT NULL OR u.role = 'student')`
			countQuery += ` AND (sp.user_id IS NOT NULL OR u.role = 'student')`
		case "university":
			baseQuery += ` AND (up.user_id IS NOT NULL OR u.role = 'university')`
			countQuery += ` AND (up.user_id IS NOT NULL OR u.role = 'university')`
		case "customer":
			baseQuery += ` AND (cp.user_id IS NOT NULL OR u.role = 'customer')`
			countQuery += ` AND (cp.user_id IS NOT NULL OR u.role = 'customer')`
		case "user":
			baseQuery += ` AND u.role = 'user'`
			countQuery += ` AND u.role = 'user'`
		}
	}

	if search != "" {
		searchPattern := "%" + search + "%"
		placeholder := "$" + strconv.Itoa(argNum)
		searchCond := `(u.login ILIKE ` + placeholder + ` OR u.email ILIKE ` + placeholder + ` OR COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') ILIKE ` + placeholder + `)`
		baseQuery += ` AND ` + searchCond
		countQuery += ` AND ` + searchCond
		args = append(args, searchPattern)
		argNum++
	}

	baseQuery += ` ORDER BY updated_at DESC LIMIT $` + strconv.Itoa(argNum) + ` OFFSET $` + strconv.Itoa(argNum+1)
	args = append(args, limit, offset)

	var total int
	err := r.db.QueryRow(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var row db.AllProfilesResult
		var enterpriseID, studentUnivDeptID, universityUnivDeptID sql.Null[uuid.UUID]
		var specialty sql.NullString
		var grade sql.Null[float64]

		err := rows.Scan(
			&row.UserID,
			&row.Email,
			&row.Login,
			&row.Role,
			&row.Avatar,
			&row.FullName,
			&row.Subrole,
			&row.Position,
			&row.Phone,
			&specialty,
			&grade,
			&enterpriseID,
			&studentUnivDeptID,
			&universityUnivDeptID,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if specialty.Valid {
			row.Specialty = &specialty.String
		}
		if grade.Valid {
			row.Grade = &grade.V
		}
		if enterpriseID.Valid {
			row.EnterpriseID = &enterpriseID.V
		}
		if studentUnivDeptID.Valid {
			row.StudentUniversityDeptID = &studentUnivDeptID.V
		}
		if universityUnivDeptID.Valid {
			row.UniversityUniversityDeptID = &universityUnivDeptID.V
		}

		results = append(results, row)
	}

	return results, total, nil
}

func (r *ProfileRepository) DeleteAllProfiles(ctx context.Context, userID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `DELETE FROM brsm_profiles WHERE user_id = $1`, userID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `DELETE FROM student_profiles WHERE user_id = $1`, userID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `DELETE FROM university_profiles WHERE user_id = $1`, userID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `DELETE FROM customer_profiles WHERE user_id = $1`, userID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `UPDATE users SET role = $1 WHERE id = $2`, "user", userID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
