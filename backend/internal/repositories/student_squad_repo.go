package repositories

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type StudentSquadRepository struct {
	db *pgxpool.Pool
}

func NewStudentSquadRepository(db *pgxpool.Pool) *StudentSquadRepository {
	return &StudentSquadRepository{db: db}
}

func (r *StudentSquadRepository) GetAll(ctx context.Context, status string, limit, offset int) ([]db.StudentSquadWithDetails, int, error) {
	results := make([]db.StudentSquadWithDetails, 0)

	baseQuery := `
		SELECT ss.id, ss.organizer_id,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as organizer_name,
			   u.role as organizer_role,
			   bp.position as organizer_position,
			   bp.phone as organizer_phone,
			   ss.title, ss.description, ss.profile,
			   ss.max_participants,
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = ss.id) as current_count,
			   ss.status_id, ss2.name as status_name,
			   ss.created_at, ss.updated_at
		FROM student_squads ss
		JOIN users u ON ss.organizer_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ss.organizer_id
		LEFT JOIN student_profiles sp ON sp.user_id = ss.organizer_id
		LEFT JOIN university_profiles up ON up.user_id = ss.organizer_id
		LEFT JOIN customer_profiles cp ON cp.user_id = ss.organizer_id
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id`

	countQuery := `SELECT COUNT(*) FROM student_squads ss JOIN squad_statuses ss2 ON ss.status_id = ss2.id`

	args := []interface{}{}
	argNum := 1

	if status != "" && status != "all" {
		baseQuery += ` WHERE ss2.name = $` + strconv.Itoa(argNum)
		countQuery += ` WHERE ss2.name = $` + strconv.Itoa(argNum)
		args = append(args, status)
		argNum++
	}

	baseQuery += ` ORDER BY ss.created_at DESC LIMIT $` + strconv.Itoa(argNum) + ` OFFSET $` + strconv.Itoa(argNum+1)

	var total int
	countArgs := args
	if len(countArgs) > 0 {
		err := r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
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
		var s db.StudentSquadWithDetails
		err := rows.Scan(
			&s.ID, &s.OrganizerID, &s.OrganizerName, &s.OrganizerRole,
			&s.OrganizerPosition, &s.OrganizerPhone,
			&s.Title, &s.Description, &s.Profile, &s.MaxParticipants,
			&s.CurrentCount, &s.StatusID, &s.StatusName,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, s)
	}

	return results, total, nil
}

