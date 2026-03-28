package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lanxre/frameby/internal/models/db"
)

type ChatRepository struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) CreateChat(ctx context.Context, chatType string, name *string, createdBy uuid.UUID) (*db.Chat, error) {
	query := `
		INSERT INTO chats (type, name, created_by)
		VALUES ($1, $2, $3)
		RETURNING id, type, name, avatar, created_by, created_at, updated_at`

	var chat db.Chat
	var namePtr sql.NullString
	var avatarPtr sql.NullString
	var createdByPtr sql.NullString

	err := r.db.QueryRow(ctx, query, chatType, name, createdBy).Scan(
		&chat.ID, &chat.Type, &namePtr, &avatarPtr, &createdByPtr, &chat.CreatedAt, &chat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if namePtr.Valid {
		chat.Name = &namePtr.String
	}
	if avatarPtr.Valid {
		chat.Avatar = &avatarPtr.String
	}
	if createdByPtr.Valid {
		chat.CreatedBy = &createdByPtr.String
	}

	return &chat, nil
}

func (r *ChatRepository) AddMember(ctx context.Context, chatID, userID uuid.UUID, role string) error {
	query := `
		INSERT INTO chat_members (chat_id, user_id, role, unread_count)
		VALUES ($1, $2, $3, 0)
		ON CONFLICT (chat_id, user_id) DO NOTHING`
	_, err := r.db.Exec(ctx, query, chatID, userID, role)
	return err
}

func (r *ChatRepository) MarkChatAsRead(ctx context.Context, chatID, userID uuid.UUID) error {
	query := `
		UPDATE chat_members 
		SET last_read_at = NOW(), unread_count = 0
		WHERE chat_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, chatID, userID)
	return err
}

func (r *ChatRepository) RemoveMember(ctx context.Context, chatID, userID uuid.UUID) error {
	query := `DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, chatID, userID)
	return err
}

func (r *ChatRepository) GetUserChats(ctx context.Context, userID uuid.UUID, limit, offset int) ([]db.ChatWithLastMessage, int, error) {
	countQuery := `
		SELECT COUNT(DISTINCT c.id)
		FROM chats c
		JOIN chat_members cm ON c.id = cm.chat_id
		WHERE cm.user_id = $1`

	var total int
	err := r.db.QueryRow(ctx, countQuery, userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT c.id, c.type,
			   COALESCE(c.name, other_user.full_name) as chat_name,
			   COALESCE(c.avatar, other_user.avatar) as chat_avatar,
			   c.created_by, c.created_at, c.updated_at,
			   lm.content as last_message, lm.created_at as last_message_at,
			   COALESCE(cm.unread_count, 0) as unread_count
		FROM chats c
		JOIN chat_members cm ON c.id = cm.chat_id
		LEFT JOIN LATERAL (
			SELECT COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name) as full_name,
				   u.avatar
			FROM chat_members cm2
			JOIN users u ON cm2.user_id = u.id
			LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
			LEFT JOIN student_profiles sp ON sp.user_id = u.id
			LEFT JOIN university_profiles up ON up.user_id = u.id
			LEFT JOIN customer_profiles cp ON cp.user_id = u.id
			WHERE cm2.chat_id = c.id AND cm2.user_id != $1
			LIMIT 1
		) other_user ON c.type = 'direct'
		LEFT JOIN (
			SELECT cm.chat_id, cmc.content, cm.created_at
			FROM chat_messages cm
			JOIN chat_message_contents cmc ON cm.id = cmc.message_id
			JOIN (
				SELECT chat_id, MAX(created_at) as max_created
				FROM chat_messages
				GROUP BY chat_id
			) latest ON cm.chat_id = latest.chat_id AND cm.created_at = latest.max_created
		) lm ON c.id = lm.chat_id
		WHERE cm.user_id = $1
		ORDER BY COALESCE(lm.created_at, c.created_at) DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var chats []db.ChatWithLastMessage
	for rows.Next() {
		var c db.ChatWithLastMessage
		var chatNamePtr, chatAvatarPtr, createdByPtr, lastMsg sql.NullString
		var lastMsgAt sql.NullTime

		err := rows.Scan(
			&c.ID, &c.Type, &chatNamePtr, &chatAvatarPtr, &createdByPtr, &c.CreatedAt, &c.UpdatedAt,
			&lastMsg, &lastMsgAt, &c.UnreadCount,
		)
		if err != nil {
			return nil, 0, err
		}

		if chatNamePtr.Valid {
			c.Name = &chatNamePtr.String
		}
		if chatAvatarPtr.Valid {
			c.Avatar = &chatAvatarPtr.String
		}
		if createdByPtr.Valid {
			c.CreatedBy = &createdByPtr.String
		}
		if lastMsg.Valid {
			c.LastMessage = &lastMsg.String
		}
		if lastMsgAt.Valid {
			c.LastMessageAt = &lastMsgAt.Time
		}

		chats = append(chats, c)
	}

	return chats, total, nil
}

