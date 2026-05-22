package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
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

type UniversityStats struct {
	TotalStudents    int
	StudentsInSquads int
	StudentsEmployed int
	TotalSquads      int
	JobInvitations   int
}

func (r *ProfileRepository) GetUniversityStats(ctx context.Context, universityDeptID uuid.UUID) (*UniversityStats, error) {
	stats := &UniversityStats{}

	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*) FROM student_profiles WHERE university_department_id = $1
	`, universityDeptID).Scan(&stats.TotalStudents)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, `
		SELECT COUNT(DISTINCT ssp.user_id)
		FROM student_squad_participants ssp
		JOIN student_squads squads ON ssp.squad_id = squads.id
		JOIN student_profiles sp ON ssp.user_id = sp.user_id
		WHERE sp.university_department_id = $1
	`, universityDeptID).Scan(&stats.StudentsInSquads)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, `
		SELECT COUNT(DISTINCT ep.user_id)
		FROM employment_participants ep
		JOIN employment_requests er ON ep.request_id = er.id
		JOIN student_profiles sp ON ep.user_id = sp.user_id
		WHERE er.university_department_id = $1 AND ep.status_id >= 2
	`, universityDeptID).Scan(&stats.StudentsEmployed)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, `
		SELECT COUNT(DISTINCT ss.id)
		FROM student_squads ss
		LEFT JOIN student_profiles sp ON ss.organizer_id = sp.user_id AND sp.university_department_id = $1
		JOIN squad_statuses squad_s ON ss.status_id = squad_s.id
		WHERE squad_s.name = 'recruitment_open'
		  AND (sp.user_id IS NOT NULL
			   OR EXISTS (
				   SELECT 1 FROM university_profiles up
				   WHERE up.user_id = ss.approved_by_university_id
				   AND up.university_department_id = $1
			   ))
	`, universityDeptID).Scan(&stats.TotalSquads)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM employment_requests er
		WHERE er.university_department_id = $1 AND er.status_id = 2
	`, universityDeptID).Scan(&stats.JobInvitations)
	if err != nil {
		return nil, err
	}

	return stats, nil
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

