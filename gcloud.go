package gcloud

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/iot/apiv1"
	"cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/translate"
	"cloud.google.com/go/videointelligence/apiv1"
	"cloud.google.com/go/vision/apiv1"
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"fmt"
	"github.com/autom8ter/gcloud/clients"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/calendar/v3"
	class "google.golang.org/api/classroom/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/content/v2.1"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/healthcare/v1alpha2"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/jobs/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/oslogin/v1"
	"google.golang.org/api/people/v1"
	photos "google.golang.org/api/photoslibrary/v1"
	"google.golang.org/api/plusdomains/v1"
	"google.golang.org/api/prediction/v1.6"
	"google.golang.org/api/redis/v1"
	run "google.golang.org/api/runtimeconfig/v1"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
	"google.golang.org/api/tasks/v1"
	"google.golang.org/api/youtube/v3"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

// Config is used to create a new GCP instance
type Config struct {
	Project   string
	Scopes    []string
	InCluster bool
	SpannerDB string
	Options   []option.ClientOption
}

// GCP ServiceSet. Make sure to pass the necessary scopes in your config to successfully initialize services.
type Services struct {
	Container    *container.Service
	HealthCare   *healthcare.Service
	Calendar     *calendar.Service
	Blogger      *blogger.Service
	CustomSearch *customsearch.Service
	ClassRoom    *class.Service
	Content      *content.APIService
	OSLogin      *oslogin.Service
	People       *people.Service
	Photos       *photos.Service
	Predicion    *prediction.Service
	Redis        *redis.Service
	Config       *run.Service
	Sheets       *sheets.Service
	Slides       *slides.Service
	Tasks        *tasks.Service
	YoutTube     *youtube.Service
	Docs         *docs.Service
	Jobs         *jobs.Service
	Domains      *plusdomains.Service
}

// GCP ClientSet
type Clients struct {
	PubSub             *pubsub.Client
	IAM                *iam.Service
	Storage            *storage.Client
	Spanner            *spanner.Client
	DBAdmin            *database.DatabaseAdminClient
	FireStore          *firestore.Client
	IOT                *iot.DeviceManagerClient
	Kube               *kubernetes.Clientset
	Keys               *kms.KeyManagementClient
	ImageAnnotator     *vision.ImageAnnotatorClient
	ImageProductSearch *vision.ProductSearchClient
	VideoIntelligence  *videointelligence.Client
	Speech             *speech.Client
	Text2Speech        *texttospeech.Client
	Translate          *translate.Client
	Language           *language.Client
}

// GCP holds Google Cloud Platform Clients and Services
type GCP struct {
	ctx     context.Context
	cfg     *Config
	hTTP    *http.Client
	trce    *stackdriver.Exporter
	clients *Clients
	svcs    *Services
}

type HandlerFunc func(*GCP) error

// New returns a new authenticated GCP instance from the provided context and config
func New(ctx context.Context, cfg *Config) (*GCP, error) {
	var err error
	cli, newErr := google.DefaultClient(ctx, cfg.Scopes...)
	if newErr != nil {
		wrapErr(err, newErr, fmt.Sprintf("failed to create http client from scopes: %v", cfg.Scopes))
	}
	trc, newErr := stackdriver.NewExporter(stackdriver.Options{
		MonitoringClientOptions: cfg.Options,
		TraceClientOptions:      cfg.Options,
	})
	if newErr != nil {
		wrapErr(err, newErr, "failed to create stackdriver trace client from options")
	}

	if err != nil {
		return &GCP{
			ctx:  ctx,
			cfg:  cfg,
			hTTP: cli,
			trce: trc,
		}, err
	}
	return &GCP{
		ctx:  ctx,
		cfg:  cfg,
		hTTP: cli,
		trce: trc,
	}, nil
}

// Close closes all clients
func (g *GCP) Close() {
	_ = g.clients.PubSub.Close()
	_ = g.clients.Storage.Close()
	_ = g.clients.DBAdmin.Close()
	_ = g.clients.Speech.Close()
	_ = g.clients.FireStore.Close()
	_ = g.clients.Text2Speech.Close()
	_ = g.clients.ImageProductSearch.Close()
	_ = g.clients.Translate.Close()
	_ = g.clients.Keys.Close()
	_ = g.clients.VideoIntelligence.Close()
	g.trce.Flush()
}

// FromContext returns the value the context is holding from the given key
func (g *GCP) FromContext(key interface{}) interface{} {
	return g.ctx.Value(key)
}

func (g *GCP) Context() context.Context {
	return g.ctx
}

// Configuration returns the config used to create the GCP instance
func (g *GCP) Configuration() *Config {
	if g.cfg == nil {
		panic("configuration is uninitialized- use gcloud.New to initialize the GCP instance")
	}
	return g.cfg
}

