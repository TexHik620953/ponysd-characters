package imagedb

import (
	"time"

	"github.com/google/uuid"
)

type FileInfo struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	CharacterID uuid.UUID
	FileHash    string
	IsMain      bool
}
