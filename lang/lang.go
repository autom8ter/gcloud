package lang

import (
	"context"
	"google.golang.org/api/option"
)

type Lang struct {
	translator *Translator
	t2p        *Text2Speech
	spch       *Speech
}

func New(ctx context.Context, opts ...option.ClientOption) (*Lang, error) {
	tr, err := NewTranslator(ctx, opts...)
	if err != nil {
		return nil, err
	}
	t2p, err := NewText2Speech(ctx, opts...)
	if err != nil {
		return nil, err
	}
	spch, err := NewSpeech(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &Lang{
		translator: tr,
		t2p:        t2p,
		spch:       spch,
	}, nil
}

func (l *Lang) Translator() *Translator {
	return l.translator
}

func (l *Lang) Text2Speech() *Text2Speech {
	return l.t2p
}

func (l *Lang) Speech() *Speech {
	return l.spch
}

func (l *Lang) Close() {
	l.translator.Close()
	l.t2p.Close()
	l.spch.Close()
}
