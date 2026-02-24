package imagedb

import (
	"context"
	query "ponysd-characters/internal/repos"
	"ponysd-characters/services/filestorage"

	"github.com/google/uuid"
)

type Service struct {
	tx query.DBTX
	fs *filestorage.Service
}

func New(
	tx query.DBTX,
	fs *filestorage.Service,
) *Service {
	svc := &Service{
		tx: tx,
		fs: fs,
	}
	return svc
}

func (svc *Service) CreateImage(ctx context.Context, characterID uuid.UUID, file []byte) (uuid.UUID, error) {
	// Save file
	fileHash, err := svc.fs.StoreFile(ctx, file)
	if err != nil {
		return uuid.Nil, err
	}

	q := query.New(svc.tx)
	return q.CreateImage(ctx, query.CreateImageParams{
		CharacterID: characterID,
		FileHash:    fileHash,
	})
}
func (svc *Service) GetImage(ctx context.Context, ID uuid.UUID) (*FileInfo, error) {
	q := query.New(svc.tx)
	img, err := q.GetImage(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		ID:          img.ID,
		CreatedAt:   img.CreatedAt.Time,
		CharacterID: img.CharacterID,
		FileHash:    img.FileHash,
		IsMain:      img.IsMain,
	}, nil
}
func (svc *Service) GetMainImage(ctx context.Context, characterID uuid.UUID) (*FileInfo, error) {
	q := query.New(svc.tx)
	img, err := q.GetMainImage(ctx, characterID)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		ID:          img.ID,
		CreatedAt:   img.CreatedAt.Time,
		CharacterID: img.CharacterID,
		FileHash:    img.FileHash,
		IsMain:      img.IsMain,
	}, nil
}
func (svc *Service) GetImageData(ctx context.Context, ID uuid.UUID) ([]byte, error) {
	q := query.New(svc.tx)
	img, err := q.GetImage(ctx, ID)
	if err != nil {
		return nil, err
	}

	return svc.fs.GetFile(ctx, img.FileHash)
}
func (svc *Service) ListImages(ctx context.Context, characterID uuid.UUID) ([]FileInfo, error) {
	q := query.New(svc.tx)
	images, err := q.ListImages(ctx, characterID)
	if err != nil {
		return nil, err
	}

	result := make([]FileInfo, 0)
	for _, img := range images {
		result = append(result, FileInfo{
			ID:          img.ID,
			CreatedAt:   img.CreatedAt.Time,
			CharacterID: img.CharacterID,
			FileHash:    img.FileHash,
			IsMain:      img.IsMain,
		})
	}

	return result, nil
}
func (svc *Service) SetMainImage(ctx context.Context, characterID, ID uuid.UUID) error {
	q := query.New(svc.tx)
	return q.SetMainImage(ctx, query.SetMainImageParams{
		ID:          ID,
		CharacterID: characterID,
	})
}
