# gcloud
--
    import "github.com/autom8ter/gcloud"


## Usage

#### type GCP

```go
type GCP struct {
	Options []option.ClientOption `validate:"required"`
}
```

GCP is the configuration used to return gcp clients and services. Use Init() to
validate GCP before using it.

#### func  NewGCP

```go
func NewGCP(options ...option.ClientOption) *GCP
```

#### func (*GCP) Blogger

```go
func (g *GCP) Blogger(ctx context.Context) (*blogger.Service, error)
```

#### func (*GCP) Calendar

```go
func (g *GCP) Calendar(ctx context.Context) (*healthcare.Service, error)
```

#### func (*GCP) ClassRoom

```go
func (g *GCP) ClassRoom(ctx context.Context) (*class.Service, error)
```

#### func (*GCP) Container

```go
func (g *GCP) Container(ctx context.Context) (*container.Service, error)
```

#### func (*GCP) Content

```go
func (g *GCP) Content(ctx context.Context) (*content.APIService, error)
```

#### func (*GCP) CustomSearch

```go
func (g *GCP) CustomSearch(ctx context.Context) (*customsearch.Service, error)
```

#### func (*GCP) DBAdmin

```go
func (g *GCP) DBAdmin(ctx context.Context) (*database.DatabaseAdminClient, error)
```

#### func (*GCP) Docs

```go
func (g *GCP) Docs(ctx context.Context) (*docs.Service, error)
```

#### func (*GCP) Domains

```go
func (g *GCP) Domains(ctx context.Context) (*plusdomains.Service, error)
```

#### func (*GCP) Firestore

```go
func (g *GCP) Firestore(ctx context.Context, project string) (*firestore.Client, error)
```

#### func (*GCP) HTTP

```go
func (g *GCP) HTTP(ctx context.Context, scopes []string) (*http.Client, error)
```

#### func (*GCP) HealthCare

```go
func (g *GCP) HealthCare(ctx context.Context) (*healthcare.Service, error)
```

#### func (*GCP) IAM

```go
func (g *GCP) IAM(ctx context.Context) (*iam.Service, error)
```

#### func (*GCP) IOT

```go
func (g *GCP) IOT(ctx context.Context) (*iot.DeviceManagerClient, error)
```

#### func (*GCP) ImageAnnotator

```go
func (g *GCP) ImageAnnotator(ctx context.Context) (*vision.ImageAnnotatorClient, error)
```

#### func (*GCP) ImageProductSearch

```go
func (g *GCP) ImageProductSearch(ctx context.Context) (*vision.ProductSearchClient, error)
```

#### func (*GCP) Init

```go
func (g *GCP) Init() error
```

#### func (*GCP) Jobs

```go
func (g *GCP) Jobs(ctx context.Context) (*jobs.Service, error)
```

#### func (*GCP) KMS

```go
func (g *GCP) KMS(ctx context.Context) (*kms.KeyManagementClient, error)
```

#### func (*GCP) Kube

```go
func (g *GCP) Kube(inCluster bool) (*kubernetes.Clientset, error)
```

#### func (*GCP) Language

```go
func (g *GCP) Language(ctx context.Context) (*language.Client, error)
```

#### func (*GCP) OSLogin

```go
func (g *GCP) OSLogin(ctx context.Context) (*oslogin.Service, error)
```

#### func (*GCP) People

```go
func (g *GCP) People(ctx context.Context) (*people.Service, error)
```

#### func (*GCP) Photos

```go
func (g *GCP) Photos(cli *http.Client) (*photos.Service, error)
```

#### func (*GCP) Prediction

```go
func (g *GCP) Prediction(cli *http.Client) (*prediction.Service, error)
```

#### func (*GCP) PubSub

```go
func (g *GCP) PubSub(ctx context.Context, project string) (*pubsub.Client, error)
```

#### func (*GCP) Redis

```go
func (g *GCP) Redis(ctx context.Context) (*redis.Service, error)
```

#### func (*GCP) RuntimeGCP

```go
func (g *GCP) RuntimeGCP(ctx context.Context) (*run.Service, error)
```

#### func (*GCP) Sheets

```go
func (g *GCP) Sheets(ctx context.Context) (*sheets.Service, error)
```

#### func (*GCP) Slides

```go
func (g *GCP) Slides(ctx context.Context) (*slides.Service, error)
```

#### func (*GCP) Spanner

```go
func (g *GCP) Spanner(ctx context.Context, database string) (*spanner.Client, error)
```

#### func (*GCP) Speech

```go
func (g *GCP) Speech(ctx context.Context) (*speech.Client, error)
```

#### func (*GCP) Storage

```go
func (g *GCP) Storage(ctx context.Context) (*storage.Client, error)
```

#### func (*GCP) Tasks

```go
func (g *GCP) Tasks(ctx context.Context) (*tasks.Service, error)
```

#### func (*GCP) Text2Speech

```go
func (g *GCP) Text2Speech(ctx context.Context) (*texttospeech.Client, error)
```

#### func (*GCP) Translate

```go
func (g *GCP) Translate(ctx context.Context) (*translate.Client, error)
```

#### func (*GCP) VideoIntelligence

```go
func (g *GCP) VideoIntelligence(ctx context.Context) (*videointelligence.Client, error)
```

#### func (*GCP) YoutTube

```go
func (g *GCP) YoutTube(ctx context.Context) (*youtube.Service, error)
```
