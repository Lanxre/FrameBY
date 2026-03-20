package static

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	UPLOAD_DIR        string = "./uploads"
	UPLOAD_AVATAR_DIR string = "./uploads/avatars"
)

func CreateStaticDirs(r *gin.Engine) error {
	if err := createDirs(); err != nil {
		return err
	}

	r.Static("/uploads", UPLOAD_DIR)

	return nil
}

func createDirs() error {
	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		return fmt.Errorf("failed to create uploads directory: %w", err)
	}

	if err := os.MkdirAll(UPLOAD_AVATAR_DIR, 0755); err != nil {
		return fmt.Errorf("failed to create avatars directory: %w", err)
	}

	return nil
}