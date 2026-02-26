package task

import "github.com/google/uuid"

type ImageGenerationTask struct {
	PositivePrompt string `json:"positive_prompt"`
	NetagivePrompt string `json:"negative_prompt"`
	Seed           int    `json:"seed"`
}

type TaskInfo struct {
	Position int
	ID       uuid.UUID
}
