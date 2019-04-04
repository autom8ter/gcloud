package video

import (
	"context"
	"google.golang.org/api/option"
)

type Intelligence struct {
}

func NewIntelligence(ctx context.Context, opts ...option.ClientOption) (*Intelligence, error) {
	return &Intelligence{}, nil
}

func (v *Intelligence) Close() {

}

func (v *Intelligence) Client() {

}
