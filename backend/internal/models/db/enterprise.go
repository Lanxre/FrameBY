package db

import (
	"time"

	"github.com/google/uuid"
)

type EnterpriseEntity struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   *string   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
