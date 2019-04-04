package video

import (
	"context"
	"google.golang.org/api/option"
)

type Video struct {
	vis *Vision
	int *Intelligence
}

func New(ctx context.Context, opts ...option.ClientOption) (*Video, error) {
	return &Video{}, nil
}

func (v *Video) Intelligence() *Intelligence {
	return v.int
}

func (v *Video) Vision() *Vision {
	return v.vis
}

func (v *Video) Close() {
	v.vis.Close()
	v.int.Close()
}
