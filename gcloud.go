package gcloud

import (
	"context"
	"github.com/autom8ter/gcloud/auth"
	"github.com/autom8ter/gcloud/cluster"
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

// GCP holds Google Cloud Platform Clients and carries some utility functions
// optional environmental variables: "GCLOUD_PROJECTID", "GCLOUD_SPANNER_DB" "GCLOUD_CLUSTER_MASTER" "GCLOUD_CLUSTER"
type GCP struct {
	Text    *text.Text
	PubSub  *pubsub.PubSub
	Vid     *video.Video
	Auth    *auth.Auth
	Storage *storage.Storage
	Trace   *trace.Trace
	Bots    *robots.Robot
	Kube    *cluster.Cluster
}

// New returns a new authenticated GCP instance from the provided api options
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error) {
	g := &GCP{}
	var err error
	var newErr error
	g.Text, newErr = text.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.PubSub, newErr = pubsub.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Vid, newErr = video.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Auth, newErr = auth.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Storage, newErr = storage.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Trace, newErr = trace.New()
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Bots, newErr = robots.New(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	g.Kube, newErr = cluster.New()
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return g, nil
}

// Close closes all clients
func (g *GCP) Close() {
	g.Text.Close()
	g.PubSub.Close()
	g.Vid.Close()
	g.Bots.Close()
	g.Auth.Close()
	g.Storage.Close()
	g.Kube.Close()
	g.Trace.Flush()
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

// DefaultClient returns an authenticated http client with the specified scopes
func (g *GCP) DefaultClient(ctx context.Context, scopes []string) (*http.Client, error) {
	return Client(ctx, scopes)
}
