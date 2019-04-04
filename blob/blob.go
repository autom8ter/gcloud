package blob

import (
	"context"
	"google.golang.org/api/option"
)

type Blob struct {
}

func New(ctx context.Context, opts ...option.ClientOption) (*Blob, error) {
	return &Blob{}, nil
}

func (v *Blob) Close() {

}

func (v *Blob) Client() {

}
