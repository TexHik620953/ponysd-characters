package tasks

import "github.com/google/uuid"

type Task struct {
	ID             uuid.UUID
	PositivePrompt string `json:"positive_prompt"`
	NegativePrompt string `json:"negative_prompt"`
	Seed           int    `json:"seed"`
}
