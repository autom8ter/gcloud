package vision

import (
	"context"
	"google.golang.org/api/option"
)

type VideoIntel struct {
}

func NewVideoIntel(ctx context.Context, opts ...option.ClientOption) (*VideoIntel, error) {
	return &VideoIntel{}, nil
}

func (v *VideoIntel) Close() {

}

func (v *VideoIntel) Client() {

}
