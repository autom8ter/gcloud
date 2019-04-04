package text

import (
	"cloud.google.com/go/translate"
	"context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type Translation struct {
	Text       []string
	TargetLang language.Tag
	Opts       *translate.Options
}

type Translator struct {
	trans *translate.Client
}

func NewTranslator(ctx context.Context, opts ...option.ClientOption) (*Translator, error) {
	tr, err := translate.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &Translator{
		trans: tr,
	}, nil
}

func (t *Translator) Translate(ctx context.Context, opts ...TranslationOption) ([]translate.Translation, error) {
	tr := newTranslationRequest(opts...)
	return t.trans.Translate(ctx, tr.Text, tr.TargetLang, tr.Opts)
}

func (t *Translator) Client() *translate.Client {
	return t.trans
}

func (l *Translator) Close() {
	_ = l.trans.Close()
}

func (l *Translator) DetectLanguage(ctx context.Context, text string) (*translate.Detection, error) {
	lang, err := l.trans.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	return &lang[0][0], nil
}
