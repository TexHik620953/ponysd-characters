package character

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Character struct {
	ID        uuid.UUID
	OwnerID   uuid.UUID
	CreatedAt time.Time

	Name string

	Nationality string
	Age         int

	Body   string
	Breast string
	Butt   string

	EyesColor string

	HairStyle string
	HairColor string

	CharacterMeta map[string]float32

	MainImage uuid.UUID
}

func getTag(tag, value string, weight float32) string {
	return fmt.Sprintf("(%s %s:%.2f)", value, tag, weight)
}

func (character *Character) BuildTags() []string {

	positiveTags := make([]string, 0)
	positiveTags = append(positiveTags, getTag("-years-old", fmt.Sprintf("%s %d", character.Nationality, character.Age), 1.5))
	positiveTags = append(positiveTags, getTag("body", character.Body, 0.3))
	positiveTags = append(positiveTags, getTag("chest", character.Breast, 1.25))
	positiveTags = append(positiveTags, getTag("butt", character.Butt, 1.25))
	positiveTags = append(positiveTags, getTag("eyes", character.EyesColor, 1.25))
	positiveTags = append(positiveTags, getTag("hairstyle", character.HairStyle, 1.25))
	positiveTags = append(positiveTags, getTag("hair", character.HairColor, 1.25))

	for k, v := range character.CharacterMeta {
		if v <= 0.01 {
			continue
		}
		positiveTags = append(positiveTags, fmt.Sprintf("(%s:%.2f)", k, v))
	}

	return positiveTags
}
