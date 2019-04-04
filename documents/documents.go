package documents

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

type Documents struct {
	fire *firestore.Client
}

func New(ctx context.Context, opts ...option.ClientOption) (*Documents, error) {
	f, err := firestore.NewClient(ctx, "", opts...)
	if err != nil {
		return nil, err
	}
	return &Documents{
		f,
	}, nil
}

func (v *Documents) Close() {
 	_ = v.fire.Close()
}

func (v *Documents) Client() *firestore.Client{
	return v.fire
}

