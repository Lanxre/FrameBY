package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/models/dto"
	"github.com/lanxre/frameby/internal/repositories"
)

type ChatService struct {
	repo *repositories.ChatRepository
}

func NewChatService(repo *repositories.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreateChat(ctx context.Context, userID uuid.UUID, req *dto.CreateChatRequest) (*db.Chat, error) {
	var name *string
	if req.Name != "" {
		name = &req.Name
	}

	chat, err := s.repo.CreateChat(ctx, req.Type, name, userID)
	if err != nil {
		return nil, err
	}

	chatUUID, _ := uuid.Parse(chat.ID)
	if err := s.repo.AddMember(ctx, chatUUID, userID, "admin"); err != nil {
		return nil, err
	}

	for _, userIDStr := range req.Users {
		newUserID, err := uuid.Parse(userIDStr)
		if err != nil {
			continue
		}
		if err := s.repo.AddMember(ctx, chatUUID, newUserID, "member"); err != nil {
			continue
		}
	}

	return chat, nil
}

func (s *ChatService) AddMember(ctx context.Context, chatID, userID uuid.UUID) error {
	return s.repo.AddMember(ctx, chatID, userID, "member")
}

func (s *ChatService) RemoveMember(ctx context.Context, chatID, userID uuid.UUID) error {
	return s.repo.RemoveMember(ctx, chatID, userID)
}

func (s *ChatService) MarkAsRead(ctx context.Context, chatID, userID uuid.UUID) error {
	return s.repo.MarkChatAsRead(ctx, chatID, userID)
}

func (s *ChatService) GetUserChats(ctx context.Context, userID uuid.UUID, limit, offset int) (*dto.ChatListResponse, error) {
	chats, total, err := s.repo.GetUserChats(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ChatResponse, len(chats))
	for i, chat := range chats {
		responses[i] = dto.ChatResponse{
			ID:          chat.ID,
			Type:        chat.Type,
			Name:        chat.Name,
			Avatar:      chat.Avatar,
			UnreadCount: chat.UnreadCount,
		}
		if chat.LastMessage != nil {
			responses[i].LastMessage = chat.LastMessage
		}
		if chat.LastMessageAt != nil {
			formatted := chat.LastMessageAt.Format("2006-01-02T15:04:05Z")
			responses[i].LastMessageAt = &formatted
		}
		responses[i].CreatedAt = chat.CreatedAt.Format("2006-01-02T15:04:05Z")

		if chat.Type == "direct" && chat.Name == nil {
			chatUUID, _ := uuid.Parse(chat.ID)
			members, err := s.repo.GetChatMembers(ctx, chatUUID)
			if err == nil {
				for _, m := range members {
					if m.UserID != userID.String() && m.UserName != "" {
						responses[i].Name = &m.UserName
						break
					}
				}
			}
		}
	}

	return &dto.ChatListResponse{
		Chats: responses,
		Total: total,
	}, nil
}

func (s *ChatService) GetChatByID(ctx context.Context, chatID uuid.UUID) (*db.ChatDetails, error) {
	chat, err := s.repo.GetChatByID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	if chat == nil {
		return nil, errors.New("chat not found")
	}

	members, err := s.repo.GetChatMembers(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return &db.ChatDetails{
		Chat:    *chat,
		Members: members,
	}, nil
}

func (s *ChatService) GetChatMembers(ctx context.Context, chatID uuid.UUID) ([]db.ChatMemberWithProfile, error) {
	return s.repo.GetChatMembers(ctx, chatID)
}

func (s *ChatService) IsMember(ctx context.Context, chatID, userID uuid.UUID) (bool, error) {
	return s.repo.IsMember(ctx, chatID, userID)
}

func (s *ChatService) FindDirectChat(ctx context.Context, userID, otherUserID uuid.UUID) (*db.Chat, error) {
	return s.repo.FindDirectChat(ctx, userID, otherUserID)
}

func (s *ChatService) SendMessage(ctx context.Context, chatID, userID uuid.UUID, req *dto.SendMessageRequest) (*db.ChatMessageWithContent, error) {
	isMember, err := s.repo.IsMember(ctx, chatID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("user is not a member of this chat")
	}

	msg, err := s.repo.CreateMessage(ctx, chatID, userID)
	if err != nil {
		return nil, err
	}

	msgUUID, _ := uuid.Parse(msg.ID)

	metadata := req.Metadata
	if metadata == nil {
		metadata = make(map[string]any)
	}

	content, err := s.repo.CreateMessageContent(ctx, msgUUID, req.Type, req.Content, metadata)
	if err != nil {
		return nil, err
	}

	userIDs, err := s.repo.GetChatMembersExcept(ctx, chatID, userID)
	if err != nil {
		return nil, err
	}

	_ = userIDs

	msgWithContent := &db.ChatMessageWithContent{
		ChatMessage: *msg,
		Content:     *content,
	}

	return msgWithContent, nil
}

func (s *ChatService) GetMessages(ctx context.Context, chatID, userID uuid.UUID, limit, offset int) (*dto.ChatMessagesResponse, error) {
	isMember, err := s.repo.IsMember(ctx, chatID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("user is not a member of this chat")
	}

	messages, total, err := s.repo.GetMessages(ctx, chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.MessageResponse, len(messages))
	for i, msg := range messages {
		responses[i] = dto.MessageResponse{
			ID:         msg.ID,
			ChatID:     msg.ChatID,
			UserID:     msg.UserID,
			UserName:   msg.UserName,
			UserAvatar: msg.UserAvatar,
			Type:       msg.Content.Type,
			Content:    msg.Content.Content,
			Metadata:   msg.Content.Metadata,
			CreatedAt:  msg.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	return &dto.ChatMessagesResponse{
		Messages: responses,
		Total:    total,
	}, nil
}

func (s *ChatService) GetUserChatsForRealtime(ctx context.Context, userID uuid.UUID) ([]string, error) {
	return s.repo.GetUserChatsForRealtime(ctx, userID)
}

func (s *ChatService) GetMessageWithUser(ctx context.Context, messageID uuid.UUID) (*db.ChatMessageWithContent, error) {
	return s.repo.GetMessageByID(ctx, messageID)
}

func (s *ChatService) GetChatMembersExcept(ctx context.Context, chatID, exceptUserID uuid.UUID) ([]uuid.UUID, error) {
	return s.repo.GetChatMembersExcept(ctx, chatID, exceptUserID)
}
