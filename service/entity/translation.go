package entity

import "errors"

var ErrNotFound = errors.New("not found")

type Translation struct {
	OriginalText string `json:"original_text"`
	Source       string `json:"source"`
	Destination  string `json:"destination"`
	ResultText   string `json:"result_text"`
}

func NewTranslation(orgText, source, dest, result string) Translation {
	return Translation{
		OriginalText: orgText,
		Source:       source,
		Destination:  dest,
		ResultText:   result,
	}
}

func (t *Translation) SetResultText(s string) {
	t.ResultText = s
}
