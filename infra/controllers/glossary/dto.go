package glossary

import (
	"ponysd-characters/services/glossary"

	"github.com/google/uuid"
)

type GlossaryRecordResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Type  string    `json:"type"`
	Value string    `json:"value"`
	Name  string    `json:"name,omitempty"`
}

func mapGlossaryRecordToResponse(c *glossary.GlossaryRecord) *GlossaryRecordResponse {
	if c == nil {
		return nil
	}

	return &GlossaryRecordResponse{
		ID:    c.ID,
		Type:  c.Type,
		Value: c.Value,
	}
}

func mapGlossaryRecordLocalToResponse(c *glossary.GlossaryRecordLocal) *GlossaryRecordResponse {
	if c == nil {
		return nil
	}

	return &GlossaryRecordResponse{
		Type:  c.Type,
		Value: c.Value,
	}
}