func (r *StudentSquadRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.StudentSquadWithParticipants, error) {
	query := `
		SELECT ss.id, ss.organizer_id,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as organizer_name,
			   u.role as organizer_role,
			   ss.title, ss.description, ss.profile,
			   ss.max_participants,
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = ss.id) as current_count,
			   ss.status_id, ss2.name as status_name,
			   ss.created_at, ss.updated_at
		FROM student_squads ss
		JOIN users u ON ss.organizer_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ss.organizer_id
		LEFT JOIN student_profiles sp ON sp.user_id = ss.organizer_id
		LEFT JOIN university_profiles up ON up.user_id = ss.organizer_id
		LEFT JOIN customer_profiles cp ON cp.user_id = ss.organizer_id
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id
		WHERE ss.id = $1`

	var s db.StudentSquadWithParticipants
	err := r.db.QueryRow(ctx, query, id).Scan(
		&s.ID, &s.OrganizerID, &s.OrganizerName, &s.OrganizerRole,
		&s.Title, &s.Description, &s.Profile, &s.MaxParticipants,
		&s.CurrentCount, &s.StatusID, &s.StatusName,
		&s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	participantQuery := `SELECT user_id FROM student_squad_participants WHERE squad_id = $1`
	rows, err := r.db.Query(ctx, participantQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	s.Participants = make([]uuid.UUID, 0)
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		s.Participants = append(s.Participants, userID)
	}

	return &s, nil
}

func (r *StudentSquadRepository) Create(ctx context.Context, organizerID uuid.UUID, title string, description, profile *string, maxParticipants int, statusName string) (*db.StudentSquad, error) {
	var statusID int
	err := r.db.QueryRow(ctx, `SELECT id FROM squad_statuses WHERE name = $1`, statusName).Scan(&statusID)
	if err != nil {
		if err == pgx.ErrNoRows {
			r.ensureStatusesExist(ctx)
			err = r.db.QueryRow(ctx, `SELECT id FROM squad_statuses WHERE name = $1`, statusName).Scan(&statusID)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	

	query := `
		INSERT INTO student_squads (organizer_id, title, description, profile, max_participants, status_id, approved_by_university_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, organizer_id, title, description, profile, max_participants, status_id, created_at, updated_at`

	var id uuid.UUID
	var createdAt, updatedAt time.Time
	err = r.db.QueryRow(ctx, query, organizerID, title, description, profile, maxParticipants, statusID, nil).Scan(
		&id, &organizerID, &title, &description, &profile, &maxParticipants, &statusID, &createdAt, &updatedAt,
	)
	if err != nil {
		return nil, err
	}

	squad := &db.StudentSquad{
		ID:              id,
		OrganizerID:     organizerID,
		Title:           title,
		Description:     description,
		Profile:         profile,
		MaxParticipants: maxParticipants,
		StatusID:        statusID,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	return squad, nil
}

func (r *StudentSquadRepository) Update(ctx context.Context, id uuid.UUID, title *string, description, profile *string, maxParticipants *int, statusName *string) error {
	query := `UPDATE student_squads SET updated_at = NOW()`
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
	if profile != nil {
		query += `, profile = $` + strconv.Itoa(argNum)
		args = append(args, *profile)
		argNum++
	}
	if maxParticipants != nil {
		query += `, max_participants = $` + strconv.Itoa(argNum)
		args = append(args, *maxParticipants)
		argNum++
	}
	if statusName != nil {
		query += `, status_id = (SELECT id FROM squad_statuses WHERE name = $` + strconv.Itoa(argNum) + `)`
		args = append(args, *statusName)
		argNum++
	}

	query += ` WHERE id = $` + strconv.Itoa(argNum)
	args = append(args, id)

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *StudentSquadRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM student_squads WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *StudentSquadRepository) Join(ctx context.Context, squadID, userID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var maxParticipants, currentCount int
	err = tx.QueryRow(ctx, `
		SELECT max_participants, 
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = $1)
		FROM student_squads WHERE id = $1
	`, squadID).Scan(&maxParticipants, &currentCount)
	if err != nil {
		return err
	}

	if currentCount >= maxParticipants {
		return errors.New("squad is full")
	}

	var statusName string
	err = tx.QueryRow(ctx, `
		SELECT ss2.name FROM student_squads ss
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id
		WHERE ss.id = $1
	`, squadID).Scan(&statusName)
	if err != nil {
		return err
	}
	if statusName != "recruitment_open" {
		return errors.New("recruitment is closed")
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO student_squad_participants (squad_id, user_id)
		VALUES ($1, $2)
		ON CONFLICT (squad_id, user_id) DO NOTHING
	`, squadID, userID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *StudentSquadRepository) Leave(ctx context.Context, squadID, userID uuid.UUID) error {
	query := `DELETE FROM student_squad_participants WHERE squad_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, squadID, userID)
	return err
}

func (r *StudentSquadRepository) GetParticipantCount(ctx context.Context, squadID uuid.UUID) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = $1`, squadID).Scan(&count)
	return count, err
}

func (r *StudentSquadRepository) IsUserParticipant(ctx context.Context, squadID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM student_squad_participants WHERE squad_id = $1 AND user_id = $2)`, squadID, userID).Scan(&exists)
	return exists, err
}

func (r *StudentSquadRepository) GetStatuses(ctx context.Context) ([]db.SquadStatus, error) {
	query := `SELECT id, name, description FROM squad_statuses ORDER BY id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []db.SquadStatus
	for rows.Next() {
		var s db.SquadStatus
		if err := rows.Scan(&s.ID, &s.Name, &s.Description); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}
	return statuses, nil
}

func (r *StudentSquadRepository) ensureStatusesExist(ctx context.Context) {
	r.db.Exec(ctx, `
		INSERT INTO squad_statuses (name, description) VALUES 
			('pending', 'Ожидает подтверждения'),
			('recruitment_open', 'Идёт набор'),
			('rejected', 'Отклонена'),
			('closed', 'Закрыта')
		ON CONFLICT (name) DO NOTHING
	`)
}

func (r *StudentSquadRepository) GetByOrganizer(ctx context.Context, organizerID uuid.UUID) ([]db.StudentSquadWithDetails, error) {
	query := `
		SELECT ss.id, ss.organizer_id,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as organizer_name,
			   u.role as organizer_role,
			   bp.position as organizer_position,
			   bp.phone as organizer_phone,
			   ss.title, ss.description, ss.profile,
			   ss.max_participants,
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = ss.id) as current_count,
			   ss.status_id, ss2.name as status_name,
			   ss.created_at, ss.updated_at
		FROM student_squads ss
		JOIN users u ON ss.organizer_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ss.organizer_id
		LEFT JOIN student_profiles sp ON sp.user_id = ss.organizer_id
		LEFT JOIN university_profiles up ON up.user_id = ss.organizer_id
		LEFT JOIN customer_profiles cp ON cp.user_id = ss.organizer_id
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id
		WHERE ss.organizer_id = $1
		ORDER BY ss.created_at DESC`

	rows, err := r.db.Query(ctx, query, organizerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []db.StudentSquadWithDetails
	for rows.Next() {
		var s db.StudentSquadWithDetails
		err := rows.Scan(
			&s.ID, &s.OrganizerID, &s.OrganizerName, &s.OrganizerRole, &s.OrganizerPosition, &s.OrganizerPhone,
			&s.Title, &s.Description, &s.Profile, &s.MaxParticipants,
			&s.CurrentCount, &s.StatusID, &s.StatusName,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	return results, nil
}

func (r *StudentSquadRepository) GetByParticipant(ctx context.Context, userID uuid.UUID) ([]db.StudentSquadWithDetails, error) {
	query := `
		SELECT ss.id, ss.organizer_id,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as organizer_name,
			   u.role as organizer_role,
			   bp.position as organizer_position,
			   bp.phone as organizer_phone,
			   ss.title, ss.description, ss.profile,
			   ss.max_participants,
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = ss.id) as current_count,
			   ss.status_id, ss2.name as status_name,
			   ss.created_at, ss.updated_at
		FROM student_squads ss
		JOIN users u ON ss.organizer_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ss.organizer_id
		LEFT JOIN student_profiles sp ON sp.user_id = ss.organizer_id
		LEFT JOIN university_profiles up ON up.user_id = ss.organizer_id
		LEFT JOIN customer_profiles cp ON cp.user_id = ss.organizer_id
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id
		JOIN student_squad_participants ssp ON ssp.squad_id = ss.id
		WHERE ssp.user_id = $1
		ORDER BY ss.created_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []db.StudentSquadWithDetails
	for rows.Next() {
		var s db.StudentSquadWithDetails
		err := rows.Scan(
			&s.ID, &s.OrganizerID, &s.OrganizerName, &s.OrganizerRole,
			&s.OrganizerPosition, &s.OrganizerPhone,
			&s.Title, &s.Description, &s.Profile, &s.MaxParticipants,
			&s.CurrentCount, &s.StatusID, &s.StatusName,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	return results, nil
}

func (r *StudentSquadRepository) UpdateStatus(ctx context.Context, squadID uuid.UUID, statusName string, approvedByUserID uuid.UUID) error {
	query := `UPDATE student_squads SET status_id = (SELECT id FROM squad_statuses WHERE name = $1), approved_by_university_id = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.Exec(ctx, query, statusName, approvedByUserID, squadID)
	return err
}

func (r *StudentSquadRepository) GetByUniversity(ctx context.Context, universityDeptID uuid.UUID) ([]db.StudentSquadWithDetails, error) {
	query := `
		SELECT ss.id, ss.organizer_id,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name, '') as organizer_name,
			   u.role as organizer_role,
			   bp.position as organizer_position,
			   bp.phone as organizer_phone,
			   ss.title, ss.description, ss.profile,
			   ss.max_participants,
			   (SELECT COUNT(*) FROM student_squad_participants WHERE squad_id = ss.id) as current_count,
			   ss.status_id, ss2.name as status_name,
			   ss.created_at, ss.updated_at,
			   u.avatar
		FROM student_squads ss
		JOIN users u ON ss.organizer_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = ss.organizer_id
		LEFT JOIN student_profiles sp ON sp.user_id = ss.organizer_id
		LEFT JOIN university_profiles up ON up.user_id = ss.organizer_id
		LEFT JOIN customer_profiles cp ON cp.user_id = ss.organizer_id
		JOIN squad_statuses ss2 ON ss.status_id = ss2.id
		WHERE EXISTS (
			SELECT 1 FROM university_profiles up3
			JOIN university_departments ud3 ON ud3.id = up3.university_department_id
			WHERE up3.user_id = ss.approved_by_university_id
			AND ud3.university_id = (SELECT university_id FROM university_departments WHERE id = $1)
		) OR EXISTS (
			SELECT 1 FROM student_profiles sp 
			JOIN university_departments ud4 ON ud4.id = sp.university_department_id
			WHERE sp.user_id = ss.organizer_id
			AND ud4.university_id = (SELECT university_id FROM university_departments WHERE id = $1)
		) OR EXISTS (
			SELECT 1 FROM university_profiles up2
			JOIN university_departments ud2 ON ud2.id = up2.university_department_id
			WHERE up2.user_id = ss.organizer_id
			AND ud2.university_id = (SELECT university_id FROM university_departments WHERE id = $1)
		) OR EXISTS (
			SELECT 1 FROM student_squad_participants ssp
			JOIN student_profiles sp2 ON ssp.user_id = sp2.user_id
			JOIN university_departments ud5 ON ud5.id = sp2.university_department_id
			WHERE ssp.squad_id = ss.id
			AND ud5.university_id = (SELECT university_id FROM university_departments WHERE id = $1)
		)
		ORDER BY ss.created_at DESC`

	rows, err := r.db.Query(ctx, query, universityDeptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []db.StudentSquadWithDetails
	for rows.Next() {
		var s db.StudentSquadWithDetails
		err := rows.Scan(
			&s.ID, &s.OrganizerID, &s.OrganizerName, &s.OrganizerRole, &s.OrganizerPosition, &s.OrganizerPhone,
			&s.Title, &s.Description, &s.Profile, &s.MaxParticipants,
			&s.CurrentCount, &s.StatusID, &s.StatusName,
			&s.CreatedAt, &s.UpdatedAt, &s.Avatar,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	
	return results, nil
}

func (r *StudentSquadRepository) GetParticipantsBySquadIDs(ctx context.Context, squadIDs []uuid.UUID) (map[uuid.UUID][]db.SquadParticipantInfo, error) {
	if len(squadIDs) == 0 {
		return make(map[uuid.UUID][]db.SquadParticipantInfo), nil
	}

	query := `
		SELECT ssp.squad_id, p.id as user_id, sp.full_name, sp.phone, sp.specialty, sp.grade, p.avatar
		FROM student_squad_participants ssp
		JOIN student_profiles sp ON ssp.user_id = sp.user_id
		JOIN users p ON p.id = ssp.user_id
		WHERE ssp.squad_id = ANY($1)
	`

	rows, err := r.db.Query(ctx, query, squadIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[uuid.UUID][]db.SquadParticipantInfo)
	for rows.Next() {
		var info db.SquadParticipantInfo
		err := rows.Scan(&info.SquadID, &info.UserID, &info.FullName, &info.Phone, &info.Specialty, &info.Grade, &info.Avatar)
		if err != nil {
			return nil, err
		}
		result[info.SquadID] = append(result[info.SquadID], info)
	}

	return result, nil
}
