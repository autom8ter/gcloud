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
	"encoding/json"
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
	validate "gopkg.in/go-playground/validator.v9"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

var valid = validate.New()

// Config is used to create a new GCP instance
type Config struct {
	Project   string   `validate:"required"`
	Scopes    []string `validate:"required"`
	InCluster bool
	SpannerDB string
	Options   []option.ClientOption `validate:"required"`
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
	ctx     context.Context       `validate:"required"`
	cfg     *Config               `validate:"required"`
	hTTP    *http.Client          `validate:"required"`
	trce    *stackdriver.Exporter `validate:"required"`
	clients *Clients
	svcs    *Services
}

// New returns a new authenticated GCP instance from the provided context and config
func New(ctx context.Context, cfg *Config) (*GCP, []error) {
	if ctx == nil {
		ctx = context.Background()
	}
	err := valid.Struct(cfg)
	if err != nil {
		panic(err.Error())
	}
	var stack []error
	cli, err := google.DefaultClient(ctx, cfg.Scopes...)
	if err != nil {
		stack = append(stack, eRR(err, cfg))
	}
	trc, err := stackdriver.NewExporter(stackdriver.Options{
		MonitoringClientOptions: cfg.Options,
		TraceClientOptions:      cfg.Options,
	})
	if err != nil {
		stack = append(stack, eRR(err, cfg))
	}

	g := &GCP{
		ctx:  ctx,
		cfg:  cfg,
		hTTP: cli,
		trce: trc,
	}
	err = valid.Struct(g)
	if err != nil {
		panic(err.Error())
	}
	if len(stack) != 0 {
		return g, stack
	}
	return g, nil
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

// WithServices adds the GCP Services to the GCP instance
func (g *GCP) WithServices() []error {
	var stack []error
	var err error
	g.svcs.Container, err = container.NewService(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.HealthCare, err = healthcare.NewService(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Calendar, err = calendar.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Blogger, err = blogger.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.CustomSearch, err = customsearch.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.ClassRoom, err = class.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Content, err = content.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.OSLogin, err = oslogin.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.People, err = people.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Photos, err = photos.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Predicion, err = prediction.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Redis, err = redis.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Config, err = run.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Sheets, err = sheets.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Slides, err = slides.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Tasks, err = tasks.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.YoutTube, err = youtube.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Docs, err = docs.NewService(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Jobs, err = jobs.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.svcs.Domains, err = plusdomains.New(g.hTTP)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}

	if len(stack) != 0 {
		return stack
	}
	return nil
}

// WithClients adds the GCP Clients to the GCP instance
func (g *GCP) WithClients() []error {
	var stack []error
	var err error
	g.clients.PubSub, err = pubsub.NewClient(g.ctx, g.cfg.Project, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.IAM, err = iam.NewService(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Storage, err = storage.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Spanner, err = spanner.NewClient(g.ctx, g.cfg.SpannerDB, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.DBAdmin, err = database.NewDatabaseAdminClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.FireStore, err = firestore.NewClient(g.ctx, g.cfg.Project, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.IOT, err = iot.NewDeviceManagerClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Kube, err = clients.NewKubernetesClientSet(g.cfg.InCluster)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Keys, err = kms.NewKeyManagementClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.ImageAnnotator, err = vision.NewImageAnnotatorClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.ImageProductSearch, err = vision.NewProductSearchClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.VideoIntelligence, err = videointelligence.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Text2Speech, err = texttospeech.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Speech, err = speech.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Translate, err = translate.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	g.clients.Language, err = language.NewClient(g.ctx, g.cfg.Options...)
	if err != nil {
		stack = append(stack, eRR(err, g.cfg))
	}
	if len(stack) != 0 {
		return stack
	}
	return nil
}

func eRR(err error, c *Config) error {
	return errors.Wrap(err, fmt.Sprintf("failed to create http client: %s", toPrettyJsonString(c)))
}

func toPrettyJsonString(obj interface{}) string {
	output, _ := json.MarshalIndent(obj, "", "  ")
	return fmt.Sprintf("%s", output)
}
