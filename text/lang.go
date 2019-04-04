package text

import (
	"context"
	"google.golang.org/api/option"
)

type Text struct {
	translator *Translator
	t2p        *Text2Speech
	spch       *Speech
}

func New(ctx context.Context, opts ...option.ClientOption) (*Text, error) {
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

	return &Text{
		translator: tr,
		t2p:        t2p,
		spch:       spch,
	}, nil
}

func (l *Text) Translator() *Translator {
	return l.translator
}

func (l *Text) Text2Speech() *Text2Speech {
	return l.t2p
}

func (l *Text) Speech() *Speech {
	return l.spch
}

func (l *Text) Close() {
	l.translator.Close()
	l.t2p.Close()
	l.spch.Close()
}
