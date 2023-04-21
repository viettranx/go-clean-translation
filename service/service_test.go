package service

import (
	"context"
	"errors"
	"go-clean-translation/service/entity"
	"testing"
)

var (
	translationOne = entity.Translation{
		OriginalText: "Apple",
		Source:       "en",
		Destination:  "vi",
		ResultText:   "trái táo",
	}

	translationTwo = entity.Translation{
		OriginalText: "trái táo",
		Source:       "vi",
		Destination:  "en",
		ResultText:   "Apple",
	}

	emptyTranslation = entity.Translation{}

	ErrGoogle = errors.New("cannot translate")
	ErrDB     = errors.New("db error")
)

type mockGoogleService struct{}

func (mockGoogleService) Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error) {
	dataSet := []entity.Translation{translationTwo}

	for _, t := range dataSet {
		if t.Source == source && t.OriginalText == orgText && t.Destination == dest {
			return t, nil
		}
	}

	return emptyTranslation, ErrGoogle
}

type mockRepository struct{}

func (mockRepository) GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error) {
	dataSet := []entity.Translation{translationOne}

	for _, t := range dataSet {
		if t.Source == source && t.OriginalText == orgText && t.Destination == dest {
			return t, nil
		}
	}

	return emptyTranslation, ErrDB
}

func (mockRepository) FindHistories(ctx context.Context) ([]entity.Translation, error) {
	return []entity.Translation{translationOne, translationTwo}, nil
}

func (mockRepository) InsertTranslation(ctx context.Context, translation entity.Translation) error {
	return nil
}

func TestService_Translate(t *testing.T) {
	testTable := []struct {
		orgText     string
		source      string
		dest        string
		expectError error
		expectData  entity.Translation
	}{
		{
			orgText:     translationOne.OriginalText,
			source:      translationOne.Source,
			dest:        translationOne.Destination,
			expectError: nil,
			expectData:  translationOne,
		},
		{
			orgText:     translationTwo.OriginalText,
			source:      translationTwo.Source,
			dest:        translationTwo.Destination,
			expectError: nil,
			expectData:  translationTwo,
		},
		{
			orgText:     "other text",
			source:      "auto",
			dest:        "any",
			expectError: ErrGoogle,
			expectData:  emptyTranslation,
		},
	}

	// setup
	repo := mockRepository{}
	googleSv := mockGoogleService{}
	translateSv := NewService(repo, googleSv)

	// Run test

	for _, item := range testTable {
		realData, realErr := translateSv.Translate(context.Background(), item.orgText, item.source, item.dest)

		if realData != item.expectData {
			t.Errorf("Failed. Expected %v but received %v", item.expectData, realData)
		}

		if realErr != item.expectError {
			t.Errorf("Failed. Expected %v but received %v", item.expectData, realData)
		}
	}

}
