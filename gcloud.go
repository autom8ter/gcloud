package gcloud

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/iot/apiv1"
	"cloud.google.com/go/kms/apiv1"
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
	"github.com/autom8ter/gcloud/clients"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"os"
)

// GCP holds Google Cloud Platform Clients and carries some utility functions
// optional environmental variables: "GCLOUD_PROJECTID", "GCLOUD_SPANNER_DB" "GCLOUD_CLUSTER_MASTER" "GCLOUD_CLUSTER", "GCLOUD_INCLUSTER"
type GCP struct {
	PubSub             *pubsub.Client
	IAM                *iam.Service
	Storage            *storage.Client
	Spanner            *spanner.Client
	DBAdmin            *database.DatabaseAdminClient
	FireStore          *firestore.Client
	Trace              *stackdriver.Exporter
	IOT                *iot.DeviceManagerClient
	Kube               *kubernetes.Clientset
	Keys               *kms.KeyManagementClient
	ImageAnnotator     *vision.ImageAnnotatorClient
	ImageProductSearch *vision.ProductSearchClient
	VideoIntelligence  *videointelligence.Client
	Speech             *speech.Client
	Text2Speech        *texttospeech.Client
	Translate          *translate.Client
}

// New returns a new authenticated GCP instance from the provided api options GCLOUD_PROJECTID, GCLOUD_SPANNERDB,
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error) {
	var err error
	sub, newErr := pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECTID"), opts...)
	if err != nil {
		wrapErr(err, newErr, "failed to create pubsub client from options")
	}
	iAM, newErr := iam.NewService(ctx, opts...)
	if err != nil {
		wrapErr(err, newErr, "failed to create iam client from options")
	}
	strg, newErr := storage.NewClient(ctx, opts...)
	if err != nil {
		wrapErr(err, newErr, "failed to create storage client from options")
	}
	span, newErr := spanner.NewClient(ctx, os.Getenv("GCLOUD_SPANNERDB"), opts...)
	if err != nil {
		wrapErr(err, newErr, "failed to create spanner client from options")
	}
	db, newErr := database.NewDatabaseAdminClient(ctx, opts...)
	if err != nil {
		wrapErr(err, newErr, "failed to create database admin client from options")
	}
	fire, newErr := firestore.NewClient(ctx, os.Getenv("GCLOUD_PROJECTID"), opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create database admin client from options")
	}
	trc, newErr := stackdriver.NewExporter(stackdriver.Options{
		MonitoringClientOptions: opts,
		TraceClientOptions:      opts,
	})
	if newErr != nil {
		wrapErr(err, newErr, "failed to create stackdriver trace client from options")
	}
	bots, newErr := iot.NewDeviceManagerClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create iot device manager client from options")
	}
	kub, newErr := clients.NewKubernetesClientSet(inCluster())
	if newErr != nil {
		wrapErr(err, newErr, "failed to create kubernetes client from options")
	}
	kys, newErr := kms.NewKeyManagementClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")
	}
	imgAnn, newErr := vision.NewImageAnnotatorClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")
	}
	imgProd, newErr := vision.NewProductSearchClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create key management client from options")

	}
	intel, err := videointelligence.NewClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create video intelligence client from options")

	}
	t2p, err := texttospeech.NewClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create text2speech client from options")
	}
	spch, err := speech.NewClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create speech client from options")
	}
	tr, err := translate.NewClient(ctx, opts...)
	if newErr != nil {
		wrapErr(err, newErr, "failed to create translation client from options")
	}
	if err != nil {
		return &GCP{
			PubSub:             sub,
			IAM:                iAM,
			Storage:            strg,
			Spanner:            span,
			DBAdmin:            db,
			FireStore:          fire,
			Trace:              trc,
			IOT:                bots,
			Kube:               kub,
			Keys:               kys,
			ImageAnnotator:     imgAnn,
			ImageProductSearch: imgProd,
			VideoIntelligence:  intel,
			Speech:             spch,
			Text2Speech:        t2p,
			Translate:          tr,
		}, err
	}
	return &GCP{
		PubSub:             sub,
		IAM:                iAM,
		Storage:            strg,
		Spanner:            span,
		DBAdmin:            db,
		FireStore:          fire,
		Trace:              trc,
		IOT:                bots,
		Kube:               kub,
		Keys:               kys,
		ImageAnnotator:     imgAnn,
		ImageProductSearch: imgProd,
		VideoIntelligence:  intel,
		Speech:             spch,
		Text2Speech:        t2p,
		Translate:          tr,
	}, nil
}

// Close closes all clients
func (g *GCP) Close() {
	g.PubSub.Close()
	g.Storage.Close()
	g.DBAdmin.Close()
	g.Speech.Close()
	g.FireStore.Close()
	g.Text2Speech.Close()
	g.ImageProductSearch.Close()
	g.Translate.Close()
	g.Keys.Close()
	g.VideoIntelligence.Close()
	g.Trace.Flush()
}

// DefaultClient returns an authenticated http client with the specified scopes
func (g *GCP) DefaultClient(ctx context.Context, scopes []string) (*http.Client, error) {
	return google.DefaultClient(ctx, scopes...)
}

func wrapErr(err error, newErr error, msg string) {
	err = errors.Wrap(newErr, msg)
}

//GCLOUD_INCLUSTER
func inCluster() bool {
	ans := os.Getenv("GCLOUD_INCLUSTER")
	switch ans {
	case "False", "f", "F", "false", "":
		return false
	default:
		return true
	}
}
