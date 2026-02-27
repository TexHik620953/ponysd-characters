package rmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Rmq struct {
	redisClient *redis.Client
}

func New(redisClient *redis.Client) *Rmq {
	return &Rmq{
		redisClient: redisClient,
	}
}

type QueueObj struct {
	ID   uuid.UUID `json:"id"`
	Data any       `json:"data"`
}

type EnqueueResult struct {
	Position  int
	MessageID uuid.UUID
}

// Ошибки
var (
	ErrMessageNotFound = errors.New("message not found")
	ErrDequeueTimeout  = errors.New("dequeue timeout")
)

// Enqueue добавляет элемент в очередь
func (r *Rmq) Enqueue(ctx context.Context, data any, queue string) (*EnqueueResult, error) {
	message := &QueueObj{
		ID:   uuid.New(),
		Data: data,
	}

	messageData, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}

	// Lua скрипт для атомарного добавления и сохранения индекса
	script := `
    local queue = KEYS[1]
    local indexKey = KEYS[2]
    local messageID = ARGV[1]
    local data = ARGV[2]
    
    -- Получаем текущую позицию
    local pos = redis.call('LLEN', queue)
    
    -- Добавляем в очередь
    redis.call('RPUSH', queue, data)
    
    -- Сохраняем индекс сообщения
    redis.call('HSET', indexKey, messageID, pos)
    
    return pos
    `

	pos, err := r.redisClient.Eval(ctx, script, []string{queue, queue + ":index"}, message.ID.String(), messageData).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to execute enqueue script: %w", err)
	}

	return &EnqueueResult{
		Position:  int(pos.(int64)),
		MessageID: message.ID,
	}, nil
}

// GetPosition возвращает текущую позицию сообщения в очереди
func (r *Rmq) GetPosition(ctx context.Context, queue string, messageID uuid.UUID) (int, error) {
	indexKey := queue + ":index"

	// Пытаемся получить позицию из индекса
	pos, err := r.redisClient.HGet(ctx, indexKey, messageID.String()).Int()
	if err == redis.Nil {
		return -1, ErrMessageNotFound
	}
	if err != nil {
		return -1, fmt.Errorf("failed to get position: %w", err)
	}

	return pos, nil
}

// DequeueBlocking извлекает элемент с блокировкой (ждет появления элементов)
func (r *Rmq) DequeueBlocking(ctx context.Context, queue string, timeout time.Duration, data any) (uuid.UUID, error) {
	// Используем BLPOP для блокирующего извлечения
	result, err := r.redisClient.BLPop(ctx, timeout, queue).Result()
	if err == redis.Nil {
		return uuid.Nil, ErrDequeueTimeout
	}
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to blocking dequeue: %w", err)
	}

	// result[0] - имя очереди, result[1] - данные
	if len(result) < 2 {
		return uuid.Nil, errors.New("unexpected BLPOP result")
	}

	message := QueueObj{
		Data: data,
	}
	if err := json.Unmarshal([]byte(result[1]), &message); err != nil {
		return uuid.Nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}

	// Удаляем индекс сообщения
	go r.cleanupIndex(context.Background(), queue, message.ID)

	return message.ID, nil
}

// cleanupIndex удаляет индекс сообщения (внутренний метод)
func (r *Rmq) cleanupIndex(ctx context.Context, queue string, messageID uuid.UUID) {
	indexKey := queue + ":index"

	// Используем отдельный контекст с таймаутом
	cleanupCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Удаляем индекс
	if err := r.redisClient.HDel(cleanupCtx, indexKey, messageID.String()).Err(); err != nil {
		// Логируем ошибку, но не возвращаем её
		fmt.Printf("Failed to cleanup index for message %s: %v\n", messageID, err)
	}
}
