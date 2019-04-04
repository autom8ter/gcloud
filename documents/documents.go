package documents

import (
	"context"
	"google.golang.org/api/option"
)

type Documents struct {
}

func New(ctx context.Context, opts ...option.ClientOption) (*Documents, error) {
	return &Documents{}, nil
}

func (v *Documents) Close() {

}

func (v *Documents) Client() {

}
