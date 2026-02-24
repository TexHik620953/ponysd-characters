package character

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	query "ponysd-characters/internal/repos"
	"ponysd-characters/services/filestorage"
	"ponysd-characters/services/glossary"
	"ponysd-characters/services/imagedb"

	"github.com/0x6flab/namegenerator"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type Service struct {
	tx       query.DBTX
	glossary *glossary.Service
	imgSvc   *imagedb.Service
}

func New(
	tx query.DBTX,
	minioClient *minio.Client,
) *Service {
	return &Service{
		tx:       tx,
		glossary: glossary.New(tx),
		imgSvc:   imagedb.New(tx, filestorage.New(minioClient)),
	}
}

func (svc *Service) ListPublicCharacters(ctx context.Context, limit, offset int) ([]Character, error) {
	q := query.New(svc.tx)
	data, err := q.ListPublicCharacters(ctx, query.ListPublicCharactersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	result := make([]Character, len(data))
	for i, val := range data {
		var meta map[string]float32
		if err := json.Unmarshal([]byte(val.MetaParams), &meta); err != nil {
			return nil, fmt.Errorf("failed to unmarshal character meta:%w", err)
		}
		result[i] = Character{
			ID:            val.ID,
			OwnerID:       val.OwnerID,
			CreatedAt:     val.CreatedAt.Time,
			Nationality:   val.Nationality,
			Name:          val.Name,
			Biography:     val.Biography,
			Age:           int(val.Age),
			Body:          val.Body,
			Breast:        val.Breast,
			Butt:          val.Butt,
			EyesColor:     val.EyesColor,
			HairStyle:     val.HairStyle,
			HairColor:     val.HairColor,
			CharacterMeta: meta,
			MainImage:     val.MainImageID,
		}
	}

	return result, nil
}
func (svc *Service) ListUserCharacters(ctx context.Context, userID uuid.UUID) ([]Character, error) {
	q := query.New(svc.tx)
	data, err := q.ListUserCharacters(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]Character, len(data))
	for i, val := range data {
		var meta map[string]float32
		if err := json.Unmarshal([]byte(val.MetaParams), &meta); err != nil {
			return nil, fmt.Errorf("failed to unmarshal character meta:%w", err)
		}
		result[i] = Character{
			ID:            val.ID,
			OwnerID:       val.OwnerID,
			CreatedAt:     val.CreatedAt.Time,
			Nationality:   val.Nationality,
			Name:          val.Name,
			Biography:     val.Biography,
			Age:           int(val.Age),
			Body:          val.Body,
			Breast:        val.Breast,
			Butt:          val.Butt,
			EyesColor:     val.EyesColor,
			HairStyle:     val.HairStyle,
			HairColor:     val.HairColor,
			CharacterMeta: meta,
			MainImage:     val.MainImageID,
		}
	}

	return result, nil
}
func (svc *Service) GetUserCharacter(ctx context.Context, userID, characterID uuid.UUID) (*Character, error) {
	q := query.New(svc.tx)
	data, err := q.GetUserCharacter(ctx, query.GetUserCharacterParams{
		OwnerID: userID,
		ID:      characterID,
	})
	if err != nil {
		return nil, err
	}

	var meta map[string]float32
	if err := json.Unmarshal([]byte(data.MetaParams), &meta); err != nil {
		return nil, fmt.Errorf("failed to unmarshal character meta:%w", err)
	}

	return &Character{
		ID:            data.ID,
		OwnerID:       data.OwnerID,
		CreatedAt:     data.CreatedAt.Time,
		Nationality:   data.Nationality,
		Biography:     data.Biography,
		Name:          data.Name,
		Age:           int(data.Age),
		Body:          data.Body,
		Breast:        data.Breast,
		Butt:          data.Butt,
		EyesColor:     data.EyesColor,
		HairStyle:     data.HairStyle,
		HairColor:     data.HairColor,
		CharacterMeta: meta,
		MainImage:     data.MainImageID,
	}, nil
}
func (svc *Service) GetPublicCharacter(ctx context.Context, characterID uuid.UUID) (*Character, error) {
	q := query.New(svc.tx)
	data, err := q.GetPublicCharacter(ctx, characterID)
	if err != nil {
		return nil, err
	}

	var meta map[string]float32
	if err := json.Unmarshal([]byte(data.MetaParams), &meta); err != nil {
		return nil, fmt.Errorf("failed to unmarshal character meta:%w", err)
	}

	return &Character{
		ID:            data.ID,
		OwnerID:       data.OwnerID,
		CreatedAt:     data.CreatedAt.Time,
		Nationality:   data.Nationality,
		Biography:     data.Biography,
		Name:          data.Name,
		Age:           int(data.Age),
		Body:          data.Body,
		Breast:        data.Breast,
		Butt:          data.Butt,
		EyesColor:     data.EyesColor,
		HairStyle:     data.HairStyle,
		HairColor:     data.HairColor,
		CharacterMeta: meta,
		MainImage:     data.MainImageID,
	}, nil
}
func (svc *Service) DeleteCharacter(ctx context.Context, userID, characterID uuid.UUID) error {
	q := query.New(svc.tx)
	return q.DeleteCharacter(ctx, query.DeleteCharacterParams{
		OwnerID: userID,
		ID:      characterID,
	})
}

// Character creating with fields validation

