package services

import (
	"context"
	"fmt"
	"ponysd-characters/appconfig"
	"ponysd-characters/services/character"
	"ponysd-characters/services/filestorage"
	"ponysd-characters/services/glossary"
	"ponysd-characters/services/imagedb"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matthewhartstonge/argon2"
	"github.com/minio/minio-go/v7"
)

type ServicesContext interface {
	BeginTx(ctx context.Context) (pgx.Tx, ServicesContext, error)
	CharacterService() *character.Service
	FileStorageService() *filestorage.Service
	ImageService() *imagedb.Service
	GlossaryService() *glossary.Service
}

type servicesContext struct {
	pg          *pgxpool.Pool
	minioClient *minio.Client
	argon       *argon2.Config
	config      *appconfig.AppConfig
}

func NewServicesContext(
	pgPool *pgxpool.Pool,
	minioClient *minio.Client,
	config *appconfig.AppConfig,
) ServicesContext {
	return &servicesContext{
		pg:          pgPool,
		minioClient: minioClient,
		config:      config,
	}
}

func (svc *servicesContext) BeginTx(ctx context.Context) (pgx.Tx, ServicesContext, error) {
	tx, err := svc.pg.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	return tx, &txContext{
		pg:          tx,
		minioClient: svc.minioClient,
		config:      svc.config,
	}, nil
}

func (svc *servicesContext) CharacterService() *character.Service {
	return character.New(svc.pg, svc.minioClient)
}

func (svc *servicesContext) FileStorageService() *filestorage.Service {
	return filestorage.New(svc.minioClient)
}

func (svc *servicesContext) ImageService() *imagedb.Service {
	return imagedb.New(svc.pg, filestorage.New(svc.minioClient))
}

func (svc *servicesContext) GlossaryService() *glossary.Service {
	return glossary.New(svc.pg)
}

type txContext struct {
	pg          pgx.Tx
	minioClient *minio.Client
	config      *appconfig.AppConfig
}

func (svc *txContext) BeginTx(ctx context.Context) (pgx.Tx, ServicesContext, error) {
	return nil, nil, fmt.Errorf("cannot create nested transaction")
}

func (svc *txContext) CharacterService() *character.Service {
	return character.New(svc.pg, svc.minioClient)
}

func (svc *txContext) FileStorageService() *filestorage.Service {
	return filestorage.New(svc.minioClient)
}

func (svc *txContext) ImageService() *imagedb.Service {
	return imagedb.New(svc.pg, filestorage.New(svc.minioClient))
}

func (svc *txContext) GlossaryService() *glossary.Service {
	return glossary.New(svc.pg)
}
