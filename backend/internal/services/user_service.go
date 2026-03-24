package services

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/draw"
)
const maxAvatarSize = 512

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID uuid.UUID) (*db.UserEntity, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) AdminAssignRole(ctx context.Context, userID uuid.UUID, role, fullName, subrole string, enterpriseID *uuid.UUID) error {
	validRoles := map[string]bool{"student": true, "brsm": true, "customer": true, "university": true}
	if !validRoles[role] {
		return errors.New("invalid role")
	}

	return s.repo.AssignProfile(ctx, userID, role, fullName, subrole, enterpriseID)
}

func (s *UserService) SaveAvatar(ctx context.Context, file *multipart.FileHeader, userID uuid.UUID) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать изображение")
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var resized image.Image = img

	if width > maxAvatarSize || height > maxAvatarSize {
		resized = resizeImage(img, maxAvatarSize, maxAvatarSize)
	}

	dir := "./uploads/avatars"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	oldAvatar, err := s.repo.GetAvatarFilename(ctx, userID)
	if err == nil && oldAvatar != "" {
		oldPath := filepath.Join(dir, oldAvatar)
		_ = os.Remove(oldPath)
	}

	newFilename := fmt.Sprintf("%s.jpg", uuid.New().String())
	path := filepath.Join(dir, newFilename)

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	err = jpeg.Encode(out, resized, &jpeg.Options{Quality: 85})
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateAvatar(ctx, userID, &newFilename)
	if err != nil {
		return "", err
	}

	return newFilename, nil
}

func resizeImage(img image.Image, maxW, maxH int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, maxW, maxH))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, login string, password string) error {
	var loginPtr *string
	var passwordPtr *string
	
	if login != "" {
		loginPtr = &login
	}

	if password != "" {
		hash, err := hashPassword(password)
		if err != nil {
			return err
		}
		passwordPtr = &hash
	}

	return s.repo.UpdateProfile(ctx, userID, loginPtr, passwordPtr)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}