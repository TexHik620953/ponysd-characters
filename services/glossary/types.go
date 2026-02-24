package glossary

import (
	"github.com/google/uuid"
)

const (
	KEY_DEFAULT_POSITIVE = "default_positive"
	KEY_DEFAULT_NEGATIVE = "default_negative"

	KEY_NATIONALITY = "nationality"
	KEY_BODY_TYPE   = "body_type"
	KEY_BREAST_TYPE = "breast_type"
	KEY_BUTT_TYPE   = "butt_type"
	KEY_EYES_COLOR  = "eyes_color"
	KEY_HAIR_COLOR  = "hair_color"
	KEY_HAIR_STYLE  = "hair_style"

	KEY_ACTION      = "action"
	KEY_CLOTHES     = "clothes"
	KEY_ENVIRONMENT = "environment"
	KEY_BACKGROUND  = "background"
	KEY_NSFW        = "nsfw"
)

type GlossaryRecord struct {
	ID    uuid.UUID
	Type  string
	Value string
}

type GlossaryRecordLocal struct {
	Type  string
	Value string
	Name  string
}

type CharacterEnvironment struct {
	Action      string
	Clothes     string
	Environment string
	Background  string
	NSFW        string
}

func (cenv *CharacterEnvironment) BuildTags() []string {
	positiveTags := make([]string, 0)

	positiveTags = append(positiveTags, string(cenv.Action))
	positiveTags = append(positiveTags, string(cenv.Clothes))
	positiveTags = append(positiveTags, string(cenv.Environment))
	positiveTags = append(positiveTags, string(cenv.Background))
	positiveTags = append(positiveTags, string(cenv.NSFW))

	return positiveTags
}
