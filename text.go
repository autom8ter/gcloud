package gcloud

import (
	"context"
	"fmt"
	"github.com/autom8ter/gcloud/text"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"io"
)

// ToSpanish returns the provided content in Spanish
func (g *GCP) ToSpanish(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.Spanish
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToSpanish")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

// ToSpanish returns the provided content in Chinese
func (g *GCP) ToChinese(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.Chinese
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToChinese")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

// ToFrench returns the provided content in French
func (g *GCP) ToFrench(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.French
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToFrench")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

// ToItalian returns the provided content in Italian
func (g *GCP) ToItalian(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.Italian
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToItalian")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

// ToGerman returns the provided content in German
func (g *GCP) ToGerman(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.German
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToGerman")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

// ToRussian returns the provided content in Russian
func (g *GCP) ToRussian(ctx context.Context, content []string) ([]string, error) {
	resp, err := g.txt.Translator().Translate(ctx, func(t *text.Translation) {
		t.Text = content
		t.TargetLang = language.Russian
	})
	if err != nil {
		return nil, errors.Wrap(err, "gcp.ToRussian")
	}
	trans := []string{}
	for _, v := range resp {
		trans = append(trans, v.Text)
	}
	return trans, nil
}

func (g *GCP) WriteTextToSpeechMP3(ctx context.Context, text string, w io.Writer) error {
	return g.txt.Text2Speech().SynthesizeSpeech(ctx, w, func(r *texttospeechpb.SynthesizeSpeechRequest) {
		r.Input.InputSource = &texttospeechpb.SynthesisInput_Text{Text: text}
		r.Voice = &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		}
		r.AudioConfig = &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		}
	})
}

func (g *GCP) WriteAudioTranscript(ctx context.Context, audioURL string, w io.Writer) error {
	resp, err := g.txt.Speech().Recognize(ctx, func(r *speechpb.RecognizeRequest) {
		r.Config = &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		}
		r.Audio = &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: audioURL},
		}

	})
	if err != nil {
		return err
	}
	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return nil
}
