package task

import (
	"context"
	"ponysd-characters/pkg/rmq"
)

const imagegenQueue = "img_gen_tasks"

type Service struct {
	rmq *rmq.Rmq
}

func New(
	rmq *rmq.Rmq,
) *Service {
	svc := &Service{
		rmq: rmq,
	}
	return svc
}

func (svc *Service) AddImageGenerationTask(ctx context.Context, task *ImageGenerationTask) (*TaskInfo, error) {
	r, err := svc.rmq.Enqueue(ctx, task, imagegenQueue)
	if err != nil {
		return nil, err
	}
	return &TaskInfo{
		ID:       r.MessageID,
		Position: r.Position,
	}, nil
}
