package text

import (
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type RecognizeOption func(r *speechpb.RecognizeRequest)

func newRecognizeRequest(opts ...RecognizeOption) *speechpb.RecognizeRequest {
	r := &speechpb.RecognizeRequest{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type Text2SpeechOption func(r *texttospeechpb.SynthesizeSpeechRequest)

func newText2SpeechRequest(opts ...Text2SpeechOption) *texttospeechpb.SynthesizeSpeechRequest {
	r := &texttospeechpb.SynthesizeSpeechRequest{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type TranslationOption func(t *Translation)

func newTranslationRequest(opts ...TranslationOption) *Translation {
	t := &Translation{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}
