package video

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Video struct {
	vis *Vision
	int *Intelligence
}

func New(ctx context.Context, opts ...option.ClientOption) (*Video, error) {
	v := &Video{}
	var err error
	var newErr error
	v.int, newErr = NewIntelligence(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	v.vis, newErr = NewVision(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return v, err
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
