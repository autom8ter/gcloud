package video

import (
	"cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Vision struct {
	annotator *vision.ImageAnnotatorClient
	prod *vision.ProductSearchClient
}

func NewVision(ctx context.Context, opts ...option.ClientOption) (*Vision, error) {
	v := &Vision{}
	var err error
	var newErr error
	v.annotator, newErr = vision.NewImageAnnotatorClient(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	v.prod, newErr = vision.NewProductSearchClient(ctx, opts ...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())

	}
	return v, err
}

func (v *Vision) Close() {
	_ = v.annotator.Close()
	_ = v.prod.Close()
}

func (v *Vision) Annotator() *vision.ImageAnnotatorClient {
	return v.annotator
}

func (v *Vision) ProductSearch() *vision.ProductSearchClient {
	return v.prod
}