func (r *ProfileRepository) SearchUsers(ctx context.Context, search string, limit int) (*dto.AllProfilesResponse, error) {
	searchPattern := "%" + search + "%"

	query := `
		SELECT u.id, u.email, u.login, u.role, u.avatar,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as full_name,
			   COALESCE(bp.position, sp.position, up.position, cp.position, '') as position,
			   COALESCE(bp.phone, sp.phone, up.phone, cp.phone, '') as phone,
			   TO_CHAR(COALESCE(bp.updated_at, sp.updated_at, up.updated_at, cp.updated_at, u.updated_at), 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as updated_at
		FROM users u
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE u.role != 'admin'
		  AND (u.login ILIKE $1 OR u.email ILIKE $1 OR COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') ILIKE $1)
		ORDER BY updated_at DESC
		LIMIT $2`

	rows, err := r.db.Query(ctx, query, searchPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []dto.ProfileResponse
	for rows.Next() {
		var p dto.ProfileResponse
		var fullName, position, phone sql.NullString

		err := rows.Scan(
			&p.UserID,
			&p.Email,
			&p.Login,
			&p.Role,
			&p.Avatar,
			&fullName,
			&position,
			&phone,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		profileData := make(map[string]interface{})
		if fullName.Valid {
			profileData["full_name"] = fullName.String
		}
		if position.Valid {
			profileData["position"] = position.String
		}
		if phone.Valid {
			profileData["phone"] = phone.String
		}
		p.Profile = profileData

		profiles = append(profiles, p)
	}

	if profiles == nil {
		profiles = []dto.ProfileResponse{}
	}

	return &dto.AllProfilesResponse{
		Profiles: profiles,
		Total:    len(profiles),
		Limit:    limit,
		Offset:   0,
	}, nil
}

func (r *ProfileRepository) GetStudentsWithEmployment(ctx context.Context, universityDeptID uuid.UUID, limit, offset int) (*dto.StudentEmploymentListResponse, error) {
	countQuery := `SELECT COUNT(*) FROM student_profiles WHERE university_department_id = $1`
	var total int
	if err := r.db.QueryRow(ctx, countQuery, universityDeptID).Scan(&total); err != nil {
		return nil, err
	}

	studentsQuery := `
		SELECT u.id, u.email, u.login, u.role, u.avatar,
			   sp.full_name, sp.specialty, sp.grade, sp.position, sp.phone,
			   TO_CHAR(sp.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as updated_at
		FROM student_profiles sp
		JOIN users u ON sp.user_id = u.id
		WHERE sp.university_department_id = $1
		ORDER BY sp.full_name ASC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(ctx, studentsQuery, universityDeptID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []dto.StudentEmploymentInfo
	for rows.Next() {
		var p dto.ProfileResponse
		var fullName, specialty, position, phone sql.NullString
		var grade sql.Null[float64]

		err := rows.Scan(
			&p.UserID, &p.Email, &p.Login, &p.Role, &p.Avatar,
			&fullName, &specialty, &grade, &position, &phone,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		profileData := make(map[string]interface{})
		if fullName.Valid {
			profileData["full_name"] = fullName.String
		}
		if specialty.Valid {
			profileData["specialty"] = specialty.String
		}
		if grade.Valid {
			profileData["grade"] = grade.V
		}
		if position.Valid {
			profileData["position"] = position.String
		}
		if phone.Valid {
			profileData["phone"] = phone.String
		}
		p.Profile = profileData

		studentUUID, _ := uuid.Parse(p.UserID)

		squads, errSquads := r.GetStudentSquads(ctx, studentUUID)
		if errSquads != nil {
			log.Printf("Error getting squads for user %s: %v", p.UserID, errSquads)
		}
		works, errWorks := r.GetStudentWorks(ctx, studentUUID, universityDeptID)
		if errWorks != nil {
			log.Printf("Error getting works for user %s: %v", p.UserID, errWorks)
		}

		students = append(students, dto.StudentEmploymentInfo{
			Student:       p,
			StudentSquads: squads,
			Works:         works,
		})
	}

	if students == nil {
		students = []dto.StudentEmploymentInfo{}
	}

	return &dto.StudentEmploymentListResponse{
		Students: students,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

func (r *ProfileRepository) GetStudentSquads(ctx context.Context, userID uuid.UUID) ([]dto.StudentSquadInfo, error) {
	query := `
		SELECT ss.id, ss.title, ss.description,
			   TO_CHAR(ssp.joined_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as joined_at
		FROM student_squad_participants ssp
		JOIN student_squads ss ON ssp.squad_id = ss.id
		LEFT JOIN squad_statuses ss_status ON ss.status_id = ss_status.id
		WHERE ssp.user_id = $1 AND ss_status.name = 'recruitment_open'
		ORDER BY ssp.joined_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		log.Printf("GetStudentSquads query error: %v", err)
		return []dto.StudentSquadInfo{}, nil
	}
	defer rows.Close()

	var squads []dto.StudentSquadInfo
	for rows.Next() {
		var s dto.StudentSquadInfo
		var desc sql.NullString

		err := rows.Scan(&s.ID, &s.Name, &desc, &s.JoinedAt)
		if err != nil {
			log.Printf("GetStudentSquads scan error: %v", err)
			continue
		}
		if desc.Valid {
			s.Description = &desc.String
		}
		squads = append(squads, s)
	}

	if squads == nil {
		squads = []dto.StudentSquadInfo{}
	}
	return squads, nil
}

func (r *ProfileRepository) GetStudentWorks(ctx context.Context, userID, universityDeptID uuid.UUID) ([]dto.WorkInfo, error) {
	query := `
		SELECT er.id, er.title, COALESCE(e.name, 'Unknown') as company, '' as position,
			   COALESCE(eps.name, 'unknown') as status,
			   TO_CHAR(ep.applied_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') as started_at
		FROM employment_participants ep
		JOIN employment_requests er ON ep.request_id = er.id
		LEFT JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN employment_participant_statuses eps ON ep.status_id = eps.id
		WHERE ep.user_id = $1 AND er.university_department_id = $2 AND ep.status_id >= 2
		ORDER BY ep.applied_at DESC`

	rows, err := r.db.Query(ctx, query, userID, universityDeptID)
	if err != nil {
		log.Printf("GetStudentWorks query error: %v", err)
		return []dto.WorkInfo{}, nil
	}
	defer rows.Close()

	var works []dto.WorkInfo
	for rows.Next() {
		var w dto.WorkInfo

		err := rows.Scan(&w.ID, &w.Title, &w.Company, &w.Position, &w.Status, &w.StartedAt)
		if err != nil {
			log.Printf("GetStudentWorks scan error: %v", err)
			continue
		}
		works = append(works, w)
	}

	if works == nil {
		works = []dto.WorkInfo{}
	}
	return works, nil
}
