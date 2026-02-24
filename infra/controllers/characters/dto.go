package characters

import (
	"ponysd-characters/services/character"
	"time"

	"github.com/google/uuid"
)

type CharacterResponse struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   string     `json:"created_at"`
	Name        string     `json:"name"`
	Biography   *string    `json:"bio"`
	Nationality string     `json:"nationality"`
	Age         int        `json:"age"`
	Body        string     `json:"body"`
	Breast      string     `json:"breast"`
	Butt        string     `json:"butt"`
	EyesColor   string     `json:"eyes_color"`
	HairStyle   string     `json:"hair_style"`
	HairColor   string     `json:"hair_color"`
	MainImage   *uuid.UUID `json:"main_image"`
}

func mapCharacterToResponse(c *character.Character) *CharacterResponse {
	if c == nil {
		return nil
	}

	return &CharacterResponse{
		ID:          c.ID,
		CreatedAt:   c.CreatedAt.Format(time.RFC3339),
		Name:        c.Name,
		Biography:   c.Biography,
		Nationality: c.Nationality,
		Age:         c.Age,
		Body:        c.Body,
		Breast:      c.Breast,
		Butt:        c.Butt,
		EyesColor:   c.EyesColor,
		HairStyle:   c.HairStyle,
		HairColor:   c.HairColor,
		MainImage:   c.MainImage,
	}
}

type CreateCharacterRequest struct {
	Name        string  `json:"name"`
	Biography   *string `json:"bio"`
	Nationality string  `json:"nationality"`
	Age         int     `json:"age"`
	Body        string  `json:"body"`
	Breast      string  `json:"breast"`
	Butt        string  `json:"butt"`
	EyesColor   string  `json:"eyes_color"`
	HairStyle   string  `json:"hair_style"`
	HairColor   string  `json:"hair_color"`

	Seed int `json:"seed"`
}
