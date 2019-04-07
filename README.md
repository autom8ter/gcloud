# gcloud
--
    import "github.com/autom8ter/gcloud"


## Usage

#### type Clients

```go
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
```

GCP ClientSet

#### type Config

```go
type Config struct {
	Project   string
	Scopes    []string
	InCluster bool
	SpannerDB string
	Options   []option.ClientOption
}
```

Config is used to create a new GCP instance

#### type GCP

```go
type GCP struct {
}
```

GCP holds Google Cloud Platform Clients and Services

#### func  New

```go
func New(ctx context.Context, cfg *Config) (*GCP, error)
```
New returns a new authenticated GCP instance from the provided context and
config

#### func (*GCP) Clients

```go
func (g *GCP) Clients() *Clients
```
Clients returns an authenticated GCP ClientSet

#### func (*GCP) Close

```go
func (g *GCP) Close()
```
Close closes all clients

#### func (*GCP) Configuration

```go
func (g *GCP) Configuration() *Config
```
Configuration returns the config used to create the GCP instance

#### func (*GCP) Context

```go
func (g *GCP) Context() context.Context
```
Context returns the context used to create the GCP instance

#### func (*GCP) HTTP

```go
func (g *GCP) HTTP() *http.Client
```
HTTP returns a google default HTTP client

#### func (*GCP) Services

```go
func (g *GCP) Services() *Services
```
Services returns an authenticated GCP ServiceSet

#### func (*GCP) Trace

```go
func (g *GCP) Trace() *stackdriver.Exporter
```
Trace returns a stackdriver exporter

#### type Services

```go
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
}
```

GCP ServiceSet. Make sure to pass the necessary scopes in your config to
successfully initialize services.
