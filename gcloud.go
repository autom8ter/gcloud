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
	"google.golang.org/api/option"
	"google.golang.org/api/oslogin/v1"
	"google.golang.org/api/people/v1"
	photos "google.golang.org/api/photoslibrary/v1"
	"google.golang.org/api/prediction/v1.6"
	"google.golang.org/api/redis/v1"
	run "google.golang.org/api/runtimeconfig/v1"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
	"google.golang.org/api/tasks/v1"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/jobs/v3"
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
	Container    	*container.Service
	HealthCare   	*healthcare.Service
	Calendar     	*calendar.Service
	Blogger      	*blogger.Service
	CustomSearch 	*customsearch.Service
	ClassRoom    	*class.Service
	Content      	*content.APIService
	OSLogin      	*oslogin.Service
	People       	*people.Service
	Photos       	*photos.Service
	Predicion    	*prediction.Service
	Redis        	*redis.Service
	Config       	*run.Service
	Sheets       	*sheets.Service
	Slides       	*slides.Service
	Tasks        	*tasks.Service
	YoutTube     	*youtube.Service
	Docs         	*docs.Service
	Jobs 			*jobs.Service
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

// New returns a new authenticated GCP instance from the provided context and config
func New(ctx context.Context, cfg *Config) (*GCP, error) {
	var err error
	cli, newErr := google.DefaultClient(ctx, cfg.Scopes...)
	if newErr != nil {
		wrapErr(err, newErr, fmt.Sprintf("failed to create http client from scopes: %v", cfg.Scopes))
	}
	sub, newErr := pubsub.NewClient(ctx, cfg.Project, cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create pubsub client from options")
	}
	iAM, newErr := iam.NewService(ctx, cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create iam client from options")
	}
	strg, newErr := storage.NewClient(ctx, cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create storage client from options")
	}
	span, newErr := spanner.NewClient(ctx, cfg.SpannerDB, cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create spanner client from options")
	}
	db, newErr := database.NewDatabaseAdminClient(ctx, cfg.Options...)
	if err != nil {
		wrapErr(err, newErr, "failed to create database admin client from options")
	}
	fire, newErr := firestore.NewClient(ctx, cfg.Project, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create database admin client from options")
	}
	trc, newErr := stackdriver.NewExporter(stackdriver.Options{
		MonitoringClientOptions: cfg.Options,
		TraceClientOptions:      cfg.Options,
	})
	if newErr != nil {
		wrapErr(err, newErr, "failed to create stackdriver trace client from options")
	}
	bots, newErr := iot.NewDeviceManagerClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create iot device manager client from options")
	}
	kub, newErr := clients.NewKubernetesClientSet(cfg.InCluster)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create kubernetes client from options")
	}
	kys, newErr := kms.NewKeyManagementClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")
	}
	imgAnn, newErr := vision.NewImageAnnotatorClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")
	}
	imgProd, newErr := vision.NewProductSearchClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")

	}
	intel, err := videointelligence.NewClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create video intelligence client from options")

	}
	t2p, err := texttospeech.NewClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create text2speech client from options")
	}
	spch, newErr := speech.NewClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create speech client from options")
	}
	tr, newErr := translate.NewClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create translation client from options")
	}

	lang, newErr := language.NewClient(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create language client from options")
	}
	con, newErr := container.NewService(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create container client from google default http  client")
	}
	hth, newErr := healthcare.NewService(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create healthcare client from options")
	}
	cal, newErr := calendar.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create calendar client from google default http client")
	}
	blg, newErr := blogger.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create blogger client from google default http client")
	}
	sch, newErr := customsearch.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create custom search client from google default http client")
	}
	cls, newErr := class.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create classroom client from google default http client")
	}
	cont, newErr := content.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create content client from google default http client")
	}
	login, newErr := oslogin.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create os login client from google default http client")
	}
	ppl, newErr := people.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create people client from google default http client")
	}
	pho, newErr := photos.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create photos client from google default http client")
	}
	pred, newErr := prediction.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create prediction client from google default http client")
	}
	red, newErr := redis.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create redis client from google default http client")
	}
	runtime, newErr := run.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create runtime client from google default http client")
	}
	shts, newErr := sheets.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create sheets client from google default http client")
	}
	slds, newErr := slides.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create slides client from google default http client")
	}
	tsks, newErr := tasks.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create tasks client from google default http client")
	}
	tube, newErr := youtube.New(cli)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create youtube client from google default http client")
	}
	dcs, newErr := docs.NewService(ctx, cfg.Options...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create docs client from google default http client")
	}

	if err != nil {
		return &GCP{
			ctx:  ctx,
			cfg:  cfg,
			hTTP: cli,
			trce: trc,
			clients: &Clients{
				PubSub:             sub,
				IAM:                iAM,
				Storage:            strg,
				Spanner:            span,
				DBAdmin:            db,
				FireStore:          fire,
				IOT:                bots,
				Kube:               kub,
				Keys:               kys,
				ImageAnnotator:     imgAnn,
				ImageProductSearch: imgProd,
				VideoIntelligence:  intel,
				Speech:             spch,
				Text2Speech:        t2p,
				Translate:          tr,
				Language:           lang,
			},
			svcs: &Services{
				Container:    con,
				HealthCare:   hth,
				Calendar:     cal,
				Blogger:      blg,
				CustomSearch: sch,
				ClassRoom:    cls,
				Content:      cont,
				OSLogin:      login,
				People:       ppl,
				Photos:       pho,
				Predicion:    pred,
				Redis:        red,
				Config:       runtime,
				Sheets:       shts,
				Slides:       slds,
				Tasks:        tsks,
				YoutTube:     tube,
				Docs:         dcs,
			},
		}, err
	}
	return &GCP{
		ctx:  ctx,
		cfg:  cfg,
		hTTP: cli,
		trce: trc,
		clients: &Clients{
			PubSub:             sub,
			IAM:                iAM,
			Storage:            strg,
			Spanner:            span,
			DBAdmin:            db,
			FireStore:          fire,
			IOT:                bots,
			Kube:               kub,
			Keys:               kys,
			ImageAnnotator:     imgAnn,
			ImageProductSearch: imgProd,
			VideoIntelligence:  intel,
			Speech:             spch,
			Text2Speech:        t2p,
			Translate:          tr,
			Language:           lang,
		},
		svcs: &Services{
			Container:    con,
			HealthCare:   hth,
			Calendar:     cal,
			Blogger:      blg,
			CustomSearch: sch,
			ClassRoom:    cls,
			Content:      cont,
			OSLogin:      login,
			People:       ppl,
			Photos:       pho,
			Predicion:    pred,
			Redis:        red,
			Config:       runtime,
			Sheets:       shts,
			Slides:       slds,
			Tasks:        tsks,
			YoutTube:     tube,
			Docs:         dcs,
		},
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

// Context returns the context used to create the GCP instance
func (g *GCP) Context() context.Context {
	return g.ctx
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

// Trace returns a stackdriver exporter
func (g *GCP) Trace() *stackdriver.Exporter {
	return g.trce
}

// HTTP returns a google default HTTP client
func (g *GCP) HTTP() *http.Client {
	return g.hTTP
}

func wrapErr(err error, newErr error, msg string) {
	err = errors.Wrap(newErr, msg)
}
