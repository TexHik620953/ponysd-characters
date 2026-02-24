package glossary

import (
	"context"
	"math/rand"
	query "ponysd-characters/internal/repos"
)

type Service struct {
	tx query.DBTX
}

func New(
	tx query.DBTX,
) *Service {
	svc := &Service{
		tx: tx,
	}
	return svc
}

func (svc *Service) ListRecordsByType(ctx context.Context, _type string) ([]GlossaryRecord, error) {
	q := query.New(svc.tx)
	data, err := q.ListGlossaryByType(ctx, _type)
	if err != nil {
		return nil, err
	}

	result := make([]GlossaryRecord, len(data))
	for i, val := range data {
		result[i] = GlossaryRecord{
			ID:    val.ID,
			Type:  val.Type,
			Value: val.Value,
		}
	}

	return result, nil
}

func (svc *Service) ListRecordLocal(ctx context.Context, _type, local string) ([]GlossaryRecordLocal, error) {
	q := query.New(svc.tx)
	data, err := q.ListGlossaryByTypeLocal(ctx, query.ListGlossaryByTypeLocalParams{
		Type:  _type,
		Local: local,
	})
	if err != nil {
		return nil, err
	}

	result := make([]GlossaryRecordLocal, len(data))
	for i, val := range data {
		result[i] = GlossaryRecordLocal{
			Type:  val.Type,
			Value: val.Value,
			Name:  val.Name,
		}
	}

	return result, nil
}

func (svc *Service) PickRandom(ctx context.Context, _type string, rng *rand.Rand) (*GlossaryRecord, error) {
	recs, err := svc.ListRecordsByType(ctx, _type)
	if err != nil {
		return nil, err
	}
	rec := recs[rng.Int31n(int32(len(recs)))]
	return &GlossaryRecord{
		ID:    rec.ID,
		Type:  rec.Type,
		Value: rec.Value,
	}, nil
}

func (svc *Service) NewEnvironment(ctx context.Context, rng *rand.Rand) (*CharacterEnvironment, error) {
	action, err := svc.PickRandom(ctx, KEY_ACTION, rng)
	if err != nil {
		return nil, err
	}
	clothes, err := svc.PickRandom(ctx, KEY_CLOTHES, rng)
	if err != nil {
		return nil, err
	}
	environment, err := svc.PickRandom(ctx, KEY_ENVIRONMENT, rng)
	if err != nil {
		return nil, err
	}
	background, err := svc.PickRandom(ctx, KEY_BACKGROUND, rng)
	if err != nil {
		return nil, err
	}
	nsfw, err := svc.PickRandom(ctx, KEY_NSFW, rng)
	if err != nil {
		return nil, err
	}

	return &CharacterEnvironment{
		Action:      action.Type,
		Clothes:     clothes.Type,
		Environment: environment.Type,
		Background:  background.Type,
		NSFW:        nsfw.Type,
	}, nil
}