func pickRandomGlossary(data []glossary.GlossaryRecord, rng *rand.Rand) string {
	if len(data) == 0 {
		return ""
	}
	return data[rng.Int31n(int32(len(data)))].Value
}
func inGlossary(data []glossary.GlossaryRecord, val string) bool {
	for _, v := range data {
		if v.Value == val {
			return true
		}
	}
	return false
}
func (svc *Service) fixCharacter(ctx context.Context, character *Character, rng *rand.Rand) error {
	// Get variables from glossary and validate character.
	// All missing fields are set to random from glossary

	glos := glossary.New(svc.tx)

	gnat, err := glos.ListRecordsByType(ctx, glossary.KEY_NATIONALITY)
	if err != nil {
		return err
	}
	gbody, err := glos.ListRecordsByType(ctx, glossary.KEY_BODY_TYPE)
	if err != nil {
		return err
	}
	gbreast, err := glos.ListRecordsByType(ctx, glossary.KEY_BREAST_TYPE)
	if err != nil {
		return err
	}
	gbutt, err := glos.ListRecordsByType(ctx, glossary.KEY_BUTT_TYPE)
	if err != nil {
		return err
	}
	geyesColor, err := glos.ListRecordsByType(ctx, glossary.KEY_EYES_COLOR)
	if err != nil {
		return err
	}
	ghairColor, err := glos.ListRecordsByType(ctx, glossary.KEY_HAIR_COLOR)
	if err != nil {
		return err
	}
	ghairStyle, err := glos.ListRecordsByType(ctx, glossary.KEY_HAIR_STYLE)
	if err != nil {
		return err
	}

	rage := int(18 + rng.Int31n(18))
	rnationality := pickRandomGlossary(gnat, rng)
	rbody := pickRandomGlossary(gbody, rng)
	rbreast := pickRandomGlossary(gbreast, rng)
	rbutt := pickRandomGlossary(gbutt, rng)
	reyesColor := pickRandomGlossary(geyesColor, rng)
	rhairColor := pickRandomGlossary(ghairColor, rng)
	rhairStyle := pickRandomGlossary(ghairStyle, rng)

	if len(character.Name) == 0 {
		gen := namegenerator.NewGenerator().WithGender(namegenerator.Female)
		character.Name = gen.Generate()
	}

	if character.Age < 18 || character.Age > 40 {
		character.Age = rage
	}
	if len(character.Nationality) == 0 || !inGlossary(gnat, character.Nationality) {
		character.Nationality = rnationality
	}
	if len(character.Body) == 0 || !inGlossary(gbody, character.Body) {
		character.Body = rbody
	}
	if len(character.Breast) == 0 || !inGlossary(gbreast, character.Breast) {
		character.Breast = rbreast
	}
	if len(character.Butt) == 0 || !inGlossary(gbutt, character.Butt) {
		character.Butt = rbutt
	}
	if len(character.EyesColor) == 0 || !inGlossary(geyesColor, character.EyesColor) {
		character.EyesColor = reyesColor
	}
	if len(character.HairColor) == 0 || !inGlossary(ghairColor, character.HairColor) {
		character.HairColor = rhairColor
	}
	if len(character.HairStyle) == 0 || !inGlossary(ghairStyle, character.HairStyle) {
		character.HairStyle = rhairStyle
	}

	if character.CharacterMeta == nil {
		character.CharacterMeta = map[string]float32{}
	}

	return nil
}
func (svc *Service) CreateCharacter(ctx context.Context, character *Character, seed int64) error {
	rng := rand.New(rand.NewSource(seed))
	// Validation and invalid fields fix
	err := svc.fixCharacter(ctx, character, rng)
	if err != nil {
		return err
	}

	// Generate image
	meta, err := json.Marshal(character.CharacterMeta)
	if err != nil {
		return err
	}
	id, err := query.New(svc.tx).CreateCharacter(ctx, query.CreateCharacterParams{
		OwnerID:     character.OwnerID,
		Nationality: character.Nationality,
		Name:        character.Name,
		Age:         int32(character.Age),
		Body:        character.Body,
		Breast:      character.Breast,
		Butt:        character.Butt,
		EyesColor:   character.EyesColor,
		HairStyle:   character.HairStyle,
		HairColor:   character.HairColor,
		MetaParams:  string(meta),
	})
	if err != nil {
		return err
	}
	character.ID = id

	environment, err := svc.glossary.NewEnvironment(ctx, rng)
	if err != nil {
		return err
	}

	defaultPositiveTags, err := svc.glossary.ListRecordsByType(ctx, glossary.KEY_DEFAULT_POSITIVE)
	if err != nil {
		return err
	}
	defaultNegativeTags, err := svc.glossary.ListRecordsByType(ctx, glossary.KEY_DEFAULT_NEGATIVE)
	if err != nil {
		return err
	}
	negativeTags := make([]string, 0)
	for _, rec := range defaultNegativeTags {
		negativeTags = append(negativeTags, rec.Value)
	}

	positiveTags := make([]string, 0)
	for _, rec := range defaultPositiveTags {
		positiveTags = append(positiveTags, rec.Value)
	}
	positiveTags = append(positiveTags, character.BuildTags()...)
	positiveTags = append(positiveTags, environment.BuildTags()...)

	/*
		// Run task for generation
		task := &task.Task{
			CharacterID:    character.ID,
			Seed:           int(seed),
			PositivePrompt: strings.Join(positiveTags, ", "),
			NegativePrompt: strings.Join(negativeTags, ", "),
			Status:         string(task.TaskStatusPending),
		}
		_, err = svc.taskSvc.CreateImageTask(ctx, task)
	*/
	return err
}
