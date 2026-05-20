package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type EmploymentRepository struct {
	db *pgxpool.Pool
}

func NewEmploymentRepository(db *pgxpool.Pool) *EmploymentRepository {
	return &EmploymentRepository{db: db}
}

func (r *EmploymentRepository) GetAll(ctx context.Context, status, universityDeptID string, limit, offset int) ([]db.EmploymentRequestWithDetails, int, error) {
	results := make([]db.EmploymentRequestWithDetails, 0)

	baseQuery := `
		SELECT er.id, er.enterprise_id, e.name as enterprise_name, e.address as enterprise_address,
			   er.university_department_id, COALESCE(u.name, '') as university_name, COALESCE(d.name, '') as department_name, ud.address as university_address,
			   er.title, er.description, er.requirements, er.salary, er.schedule,
			   er.max_participants,
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = er.id) as current_participants,
			   er.status_id, es.name as status_name,
			   er.created_at, er.updated_at
		FROM employment_requests er
		JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN university_departments ud ON er.university_department_id = ud.id
		LEFT JOIN universities u ON ud.university_id = u.id
		LEFT JOIN departments d ON ud.department_id = d.id
		JOIN employment_statuses es ON er.status_id = es.id`

	countQuery := `SELECT COUNT(*) FROM employment_requests er JOIN employment_statuses es ON er.status_id = es.id`

	args := []interface{}{}
	argNum := 1
	whereClause := ""

	if status != "" && status != "all" {
		whereClause = ` WHERE es.name = $` + strconv.Itoa(argNum)
		args = append(args, status)
		argNum++
	}

	if universityDeptID != "" {
		if whereClause == "" {
			whereClause = ` WHERE er.university_department_id = $` + strconv.Itoa(argNum)
		} else {
			whereClause += ` AND er.university_department_id = $` + strconv.Itoa(argNum)
		}
		args = append(args, universityDeptID)
		argNum++
	}

	baseQuery += whereClause + ` ORDER BY er.created_at DESC LIMIT $` + strconv.Itoa(argNum) + ` OFFSET $` + strconv.Itoa(argNum+1)
	countQuery += whereClause

	var total int
	if len(args) > 0 {
		err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := r.db.QueryRow(ctx, countQuery).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
	}

	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var e db.EmploymentRequestWithDetails
		err := rows.Scan(
			&e.ID, &e.EnterpriseID, &e.EnterpriseName, &e.EnterpriseAddress,
			&e.UniversityDepartmentID, &e.UniversityName, &e.DepartmentName, &e.UniversityAddress,
			&e.Title, &e.Description, &e.Requirements, &e.Salary, &e.Schedule,
			&e.MaxParticipants, &e.CurrentParticipants,
			&e.StatusID, &e.StatusName,
			&e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, e)
	}

	return results, total, nil
}

