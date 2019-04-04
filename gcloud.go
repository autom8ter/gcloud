package gcloud

import (
	"context"
	"github.com/autom8ter/gcloud/auth"
	"github.com/autom8ter/gcloud/pubsub"
	"github.com/autom8ter/gcloud/robots"
	"github.com/autom8ter/gcloud/storage"
	"github.com/autom8ter/gcloud/text"
	"github.com/autom8ter/gcloud/trace"
	"github.com/autom8ter/gcloud/video"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"io"
	"net/http"
)

// HandlerFunc is used to run a function using a GCP object (see GCP.Execute)
// Creating a HandlerFunc is easy...
/*
	func NewHandlerFunc() HandlerFunc {
		return func(g *GCP) error {

		this is similar to http.HandlerFunc...

		return nil
	}}
*/
type HandlerFunc func(g *GCP) error

// GCP holds Google Cloud Platform Clients and carries some utility functions
type GCP struct {
	txt  *text.Text
	sub  *pubsub.PubSub
	vid  *video.Video
	ath  *auth.Auth
	strg *storage.Storage
	trc  *trace.Trace
	bots *robots.Robot
}

// New returns a new authenticated GCP instance from the provided api options
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error) {
	g := &GCP{}
	var err error
	var newErr error
	g.txt, newErr = text.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.sub, newErr = pubsub.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.vid, newErr = video.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.ath, newErr = auth.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.strg, newErr = storage.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.trc, err = trace.New()
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.bots, err = robots.New(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return g, nil
}

// Trace returns a registered stackdriver exporter
func (g *GCP) Trace() *trace.Trace {
	return g.trc
}

// Text returns a client used for common text operations: GCP text2speech, translation, and speech services
func (g *GCP) Text() *text.Text {
	return g.txt
}

// PubSub returns a client used for GCP pubsub
func (g *GCP) PubSub() *pubsub.PubSub {
	return g.sub
}

// Video returns a client used for GCP video intelligence and computer vision
func (g *GCP) Video() *video.Video {
	return g.vid
}

// Auth returns a client used for GCP key management and IAM
func (g *GCP) Auth() *auth.Auth {
	return g.ath
}

// Auth returns a client used for GCP key management and IAM
func (g *GCP) Robots() *robots.Robot {
	return g.bots
}

// Blob returns a client used for GCP blob storage
func (g *GCP) Storage() *storage.Storage {
	return g.strg
}

// Close closes all clients
func (g *GCP) Close() {
	g.txt.Close()
	g.sub.Close()
	g.vid.Close()
}

// JSON formats an object and turns it into JSON bytes
func (g *GCP) JSON(obj interface{}) []byte {
	return JSON(obj)
}

// XML formats an object and turns it into XML bytes
func (g *GCP) XML(obj interface{}) []byte {
	return XML(obj)
}

// YAML formats an object and turns it into YAML bytes
func (g *GCP) YAML(obj interface{}) []byte {
	return YAML(obj)
}

// Proto formats an object and turns it into  Proto bytes
func (g *GCP) Proto(m proto.Message) []byte {
	return Proto(m)
}

// Render uses html/template along with the sprig funcmap functions to render a strings to an io writer ref: https://github.com/Masterminds/sprig
func (g *GCP) Render(text string, data interface{}, w io.Writer) error {
	return Render(text, data, w)
}

// Execute runs all functions and returns a wrapped error
func (g *GCP) Execute(fns ...HandlerFunc) error {
	var err error
	for _, f := range fns {
		if newErr := f(g); newErr != nil {
			err = errors.Wrap(err, newErr.Error())
		}
	}
	return err
}

// DefaultClient returns an authenticated http client with the specified scopes
func (g *GCP) DefaultClient(ctx context.Context, scopes []string) (*http.Client, error) {
	return Client(ctx, scopes)
}
