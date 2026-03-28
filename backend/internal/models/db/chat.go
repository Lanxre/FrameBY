package db

import "time"

type Chat struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Name      *string   `json:"name"`
	Avatar    *string   `json:"avatar"`
	CreatedBy *string   `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatMember struct {
	ID       string    `json:"id"`
	ChatID   string    `json:"chat_id"`
	UserID   string    `json:"user_id"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type ChatMessage struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatMessageContent struct {
	ID        string                 `json:"id"`
	MessageID string                 `json:"message_id"`
	Type      string                 `json:"type"`
	Content   string                 `json:"content"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type ChatWithLastMessage struct {
	Chat
	LastMessage   *string    `json:"last_message"`
	LastMessageAt *time.Time `json:"last_message_at"`
	UnreadCount   int        `json:"unread_count"`
}

type ChatMessageWithContent struct {
	ChatMessage
	Content    ChatMessageContent `json:"content"`
	UserName   string             `json:"user_name"`
	UserAvatar *string            `json:"user_avatar"`
}

type ChatDetails struct {
	Chat
	Members []ChatMemberWithProfile `json:"members"`
}

type ChatMemberWithProfile struct {
	ChatMember
	UserName   string  `json:"user_name"`
	UserLogin  string  `json:"user_login"`
	UserAvatar *string `json:"user_avatar"`
}
