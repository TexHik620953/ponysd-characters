package task

import (
	"context"
	query "ponysd-characters/internal/repos"
	"ponysd-characters/pkg/rmq"

	"github.com/google/uuid"
)

const imagegenQueue = "img_gen_tasks"

type Service struct {
	tx  query.DBTX
	rmq *rmq.Rmq
}

func New(
	tx query.DBTX,
	rmq *rmq.Rmq,
) *Service {
	svc := &Service{
		tx:  tx,
		rmq: rmq,
	}
	return svc
}

func (svc *Service) AddImageGenerationTask(ctx context.Context, task *ImageGenerationTask) (uuid.UUID, error) {
	q := query.New(svc.tx)
	id, err := q.CreateTask(ctx, query.CreateTaskParams{
		Status:      string(StatusPending),
		Positive:    task.PositivePrompt,
		Negative:    task.NegativePrompt,
		Seed:        int32(task.Seed),
		CharacterID: task.CharacterID,
		UserID:      task.UserID,
	})
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (svc *Service) PullPendingTask(ctx context.Context) (*Task, error) {
	q := query.New(svc.tx)
	task, err := q.GetTaskWithStatus(ctx, string(StatusPending))
	if err != nil {
		return nil, err
	}

	_, err = q.UpdateTask(ctx, query.UpdateTaskParams{
		ID:     task.ID,
		Status: string(StatusInWork),
	})
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:             task.ID,
		UserID:         task.UserID,
		CharacterID:    task.CharacterID,
		PositivePrompt: task.Positive,
		NegativePrompt: task.Negative,
		Seed:           int(task.Seed),
	}, nil
}

func (svc *Service) SetTaskStatus(ctx context.Context, id uuid.UUID, status TaskStatus) error {
	q := query.New(svc.tx)

	_, err := q.UpdateTask(ctx, query.UpdateTaskParams{
		ID:     id,
		Status: string(StatusPending),
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetTaskByID(ctx context.Context, id uuid.UUID) (*Task, error) {
	q := query.New(svc.tx)

	task, err := q.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:             task.ID,
		UserID:         task.UserID,
		CharacterID:    task.CharacterID,
		PositivePrompt: task.Positive,
		NegativePrompt: task.Negative,
		Seed:           int(task.Seed),
	}, nil
}
