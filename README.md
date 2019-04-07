# gcloud
--
    import "github.com/autom8ter/gcloud"


## Usage

#### type GCP

```go
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
	Language           *language.Client
}
```

GCP holds Google Cloud Platform Clients and carries some utility functions
optional environmental variables: "GCLOUD_PROJECTID", "GCLOUD_SPANNER_DB"
"GCLOUD_CLUSTER_MASTER" "GCLOUD_CLUSTER", "GCLOUD_INCLUSTER"

#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error)
```
New returns a new authenticated GCP instance from the provided api options
GCLOUD_PROJECTID, GCLOUD_SPANNERDB,

#### func (*GCP) Close

```go
func (g *GCP) Close()
```
Close closes all clients

#### func (*GCP) DefaultClient

```go
func (g *GCP) DefaultClient(ctx context.Context, scopes []string) (*http.Client, error)
```
DefaultClient returns an authenticated http client with the specified scopes
