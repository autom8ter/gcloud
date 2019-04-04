package gcloud

import (
	"context"
	"github.com/autom8ter/gcloud/lang"
	"github.com/autom8ter/gcloud/pubsub"
	"github.com/autom8ter/gcloud/vision"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"io"
)

// Func is used to run a function using a GCP object (see GCP.Execute)
type Func func(g *GCP) error

// GCP holds Google Cloud Platform Clients and carries some utility functions
type GCP struct {
	lng *lang.Lang
	sub *pubsub.PubSub
	vis *vision.Vision
}

// New returns a new authenticated GCP instance from the provided api options
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error) {
	g := &GCP{}
	var err error
	var newErr error
	g.lng, newErr = lang.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.sub, err = pubsub.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.vis, err = vision.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return g, nil
}

// Lang returns a client used for GCP text2speech, translation, and speech services
func (g *GCP) Lang() *lang.Lang {
	return g.lng
}

// PubSub returns a client used for GCP pubsub
func (g *GCP) PubSub() *pubsub.PubSub {
	return g.sub
}

// PubSub returns a client used for GCP video intelligence and computer vision
func (g *GCP) Vision() *vision.Vision {
	return g.vis
}

// Close closes all clients
func (g *GCP) Close() {
	g.lng.Close()
	g.sub.Close()
}

// JSON formats an object and turns it into JSON bytes
func (g *GCP) JSON(obj interface{}) []byte {
	return toJSON(obj)
}

// XML formats an object and turns it into XML bytes
func (g *GCP) XML(obj interface{}) []byte {
	return toXML(obj)
}

// YAML formats an object and turns it into YAML bytes
func (g *GCP) YAML(obj interface{}) []byte {
	return toYAML(obj)
}

// Proto formats an object and turns it into  Proto bytes
func (g *GCP) Proto(m proto.Message) []byte {
	return toProto(m)
}

// Render uses html/template along with the sprig funcmap functions to render a strings to an io writer ref: https://github.com/Masterminds/sprig
func (g *GCP) Render(text string, data interface{}, w io.Writer) error {
	return render(text, data, w)
}

// Execute runs all functions and returns a wrapped error
func (g *GCP) Execute(fns ...Func) error {
	var err error
	for _, f := range fns {
		if newErr := f(g); newErr != nil {
			err = errors.Wrap(err, newErr.Error())
		}
	}
	return err
}
