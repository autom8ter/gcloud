package vision

import (
	"context"
	"google.golang.org/api/option"
)

type Vision struct {
}

func New(ctx context.Context, opts ...option.ClientOption) (*Vision, error) {
	return &Vision{}, nil
}

func (v *Vision) Close() {

}

func (v *Vision) Client() {

}
