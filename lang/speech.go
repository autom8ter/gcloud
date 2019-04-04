package lang

import (
	"fmt"
	"google.golang.org/api/option"
	// [START imports]
	"context"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"

	"cloud.google.com/go/speech/apiv1"
)

type Speech struct {
	spch *speech.Client
}

func NewSpeech(ctx context.Context, opts ...option.ClientOption) (*Speech, error) {
	client, err := speech.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("New Speech Client: %v", err)
	}
	return &Speech{
		spch: client,
	}, nil
}

func (s *Speech) Close() {
	_ = s.spch.Close()
}

func (s *Speech) Client() *speech.Client {
	return s.spch
}

func (s *Speech) Recognize(ctx context.Context, opts ...RecognizeOption) (*speechpb.RecognizeResponse, error) {
	return s.spch.Recognize(ctx, newRecognizeRequest(opts...))
}
