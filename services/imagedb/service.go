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

func (svc *Service) InsertImage(ctx context.Context, userID, characterID uuid.UUID, fileHash string) (uuid.UUID, error) {
	q := query.New(svc.tx)
	return q.InsertImage(ctx, query.InsertImageParams{
		OwnerID:     userID,
		CharacterID: characterID,
		FileHash:    fileHash,
	})
}
func (svc *Service) GetImage(ctx context.Context, userID, imageID uuid.UUID) (*FileInfo, error) {
	q := query.New(svc.tx)
	img, err := q.GetImage(ctx, query.GetImageParams{
		OwnerID: userID,
		ID:      imageID,
	})
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
func (svc *Service) GetImageData(ctx context.Context, userID, imageID uuid.UUID) ([]byte, error) {
	q := query.New(svc.tx)
	img, err := q.GetImage(ctx, query.GetImageParams{
		OwnerID: userID,
		ID:      imageID,
	})
	if err != nil {
		return nil, err
	}

	return svc.fs.GetFile(ctx, img.FileHash)
}
func (svc *Service) ListCharacterImages(ctx context.Context, userID, characterID uuid.UUID) ([]FileInfo, error) {
	q := query.New(svc.tx)
	images, err := q.ListCharacterImages(ctx, query.ListCharacterImagesParams{
		OwnerID:     userID,
		CharacterID: characterID,
	})
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
