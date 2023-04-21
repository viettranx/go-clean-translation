package googlesv

import (
	"context"
	"github.com/Conight/go-googletrans"
	"go-clean-translation/service"
	"go-clean-translation/service/entity"
)

type googleTranslateAPI struct{}

func New() service.GoogleService {
	return googleTranslateAPI{}
}

func (api googleTranslateAPI) Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error) {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	trans := translator.New(conf)

	result, err := trans.Translate(orgText, source, dest)
	if err != nil {
		return entity.Translation{}, err
	}

	return entity.NewTranslation(orgText, source, dest, result.Text), nil
}
