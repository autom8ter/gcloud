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
	"github.com/autom8ter/gcloud/clients"
	"github.com/autom8ter/objectify"
	"github.com/hashicorp/go-multierror"
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

var tool = objectify.New()

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
	Prediction   *prediction.Service
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
	PubSub             *pubsub.Client `validate:"required"`
	IAM                *iam.Service
	Storage            *storage.Client `validate:"required"`
	Spanner            *spanner.Client
	DBAdmin            *database.DatabaseAdminClient
	FireStore          *firestore.Client `validate:"required"`
	IOT                *iot.DeviceManagerClient
	Kube               *kubernetes.Clientset `validate:"required"`
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
	ctx     context.Context `validate:"required"`
	cfg     *Config         `validate:"required"`
	httP    *http.Client    `validate:"required"`
	clients *Clients        `validate:"required"`
	svcs    *Services       `validate:"required"`
	err     *multierror.Error
}

// New returns a new authenticated GCP instance from the provided context and config
func New(ctx context.Context, cfg *Config) *GCP {
	var errs []error
	if ctx == nil {
		ctx = context.Background()
	}
	err := tool.Validate(cfg)
	if err != nil {
		panic("validation error: " + err.Error())
	}
	cli, err := google.DefaultClient(ctx, cfg.Scopes...)
	if err != nil {
		errs = append(errs, err)
	}

	sub, err := pubsub.NewClient(ctx, cfg.Project, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	pol, err := iam.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	strg, err := storage.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	span, err := spanner.NewClient(ctx, cfg.SpannerDB, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	db, err := database.NewDatabaseAdminClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	fire, err := firestore.NewClient(ctx, cfg.Project, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	bots, err := iot.NewDeviceManagerClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	kub, err := clients.NewKubernetesClientSet(cfg.InCluster)
	if err != nil {
		errs = append(errs, err)
	}
	kys, err := kms.NewKeyManagementClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	ann, err := vision.NewImageAnnotatorClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	prodsch, err := vision.NewProductSearchClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	intel, err := videointelligence.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	t2p, err := texttospeech.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	spch, err := speech.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	trans, err := translate.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	lang, err := language.NewClient(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}

	tain, err := container.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	hcare, err := healthcare.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	cal, err := calendar.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	blg, err := blogger.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	sch, err := customsearch.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	cRoom, err := class.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	cont, err := content.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	login, err := oslogin.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	peeps, err := people.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	phot, err := photos.New(cli)
	if err != nil {
		errs = append(errs, err)
	}
	pred, err := prediction.New(cli)
	if err != nil {
		errs = append(errs, err)
	}
	red, err := redis.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	runtime, err := run.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	shts, err := sheets.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	slds, err := slides.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	tsks, err := tasks.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	tube, err := youtube.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	dcs, err := docs.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	jbs, err := jobs.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	doms, err := plusdomains.NewService(ctx, cfg.Options...)
	if err != nil {
		errs = append(errs, err)
	}
	multiErr := appendErr(errs[0], errs[1:]...)
	g := &GCP{
		ctx:  ctx,
		cfg:  cfg,
		httP: cli,
		err:  multiErr,
		svcs: &Services{
			Container:    tain,
			HealthCare:   hcare,
			Calendar:     cal,
			Blogger:      blg,
			CustomSearch: sch,
			ClassRoom:    cRoom,
			Content:      cont,
			OSLogin:      login,
			People:       peeps,
			Photos:       phot,
			Prediction:   pred,
			Redis:        red,
			Config:       runtime,
			Sheets:       shts,
			Slides:       slds,
			Tasks:        tsks,
			YoutTube:     tube,
			Docs:         dcs,
			Jobs:         jbs,
			Domains:      doms,
		},
		clients: &Clients{
			PubSub:             sub,
			IAM:                pol,
			Storage:            strg,
			Spanner:            span,
			DBAdmin:            db,
			FireStore:          fire,
			IOT:                bots,
			Kube:               kub,
			Keys:               kys,
			ImageAnnotator:     ann,
			ImageProductSearch: prodsch,
			VideoIntelligence:  intel,
			Speech:             spch,
			Text2Speech:        t2p,
			Translate:          trans,
			Language:           lang,
		},
	}
	err = tool.Validate(g)
	if err != nil {
		panic("validation error: " + err.Error())
	}
	return g
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
}

// FromContext returns the value the context is holding from the given key
func (g *GCP) FromContext(key interface{}) interface{} {
	return g.ctx.Value(key)
}

func (g *GCP) Context() context.Context {
	return g.ctx
}

func (g *GCP) Error() error {
	return g.err.ErrorOrNil()
}

// Configuration returns the config used to create the GCP instance
func (g *GCP) Configuration() *Config {
	return g.cfg
}

// Services returns an authenticated GCP ServiceSet
func (g *GCP) Services() *Services {
	return g.svcs
}

// Clients returns an authenticated GCP ClientSet
func (g *GCP) Clients() *Clients {
	return g.clients
}

// HTTP returns a google default HTTP client
func (g *GCP) HTTP() *http.Client {
	return g.httP
}

func appendErr(base error, toAppend ...error) *multierror.Error {
	return multierror.Append(base, toAppend...)
}
