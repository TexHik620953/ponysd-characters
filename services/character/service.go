package character

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	query "ponysd-characters/internal/repos"
	"ponysd-characters/pkg/rmq"
	"ponysd-characters/services/glossary"
	"ponysd-characters/services/task"
	"strings"

	"github.com/0x6flab/namegenerator"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	tx       query.DBTX
	glossary *glossary.Service
	task     *task.Service
}

func New(
	tx query.DBTX,
	redis *redis.Client,
	minioClient *minio.Client,
) *Service {
	return &Service{
		tx:       tx,
		glossary: glossary.New(tx),
		task:     task.New(rmq.New(redis)),
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
func (svc *Service) GetCharacter(ctx context.Context, userID, characterID uuid.UUID) (*Character, error) {
	q := query.New(svc.tx)
	data, err := q.GetCharacter(ctx, query.GetCharacterParams{
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
func (svc *Service) DeleteCharacter(ctx context.Context, userID, characterID uuid.UUID) error {
	q := query.New(svc.tx)
	return q.DeleteCharacter(ctx, query.DeleteCharacterParams{
		OwnerID: userID,
		ID:      characterID,
	})
}

// Character creating with fields validation

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
	rnationality := glossary.PickRandomGlossary(gnat, rng)
	rbody := glossary.PickRandomGlossary(gbody, rng)
	rbreast := glossary.PickRandomGlossary(gbreast, rng)
	rbutt := glossary.PickRandomGlossary(gbutt, rng)
	reyesColor := glossary.PickRandomGlossary(geyesColor, rng)
	rhairColor := glossary.PickRandomGlossary(ghairColor, rng)
	rhairStyle := glossary.PickRandomGlossary(ghairStyle, rng)

	if len(character.Name) == 0 {
		gen := namegenerator.NewGenerator().WithGender(namegenerator.Female)
		character.Name = gen.Generate()
	}

	if character.Age < 18 || character.Age > 40 {
		character.Age = rage
	}
	glossary.SetIfInvalid(&character.Nationality, gnat, rnationality)
	glossary.SetIfInvalid(&character.Body, gnat, rbody)
	glossary.SetIfInvalid(&character.Breast, gnat, rbreast)
	glossary.SetIfInvalid(&character.Butt, gnat, rbutt)
	glossary.SetIfInvalid(&character.EyesColor, gnat, reyesColor)
	glossary.SetIfInvalid(&character.HairColor, gnat, rhairColor)
	glossary.SetIfInvalid(&character.HairStyle, gnat, rhairStyle)

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

	env, err := svc.glossary.NewEnvironment(ctx, rng)
	if err != nil {
		return err
	}
	img, err := svc.CreateImage(ctx, character.OwnerID, character.ID, env, seed)
	_ = img
	return err
}

// Image generation

func (svc *Service) CreateImage(ctx context.Context, userID, characterID uuid.UUID, environment *glossary.CharacterEnvironment, seed int64) (*task.TaskInfo, error) {
	rng := rand.New(rand.NewSource(seed))

	// Get character
	character, err := svc.GetCharacter(ctx, userID, characterID)
	if err != nil {
		return nil, err
	}

	err = svc.glossary.FixEnvironment(ctx, environment, rng)
	if err != nil {
		return nil, err
	}

	defaultPositiveTags, err := svc.glossary.ListRecordsByType(ctx, glossary.KEY_DEFAULT_POSITIVE)
	if err != nil {
		return nil, err
	}
	defaultNegativeTags, err := svc.glossary.ListRecordsByType(ctx, glossary.KEY_DEFAULT_NEGATIVE)
	if err != nil {
		return nil, err
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

	task, err := svc.task.AddImageGenerationTask(ctx, &task.ImageGenerationTask{
		PositivePrompt: strings.Join(positiveTags, ", "),
		NetagivePrompt: strings.Join(negativeTags, ", "),
		Seed:           int(seed),
	})

	if err != nil {
		return nil, err
	}

	return task, nil
}
