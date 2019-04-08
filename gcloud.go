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
	"golang.org/x/oauth2/google"
	"google.golang.org/api/blogger/v3"
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

// GCP is the configuration used to return gcp clients and services. Use Init() to validate GCP before using it.
type GCP struct {
	Project string                `validate:"required"`
	Scopes  []string              `validate:"required"`
	Options []option.ClientOption `validate:"required"`
}

func NewGCP(project string, scopes []string, options ...option.ClientOption) *GCP {
	return &GCP{Project: project, Scopes: scopes, Options: options}
}

func (g *GCP) Init() error {
	return tool.Validate(g)
}

func (g *GCP) HTTP(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, g.Scopes...)
}

func (g *GCP) PubSub(ctx context.Context) (*pubsub.Client, error) {
	return pubsub.NewClient(ctx, g.Project, g.Options...)
}

func (g *GCP) Firestore(ctx context.Context) (*firestore.Client, error) {
	return firestore.NewClient(ctx, g.Project, g.Options...)
}

func (g *GCP) Translate(ctx context.Context) (*translate.Client, error) {
	return translate.NewClient(ctx, g.Options...)
}

func (g *GCP) IAM(ctx context.Context) (*iam.Service, error) {
	return iam.NewService(ctx, g.Options...)
}

func (g *GCP) Storage(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx, g.Options...)
}

func (g *GCP) IOT(ctx context.Context) (*iot.DeviceManagerClient, error) {
	return iot.NewDeviceManagerClient(ctx, g.Options...)
}

func (g *GCP) Kube(inCluster bool) (*kubernetes.Clientset, error) {
	return clients.NewKubernetesClientSet(inCluster)
}

func (g *GCP) Language(ctx context.Context) (*language.Client, error) {
	return language.NewClient(ctx, g.Options...)
}

func (g *GCP) Spanner(ctx context.Context, database string) (*spanner.Client, error) {
	return spanner.NewClient(ctx, database, g.Options...)
}

func (g *GCP) DBAdmin(ctx context.Context) (*database.DatabaseAdminClient, error) {
	return database.NewDatabaseAdminClient(ctx, g.Options...)
}

func (g *GCP) KMS(ctx context.Context) (*kms.KeyManagementClient, error) {
	return kms.NewKeyManagementClient(ctx, g.Options...)
}

func (g *GCP) VideoIntelligence(ctx context.Context) (*videointelligence.Client, error) {
	return videointelligence.NewClient(ctx, g.Options...)
}

func (g *GCP) ImageAnnotator(ctx context.Context) (*vision.ImageAnnotatorClient, error) {
	return vision.NewImageAnnotatorClient(ctx, g.Options...)
}

func (g *GCP) ImageProductSearch(ctx context.Context) (*vision.ProductSearchClient, error) {
	return vision.NewProductSearchClient(ctx, g.Options...)
}

func (g *GCP) Text2Speech(ctx context.Context) (*texttospeech.Client, error) {
	return texttospeech.NewClient(ctx, g.Options...)
}

func (g *GCP) Speech(ctx context.Context) (*speech.Client, error) {
	return speech.NewClient(ctx, g.Options...)
}

func (g *GCP) Container(ctx context.Context) (*container.Service, error) {
	return container.NewService(ctx, g.Options...)
}

func (g *GCP) HealthCare(ctx context.Context) (*healthcare.Service, error) {
	return healthcare.NewService(ctx, g.Options...)
}

func (g *GCP) Calendar(ctx context.Context) (*healthcare.Service, error) {
	return healthcare.NewService(ctx, g.Options...)
}

func (g *GCP) Blogger(ctx context.Context) (*blogger.Service, error) {
	return blogger.NewService(ctx, g.Options...)
}

func (g *GCP) CustomSearch(ctx context.Context) (*customsearch.Service, error) {
	return customsearch.NewService(ctx, g.Options...)
}

func (g *GCP) ClassRoom(ctx context.Context) (*class.Service, error) {
	return class.NewService(ctx, g.Options...)
}

func (g *GCP) Content(ctx context.Context) (*content.APIService, error) {
	return content.NewService(ctx, g.Options...)
}

func (g *GCP) OSLogin(ctx context.Context) (*oslogin.Service, error) {
	return oslogin.NewService(ctx, g.Options...)
}

func (g *GCP) People(ctx context.Context) (*people.Service, error) {
	return people.NewService(ctx, g.Options...)
}

func (g *GCP) Photos(cli *http.Client) (*photos.Service, error) {
	return photos.New(cli)
}

func (g *GCP) Prediction(cli *http.Client) (*prediction.Service, error) {
	return prediction.New(cli)
}

func (g *GCP) Redis(ctx context.Context) (*redis.Service, error) {
	return redis.NewService(ctx, g.Options...)
}

func (g *GCP) RuntimeGCP(ctx context.Context) (*run.Service, error) {
	return run.NewService(ctx, g.Options...)
}

func (g *GCP) Sheets(ctx context.Context) (*sheets.Service, error) {
	return sheets.NewService(ctx, g.Options...)
}

func (g *GCP) Slides(ctx context.Context) (*slides.Service, error) {
	return slides.NewService(ctx, g.Options...)
}

func (g *GCP) Tasks(ctx context.Context) (*tasks.Service, error) {
	return tasks.NewService(ctx, g.Options...)
}

func (g *GCP) YoutTube(ctx context.Context) (*youtube.Service, error) {
	return youtube.NewService(ctx, g.Options...)
}

func (g *GCP) Docs(ctx context.Context) (*docs.Service, error) {
	return docs.NewService(ctx, g.Options...)
}

func (g *GCP) Jobs(ctx context.Context) (*jobs.Service, error) {
	return jobs.NewService(ctx, g.Options...)
}

func (g *GCP) Domains(ctx context.Context) (*plusdomains.Service, error) {
	return plusdomains.NewService(ctx, g.Options...)
}
