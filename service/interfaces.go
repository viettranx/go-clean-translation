package service

import (
	"context"
	"go-clean-translation/service/entity"
)

type TranslateUseCase interface {
	Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
	FetchHistories(ctx context.Context) ([]entity.Translation, error)
}

type TranslateRepository interface {
	GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
	FindHistories(ctx context.Context) ([]entity.Translation, error)
	InsertTranslation(ctx context.Context, translation entity.Translation) error
}

type GoogleService interface {
	Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
}
