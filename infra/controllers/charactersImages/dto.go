package charactersimages

import (
	"ponysd-characters/services/imagedb"
	"time"

	"github.com/google/uuid"
)

type FileInfoResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	CharacterID uuid.UUID `json:"character_id"`
	FileHash    string    `json:"file_hash"`
}

func mapServiceFileInfoToResponse(c *imagedb.FileInfo) *FileInfoResponse {
	if c == nil {
		return nil
	}

	return &FileInfoResponse{
		ID:          c.ID,
		CreatedAt:   c.CreatedAt,
		CharacterID: c.CharacterID,
		FileHash:    c.FileHash,
	}
}

type GenerateImageRequest struct {
	Action      string `json:"action"`
	Clothes     string `json:"clothes"`
	Environment string `json:"environment"`
	Background  string `json:"background"`
	NSFW        string `json:"nsfw"`

	AdditionalTags []string `json:"tags"`

	CustomPositive string `json:"custom_positive"`
}
