package mysql

import (
	"context"
	"go-clean-translation/service"
	"go-clean-translation/service/entity"
	"gorm.io/gorm"
)

const tbName = "translations"

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.TranslateRepository {
	return mysqlRepo{db: db}
}

func (repo mysqlRepo) InsertTranslation(ctx context.Context, translation entity.Translation) error {
	return repo.db.Create(&translation).Error
}

func (repo mysqlRepo) GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error) {
	var data entity.Translation

	if err := repo.db.Table(tbName).
		Where(map[string]interface{}{
			"original_text": orgText,
			"source":        source,
			"destination":   dest,
		}).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Translation{}, entity.ErrNotFound
		}

		return entity.Translation{}, err
	}

	return data, nil
}

func (repo mysqlRepo) FindHistories(ctx context.Context) ([]entity.Translation, error) {
	var histories []entity.Translation

	if err := repo.db.Table(tbName).Find(&histories).Error; err != nil {
		return nil, err
	}

	return histories, nil
}
