package task

import "github.com/google/uuid"

type ImageGenerationTask struct {
	UserID         uuid.UUID `json:"user_id"`
	CharacterID    uuid.UUID `json:"character_id"`
	PositivePrompt string    `json:"positive_prompt"`
	NegativePrompt string    `json:"negative_prompt"`
	Seed           int       `json:"seed"`
}
type Task struct {
	ID             uuid.UUID
	UserID         uuid.UUID `json:"user_id"`
	CharacterID    uuid.UUID `json:"character_id"`
	PositivePrompt string    `json:"positive_prompt"`
	NegativePrompt string    `json:"negative_prompt"`
	Seed           int       `json:"seed"`
}

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusInWork  TaskStatus = "inwork"
	StatusDone    TaskStatus = "done"
	StatusFailed  TaskStatus = "failed"
)
