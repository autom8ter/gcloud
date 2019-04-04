package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

type Document struct {
	fire *firestore.Client
}

func NewDocument(ctx context.Context, opts ...option.ClientOption) (*Document, error) {
	f, err := firestore.NewClient(ctx, "", opts...)
	if err != nil {
		return nil, err
	}
	return &Document{
		f,
	}, nil
}

func (v *Document) Close() {
	_ = v.fire.Close()
}

func (v *Document) Client() *firestore.Client {
	return v.fire
}
