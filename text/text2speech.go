package text

import (
	"cloud.google.com/go/texttospeech/apiv1"
	"context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"io"
)

type Text2Speech struct {
	speech *texttospeech.Client
}

func NewText2Speech(ctx context.Context, opts ...option.ClientOption) (*Text2Speech, error) {
	client, err := texttospeech.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Text2Speech{
		speech: client,
	}, nil
}

func (l *Text2Speech) Close() {
	_ = l.speech.Close()
}

// ListVoices lists the available text to speech voices.
func (l *Text2Speech) ListVoices(ctx context.Context) ([]*texttospeechpb.Voice, error) {
	// Performs the list voices request.
	resp, err := l.speech.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Voices, nil
}

func (l *Text2Speech) SynthesizeSpeech(ctx context.Context, dest io.Writer, opts ...Text2SpeechOption) error {
	resp, err := l.speech.SynthesizeSpeech(ctx, newText2SpeechRequest(opts...))
	if err != nil {
		return err
	}
	_, err = dest.Write(resp.AudioContent)
	if err != nil {
		return err
	}
	return nil
}

func (l *Text2Speech) ParseLanguage(targetLanguage string) (language.Tag, error) {
	return language.Parse(targetLanguage)
}