// Services returns an authenticated GCP ServiceSet
func (g *GCP) Services() *Services {
	if g.svcs == nil {
		panic("services are uninitialized- use WithServices to add the GCP service set")
	}
	return g.svcs
}

// Clients returns an authenticated GCP ClientSet
func (g *GCP) Clients() *Clients {
	if g.clients == nil {
		panic("clients are uninitialized- use WithClients to add the GCP client set")
	}
	return g.clients
}

// Trace returns a stackdriver exporter
func (g *GCP) Trace() *stackdriver.Exporter {
	if g.trce == nil {
		panic("exporter is uninitialized- use gcloud.New to initialize the GCP instance")
	}
	return g.trce
}

// HTTP returns a google default HTTP client
func (g *GCP) HTTP() *http.Client {
	if g.trce == nil {
		panic("http client is uninitialized- use gcloud.New to initialize the GCP instance")
	}
	return g.hTTP
}

func wrapErr(err error, newErr error, msg string) {
	err = errors.Wrap(newErr, msg)
}

// WithServices adds the GCP Services to the GCP instance
func (g *GCP) WithServices() error {
	var err error
	var newErr error
	g.svcs.Container, newErr = container.NewService(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create container client")
	}
	g.svcs.HealthCare, newErr = healthcare.NewService(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create healthcare client")
	}
	g.svcs.Calendar, newErr = calendar.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create calendar client")
	}
	g.svcs.Blogger, newErr = blogger.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create blogger client")
	}
	g.svcs.CustomSearch, newErr = customsearch.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create custom search client")
	}
	g.svcs.ClassRoom, newErr = class.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create classroom client")
	}
	g.svcs.Content, newErr = content.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create content client")
	}
	g.svcs.OSLogin, newErr = oslogin.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create os login client")
	}
	g.svcs.People, newErr = people.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create people client")
	}
	g.svcs.Photos, newErr = photos.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create photos client")
	}
	g.svcs.Predicion, newErr = prediction.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create prediction client")
	}
	g.svcs.Redis, newErr = redis.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create redis client")
	}
	g.svcs.Config, newErr = run.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create runtime client")
	}
	g.svcs.Sheets, newErr = sheets.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create sheets client")
	}
	g.svcs.Slides, newErr = slides.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create slides client")
	}
	g.svcs.Tasks, newErr = tasks.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create tasks client")
	}
	g.svcs.YoutTube, newErr = youtube.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create youtube client")
	}
	g.svcs.Docs, newErr = docs.NewService(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create docs client")
	}
	g.svcs.Jobs, newErr = jobs.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create jobs client")
	}
	g.svcs.Domains, newErr = plusdomains.New(g.hTTP)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create domains client")
	}
	return err
}

// WithClients adds the GCP Clients to the GCP instance
func (g *GCP) WithClients() error {
	var err error
	var newErr error
	g.clients.PubSub, newErr = pubsub.NewClient(g.ctx, g.cfg.Project, g.cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create pubsub client")
	}
	g.clients.IAM, newErr = iam.NewService(g.ctx, g.cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create iam client")
	}
	g.clients.Storage, newErr = storage.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create storage client")
	}
	g.clients.Spanner, newErr = spanner.NewClient(g.ctx, g.cfg.SpannerDB, g.cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create spanner client")
	}
	g.clients.DBAdmin, newErr = database.NewDatabaseAdminClient(g.ctx, g.cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create database admin client")
	}
	g.clients.FireStore, newErr = firestore.NewClient(g.ctx, g.cfg.Project, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create database admin client")
	}

	if newErr != nil {
		wrapErr(err, newErr, "failed to create stackdriver trace client")
	}
	g.clients.IOT, newErr = iot.NewDeviceManagerClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create iot device manager client")
	}
	g.clients.Kube, newErr = clients.NewKubernetesClientSet(g.cfg.InCluster)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create kubernetes client")
	}
	g.clients.Keys, newErr = kms.NewKeyManagementClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client")
	}
	g.clients.ImageAnnotator, newErr = vision.NewImageAnnotatorClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client")
	}
	g.clients.ImageProductSearch, newErr = vision.NewProductSearchClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client")

	}
	g.clients.VideoIntelligence, err = videointelligence.NewClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create video intelligence client")

	}
	g.clients.Text2Speech, err = texttospeech.NewClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create text2speech client")
	}
	g.clients.Speech, newErr = speech.NewClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create speech client")
	}
	g.clients.Translate, newErr = translate.NewClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create translation client")
	}

	g.clients.Language, newErr = language.NewClient(g.ctx, g.cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create language client")
	}
	return err
}

func (g *GCP) Execute(fns ...HandlerFunc) error {
	for _, fn := range fns {
		err := fn(g)
		if err != nil {
			return err
		}
	}
	return nil
}