func (r *ChatRepository) GetChatByID(ctx context.Context, chatID uuid.UUID) (*db.Chat, error) {
	query := `
		SELECT id, type, name, avatar, created_by, created_at, updated_at
		FROM chats WHERE id = $1`

	var chat db.Chat
	var namePtr, avatarPtr, createdByPtr sql.NullString

	err := r.db.QueryRow(ctx, query, chatID).Scan(
		&chat.ID, &chat.Type, &namePtr, &avatarPtr, &createdByPtr, &chat.CreatedAt, &chat.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if namePtr.Valid {
		chat.Name = &namePtr.String
	}
	if avatarPtr.Valid {
		chat.Avatar = &avatarPtr.String
	}
	if createdByPtr.Valid {
		chat.CreatedBy = &createdByPtr.String
	}

	return &chat, nil
}

func (r *ChatRepository) GetChatMembers(ctx context.Context, chatID uuid.UUID) ([]db.ChatMemberWithProfile, error) {
	query := `
		SELECT cm.id, cm.chat_id, cm.user_id, cm.role, cm.joined_at,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name) as user_name,
			   u.login, u.avatar
		FROM chat_members cm
		JOIN users u ON cm.user_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE cm.chat_id = $1
		ORDER BY cm.joined_at`

	rows, err := r.db.Query(ctx, query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []db.ChatMemberWithProfile
	for rows.Next() {
		var m db.ChatMemberWithProfile
		var userName, userLogin sql.NullString
		var userAvatar sql.NullString

		err := rows.Scan(
			&m.ID, &m.ChatID, &m.UserID, &m.Role, &m.JoinedAt,
			&userName, &userLogin, &userAvatar,
		)
		if err != nil {
			return nil, err
		}

		if userName.Valid {
			m.UserName = userName.String
		}
		if userLogin.Valid {
			m.UserLogin = userLogin.String
		}
		if userAvatar.Valid {
			m.UserAvatar = &userAvatar.String
		}

		members = append(members, m)
	}

	return members, nil
}

func (r *ChatRepository) IsMember(ctx context.Context, chatID, userID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM chat_members WHERE chat_id = $1 AND user_id = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, query, chatID, userID).Scan(&exists)
	return exists, err
}

func (r *ChatRepository) FindDirectChat(ctx context.Context, userID1, userID2 uuid.UUID) (*db.Chat, error) {
	query := `
		SELECT c.id, c.type, c.name, c.avatar, c.created_by, c.created_at, c.updated_at
		FROM chats c
		JOIN chat_members cm1 ON c.id = cm1.chat_id AND cm1.user_id = $1
		JOIN chat_members cm2 ON c.id = cm2.chat_id AND cm2.user_id = $2
		WHERE c.type = 'direct'
		LIMIT 1`
	var chat db.Chat
	err := r.db.QueryRow(ctx, query, userID1, userID2).Scan(
		&chat.ID, &chat.Type, &chat.Name, &chat.Avatar, &chat.CreatedBy, &chat.CreatedAt, &chat.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) CreateMessage(ctx context.Context, chatID, userID uuid.UUID) (*db.ChatMessage, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO chat_messages (chat_id, user_id)
		VALUES ($1, $2)
		RETURNING id, chat_id, user_id, created_at`

	var msg db.ChatMessage
	err = tx.QueryRow(ctx, query, chatID, userID).Scan(
		&msg.ID, &msg.ChatID, &msg.UserID, &msg.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `
		UPDATE chat_members 
		SET unread_count = unread_count + 1 
		WHERE chat_id = $1 AND user_id != $2`, chatID, userID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (r *ChatRepository) CreateMessageContent(ctx context.Context, messageID uuid.UUID, msgType, content string, metadata map[string]interface{}) (*db.ChatMessageContent, error) {
	metadataJSON, _ := json.Marshal(metadata)

	query := `
		INSERT INTO chat_message_contents (message_id, type, content, metadata)
		VALUES ($1, $2, $3, $4)
		RETURNING id, message_id, type, content, metadata`

	var mc db.ChatMessageContent
	var metadataBytes []byte

	err := r.db.QueryRow(ctx, query, messageID, msgType, content, metadataJSON).Scan(
		&mc.ID, &mc.MessageID, &mc.Type, &mc.Content, &metadataBytes,
	)
	if err != nil {
		return nil, err
	}

	if len(metadataBytes) > 0 {
		json.Unmarshal(metadataBytes, &mc.Metadata)
	}

	return &mc, nil
}

func (r *ChatRepository) GetMessages(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]db.ChatMessageWithContent, int, error) {
	countQuery := `SELECT COUNT(*) FROM chat_messages WHERE chat_id = $1`
	var total int
	err := r.db.QueryRow(ctx, countQuery, chatID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT cm.id, cm.chat_id, cm.user_id, cm.created_at,
			   cmc.id, cmc.message_id, cmc.type, cmc.content, cmc.metadata,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name) as user_name,
			   u.avatar
		FROM chat_messages cm
		JOIN chat_message_contents cmc ON cm.id = cmc.message_id
		JOIN users u ON cm.user_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE cm.chat_id = $1
		ORDER BY cm.created_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(ctx, query, chatID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var messages []db.ChatMessageWithContent
	for rows.Next() {
		var m db.ChatMessageWithContent
		var mc db.ChatMessageContent
		var userName sql.NullString
		var userAvatar sql.NullString
		var metadataBytes []byte

		err := rows.Scan(
			&m.ID, &m.ChatID, &m.UserID, &m.CreatedAt,
			&mc.ID, &mc.MessageID, &mc.Type, &mc.Content, &metadataBytes,
			&userName, &userAvatar,
		)
		if err != nil {
			return nil, 0, err
		}

		m.Content = mc
		if userName.Valid {
			m.UserName = userName.String
		}
		if userAvatar.Valid {
			m.UserAvatar = &userAvatar.String
		}
		if len(metadataBytes) > 0 {
			json.Unmarshal(metadataBytes, &m.Content.Metadata)
		}

		messages = append(messages, m)
	}

	return messages, total, nil
}

func (r *ChatRepository) GetMessageByID(ctx context.Context, messageID uuid.UUID) (*db.ChatMessageWithContent, error) {
	query := `
		SELECT cm.id, cm.chat_id, cm.user_id, cm.created_at,
			   cmc.id, cmc.message_id, cmc.type, cmc.content, cmc.metadata,
			   COALESCE(bp.full_name, sp.full_name, up.full_name, cp.full_name) as user_name,
			   u.avatar
		FROM chat_messages cm
		JOIN chat_message_contents cmc ON cm.id = cmc.message_id
		JOIN users u ON cm.user_id = u.id
		LEFT JOIN brsm_profiles bp ON bp.user_id = u.id
		LEFT JOIN student_profiles sp ON sp.user_id = u.id
		LEFT JOIN university_profiles up ON up.user_id = u.id
		LEFT JOIN customer_profiles cp ON cp.user_id = u.id
		WHERE cm.id = $1`

	var m db.ChatMessageWithContent
	var mc db.ChatMessageContent
	var userName sql.NullString
	var userAvatar sql.NullString
	var metadataBytes []byte

	err := r.db.QueryRow(ctx, query, messageID).Scan(
		&m.ID, &m.ChatID, &m.UserID, &m.CreatedAt,
		&mc.ID, &mc.MessageID, &mc.Type, &mc.Content, &metadataBytes,
		&userName, &userAvatar,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	m.Content = mc
	if userName.Valid {
		m.UserName = userName.String
	}
	if userAvatar.Valid {
		m.UserAvatar = &userAvatar.String
	}
	if len(metadataBytes) > 0 {
		json.Unmarshal(metadataBytes, &m.Content.Metadata)
	}

	return &m, nil
}

func (r *ChatRepository) GetChatMembersExcept(ctx context.Context, chatID uuid.UUID, exceptUserID uuid.UUID) ([]uuid.UUID, error) {
	query := `SELECT user_id FROM chat_members WHERE chat_id = $1 AND user_id != $2`
	rows, err := r.db.Query(ctx, query, chatID, exceptUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []uuid.UUID
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}

func (r *ChatRepository) GetUserChatsForRealtime(ctx context.Context, userID uuid.UUID) ([]string, error) {
	query := `SELECT chat_id FROM chat_members WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatIDs []string
	for rows.Next() {
		var chatID string
		if err := rows.Scan(&chatID); err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}

	return chatIDs, nil
}