func (r *EmploymentRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.EmploymentRequestWithParticipants, error) {
	query := `
		SELECT er.id, er.enterprise_id, e.name as enterprise_name, e.address as enterprise_address,
			   er.university_department_id, COALESCE(u.name, '') as university_name, COALESCE(d.name, '') as department_name, ud.address as university_address,
			   er.title, er.description, er.requirements, er.salary, er.schedule,
			   er.max_participants,
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = er.id) as current_participants,
			   er.status_id, es.name as status_name,
			   er.created_at, er.updated_at
		FROM employment_requests er
		JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN university_departments ud ON er.university_department_id = ud.id
		LEFT JOIN universities u ON ud.university_id = u.id
		LEFT JOIN departments d ON ud.department_id = d.id
		JOIN employment_statuses es ON er.status_id = es.id
		WHERE er.id = $1`

	var e db.EmploymentRequestWithParticipants
	err := r.db.QueryRow(ctx, query, id).Scan(
		&e.ID, &e.EnterpriseID, &e.EnterpriseName, &e.EnterpriseAddress,
		&e.UniversityDepartmentID, &e.UniversityName, &e.DepartmentName, &e.UniversityAddress,
		&e.Title, &e.Description, &e.Requirements, &e.Salary, &e.Schedule,
		&e.MaxParticipants, &e.CurrentParticipants,
		&e.StatusID, &e.StatusName,
		&e.CreatedAt, &e.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	participantQuery := `SELECT user_id FROM employment_participants WHERE request_id = $1`
	rows, err := r.db.Query(ctx, participantQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e.ParticipantUserIDs = make([]uuid.UUID, 0)
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		e.ParticipantUserIDs = append(e.ParticipantUserIDs, userID)
	}

	return &e, nil
}

func (r *EmploymentRepository) Create(ctx context.Context, enterpriseID uuid.UUID, title string, description, requirements, salary, schedule *string, maxParticipants int) (*db.EmploymentRequest, error) {
	var id uuid.UUID

	query := `
		INSERT INTO employment_requests (enterprise_id, title, description, requirements, salary, schedule, max_participants)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, enterprise_id, university_department_id, title, description, requirements, salary, schedule, max_participants, status_id, created_at, updated_at`

	var req db.EmploymentRequest
	err := r.db.QueryRow(ctx, query, enterpriseID, title, description, requirements, salary, schedule, maxParticipants).Scan(
		&id, &req.EnterpriseID, &req.UniversityDepartmentID, &req.Title, &req.Description, &req.Requirements, &req.Salary, &req.Schedule, &req.MaxParticipants, &req.StatusID, &req.CreatedAt, &req.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	req.ID = id
	return &req, nil
}

func (r *EmploymentRepository) Update(ctx context.Context, id uuid.UUID, title, description, requirements, salary, schedule *string, maxParticipants *int, statusName *string) error {
	query := `UPDATE employment_requests SET updated_at = NOW()`
	args := []interface{}{}
	argNum := 1

	if title != nil {
		query += `, title = $` + strconv.Itoa(argNum)
		args = append(args, *title)
		argNum++
	}
	if description != nil {
		query += `, description = $` + strconv.Itoa(argNum)
		args = append(args, *description)
		argNum++
	}
	if requirements != nil {
		query += `, requirements = $` + strconv.Itoa(argNum)
		args = append(args, *requirements)
		argNum++
	}
	if salary != nil {
		query += `, salary = $` + strconv.Itoa(argNum)
		args = append(args, *salary)
		argNum++
	}
	if schedule != nil {
		query += `, schedule = $` + strconv.Itoa(argNum)
		args = append(args, *schedule)
		argNum++
	}
	if maxParticipants != nil {
		query += `, max_participants = $` + strconv.Itoa(argNum)
		args = append(args, *maxParticipants)
		argNum++
	}
	if statusName != nil {
		query += `, status_id = (SELECT id FROM employment_statuses WHERE name = $` + strconv.Itoa(argNum) + `)`
		args = append(args, *statusName)
		argNum++
	}

	query += ` WHERE id = $` + strconv.Itoa(argNum)
	args = append(args, id)

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *EmploymentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employment_requests WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *EmploymentRepository) Apply(ctx context.Context, requestID, userID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var maxParticipants, currentCount, statusID int
	err = tx.QueryRow(ctx, `
		SELECT max_participants, 
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = $1),
			   status_id
		FROM employment_requests WHERE id = $1
	`, requestID).Scan(&maxParticipants, &currentCount, &statusID)
	if err != nil {
		return err
	}

	if statusID != 2 {
		return errors.New("applications are only accepted for approved requests")
	}

	if currentCount >= maxParticipants {
		return errors.New("all positions are filled")
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO employment_participants (request_id, user_id)
		VALUES ($1, $2)
		ON CONFLICT (request_id, user_id) DO NOTHING
	`, requestID, userID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *EmploymentRepository) UpdateParticipantStatus(ctx context.Context, requestID, userID uuid.UUID, statusName string) error {
	var statusID int
	err := r.db.QueryRow(ctx, `SELECT id FROM employment_participant_statuses WHERE name = $1`, statusName).Scan(&statusID)
	if err != nil {
		if err == pgx.ErrNoRows {
			r.ensureParticipantStatusesExist(ctx)
			err = r.db.QueryRow(ctx, `SELECT id FROM employment_participant_statuses WHERE name = $1`, statusName).Scan(&statusID)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	var updateQuery string
	var args []interface{}

	if statusName == "contracted" {
		updateQuery = `UPDATE employment_participants SET status_id = $1, contracted_at = NOW() WHERE request_id = $2 AND user_id = $3`
		args = []interface{}{statusID, requestID, userID}
	} else {
		updateQuery = `UPDATE employment_participants SET status_id = $1 WHERE request_id = $2 AND user_id = $3`
		args = []interface{}{statusID, requestID, userID}
	}

	_, err = r.db.Exec(ctx, updateQuery, args...)
	return err
}

func (r *EmploymentRepository) CancelApplication(ctx context.Context, requestID, userID uuid.UUID) error {
	query := `DELETE FROM employment_participants WHERE request_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, requestID, userID)
	return err
}

func (r *EmploymentRepository) GetParticipants(ctx context.Context, requestID uuid.UUID) ([]db.EmploymentParticipantWithDetails, error) {
	query := `
		SELECT ep.id, ep.request_id, ep.user_id,
			   COALESCE(sp.full_name, up.full_name, bp.full_name, '') as student_name,
			   ep.status_id, eps.name as status_name,
			   ep.applied_at, ep.contracted_at
		FROM employment_participants ep
		JOIN users u ON ep.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = ep.user_id
		LEFT JOIN university_profiles up ON up.user_id = ep.user_id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ep.user_id
		JOIN employment_participant_statuses eps ON ep.status_id = eps.id
		WHERE ep.request_id = $1
		ORDER BY ep.applied_at DESC`

	rows, err := r.db.Query(ctx, query, requestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []db.EmploymentParticipantWithDetails
	for rows.Next() {
		var p db.EmploymentParticipantWithDetails
		err := rows.Scan(&p.ID, &p.RequestID, &p.UserID, &p.StudentName, &p.StatusID, &p.StatusName, &p.AppliedAt, &p.ContractedAt)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}
	return participants, nil
}

func (r *EmploymentRepository) GetStatuses(ctx context.Context) ([]db.EmploymentStatus, error) {
	query := `SELECT id, name, description FROM employment_statuses ORDER BY id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []db.EmploymentStatus
	for rows.Next() {
		var s db.EmploymentStatus
		if err := rows.Scan(&s.ID, &s.Name, &s.Description); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}

func (r *EmploymentRepository) GetParticipantStatuses(ctx context.Context) ([]db.EmploymentParticipantStatus, error) {
	query := `SELECT id, name, description FROM employment_participant_statuses ORDER BY id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []db.EmploymentParticipantStatus
	for rows.Next() {
		var s db.EmploymentParticipantStatus
		if err := rows.Scan(&s.ID, &s.Name, &s.Description); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}

func (r *EmploymentRepository) GetByEnterprise(ctx context.Context, enterpriseID uuid.UUID) ([]db.EmploymentRequestWithDetails, error) {
	query := `
		SELECT er.id, er.enterprise_id, e.name as enterprise_name, e.address as enterprise_address,
			   er.university_department_id, COALESCE(u.name, '') as university_name, COALESCE(d.name, '') as department_name, ud.address as university_address,
			   er.title, er.description, er.requirements, er.salary, er.schedule,
			   er.max_participants,
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = er.id) as current_participants,
			   er.status_id, es.name as status_name,
			   er.created_at, er.updated_at
		FROM employment_requests er
		JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN university_departments ud ON er.university_department_id = ud.id
		LEFT JOIN universities u ON ud.university_id = u.id
		LEFT JOIN departments d ON ud.department_id = d.id
		JOIN employment_statuses es ON er.status_id = es.id
		WHERE er.enterprise_id = $1
		ORDER BY er.created_at DESC`

	rows, err := r.db.Query(ctx, query, enterpriseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []db.EmploymentRequestWithDetails
	for rows.Next() {
		var e db.EmploymentRequestWithDetails
		err := rows.Scan(
			&e.ID, &e.EnterpriseID, &e.EnterpriseName, &e.EnterpriseAddress,
			&e.UniversityDepartmentID, &e.UniversityName, &e.DepartmentName, &e.UniversityAddress,
			&e.Title, &e.Description, &e.Requirements, &e.Salary, &e.Schedule,
			&e.MaxParticipants, &e.CurrentParticipants,
			&e.StatusID, &e.StatusName,
			&e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, e)
	}
	return results, nil
}

func (r *EmploymentRepository) GetByUniversityDepartment(ctx context.Context, universityDeptID uuid.UUID) ([]db.EmploymentRequestWithDetails, error) {
	query := `
		SELECT er.id, er.enterprise_id, e.name as enterprise_name, e.address as enterprise_address,
			   er.university_department_id, COALESCE(u.name, '') as university_name, COALESCE(d.name, '') as department_name, ud.address as university_address,
			   er.title, er.description, er.requirements, er.salary, er.schedule,
			   er.max_participants,
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = er.id) as current_participants,
			   er.status_id, es.name as status_name,
			   er.created_at, er.updated_at
		FROM employment_requests er
		JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN university_departments ud ON er.university_department_id = ud.id
		LEFT JOIN universities u ON ud.university_id = u.id
		LEFT JOIN departments d ON ud.department_id = d.id
		JOIN employment_statuses es ON er.status_id = es.id
		WHERE er.university_department_id = $1
		ORDER BY er.created_at DESC`

	rows, err := r.db.Query(ctx, query, universityDeptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []db.EmploymentRequestWithDetails
	for rows.Next() {
		var e db.EmploymentRequestWithDetails
		err := rows.Scan(
			&e.ID, &e.EnterpriseID, &e.EnterpriseName, &e.EnterpriseAddress,
			&e.UniversityDepartmentID, &e.UniversityName, &e.DepartmentName, &e.UniversityAddress,
			&e.Title, &e.Description, &e.Requirements, &e.Salary, &e.Schedule,
			&e.MaxParticipants, &e.CurrentParticipants,
			&e.StatusID, &e.StatusName,
			&e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, e)
	}
	return results, nil
}

func (r *EmploymentRepository) IsUserApplied(ctx context.Context, requestID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM employment_participants WHERE request_id = $1 AND user_id = $2)`, requestID, userID).Scan(&exists)
	return exists, err
}

func (r *EmploymentRepository) UpdateStatus(ctx context.Context, id uuid.UUID, approved bool) error {
	var statusName string
	if approved {
		statusName = "approved"
	} else {
		statusName = "rejected"
	}

	query := `UPDATE employment_requests SET status_id = (SELECT id FROM employment_statuses WHERE name = $1), updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(ctx, query, statusName, id)
	return err
}

func (r *EmploymentRepository) GetUserApplications(ctx context.Context, userID uuid.UUID) ([]db.EmploymentApplicationWithDetails, error) {
	query := `
		SELECT er.id, er.enterprise_id, e.name as enterprise_name, e.address as enterprise_address,
			   er.university_department_id, COALESCE(u.name, '') as university_name, COALESCE(d.name, '') as department_name, ud.address as university_address,
			   er.title, er.description, er.requirements, er.salary, er.schedule,
			   er.max_participants,
			   (SELECT COUNT(*) FROM employment_participants WHERE request_id = er.id) as current_participants,
			   er.status_id, es.name as status_name,
			   er.created_at, er.updated_at,
			   ep.status_id as participant_status_id, eps.name as participant_status_name, ep.applied_at as participant_applied_at
		FROM employment_requests er
		JOIN enterprises e ON er.enterprise_id = e.id
		LEFT JOIN university_departments ud ON er.university_department_id = ud.id
		LEFT JOIN universities u ON ud.university_id = u.id
		LEFT JOIN departments d ON ud.department_id = d.id
		JOIN employment_statuses es ON er.status_id = es.id
		JOIN employment_participants ep ON ep.request_id = er.id
		JOIN employment_participant_statuses eps ON ep.status_id = eps.id
		WHERE ep.user_id = $1
		ORDER BY ep.applied_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]db.EmploymentApplicationWithDetails, 0)
	for rows.Next() {
		var a db.EmploymentApplicationWithDetails
		err := rows.Scan(
			&a.ID, &a.EnterpriseID, &a.EnterpriseName, &a.EnterpriseAddress,
			&a.UniversityDepartmentID, &a.UniversityName, &a.DepartmentName, &a.UniversityAddress,
			&a.Title, &a.Description, &a.Requirements, &a.Salary, &a.Schedule,
			&a.MaxParticipants, &a.CurrentParticipants,
			&a.StatusID, &a.StatusName,
			&a.CreatedAt, &a.UpdatedAt,
			&a.ParticipantStatusID, &a.ParticipantStatusName, &a.ParticipantAppliedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, a)
	}
	return results, nil
}

func (r *EmploymentRepository) ensureEmploymentStatusesExist(ctx context.Context) {
	r.db.Exec(ctx, `
		INSERT INTO employment_statuses (name, description) VALUES 
			('pending', 'Ожидает подтверждения'),
			('approved', 'Подтверждена'),
			('closed', 'Закрыта'),
			('rejected', 'Отклонена')
		ON CONFLICT DO NOTHING
	`)
}

func (r *EmploymentRepository) ensureParticipantStatusesExist(ctx context.Context) {
	r.db.Exec(ctx, `
		INSERT INTO employment_participant_statuses (name, description) VALUES 
			('applied', 'Подана'),
			('accepted', 'Принята'),
			('rejected', 'Отклонена'),
			('contracted', 'Трудоустроен')
		ON CONFLICT DO NOTHING
	`)
}
