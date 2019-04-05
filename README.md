# gcloud
--
    import "github.com/autom8ter/gcloud"


## Usage

#### func  Client

```go
func Client(ctx context.Context, scopes []string) (*http.Client, error)
```
DefaultClient returns an authenticated http client with the specified scopes

#### func  JSON

```go
func JSON(v interface{}) []byte
```

#### func  MustGetEnv

```go
func MustGetEnv(envKey, defaultValue string) string
```

#### func  Proto

```go
func Proto(msg proto.Message) []byte
```

#### func  Render

```go
func Render(text string, data interface{}, w io.Writer) error
```

#### func  XML

```go
func XML(v interface{}) []byte
```

#### func  YAML

```go
func YAML(v interface{}) []byte
```

#### type GCP

```go
type GCP struct {
}
```

GCP holds Google Cloud Platform Clients and carries some utility functions
optional environmental variables: "GCLOUD_PROJECTID", "GCLOUD_SPANNER_DB"
"GCLOUD_CLUSTER_MASTER" "GCLOUD_CLUSTER"

#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*GCP, error)
```
New returns a new authenticated GCP instance from the provided api options

#### func (*GCP) Auth

```go
func (g *GCP) Auth() *auth.Auth
```
Auth returns a client used for GCP key management and IAM

#### func (*GCP) Close

```go
func (g *GCP) Close()
```
Close closes all clients

#### func (*GCP) Cluster

```go
func (g *GCP) Cluster() *cluster.Cluster
```
Cluster returns a registered kubernetes clientset "GCLOUD_CLUSTER_MASTER"
"GCLOUD_CLUSTER"

#### func (*GCP) DefaultClient

```go
func (g *GCP) DefaultClient(ctx context.Context, scopes []string) (*http.Client, error)
```
DefaultClient returns an authenticated http client with the specified scopes

#### func (*GCP) Execute

```go
func (g *GCP) Execute(ctx context.Context, fns ...HandlerFunc) error
```
Execute runs all functions and returns a wrapped error

#### func (*GCP) JSON

```go
func (g *GCP) JSON(obj interface{}) []byte
```
JSON formats an object and turns it into JSON bytes

#### func (*GCP) Proto

```go
func (g *GCP) Proto(m proto.Message) []byte
```
Proto formats an object and turns it into Proto bytes

#### func (*GCP) PubSub

```go
func (g *GCP) PubSub() *pubsub.PubSub
```
PubSub returns a client used for GCP pubsub

#### func (*GCP) Render

```go
func (g *GCP) Render(text string, data interface{}, w io.Writer) error
```
Render uses html/template along with the sprig funcmap functions to render a
strings to an io writer ref: https://github.com/Masterminds/sprig

#### func (*GCP) Robots

```go
func (g *GCP) Robots() *robots.Robot
```
Auth returns a client used for GCP key management and IAM

#### func (*GCP) Storage

```go
func (g *GCP) Storage() *storage.Storage
```
Storage returns a client used for GCP blob storage, firestore (documents), and
cloud sql spanner

#### func (*GCP) Text

```go
func (g *GCP) Text() *text.Text
```
Text returns a client used for common text operations: GCP text2speech,
translation, and speech services

#### func (*GCP) Trace

```go
func (g *GCP) Trace() *trace.Trace
```
Trace returns a registered stackdriver exporter

#### func (*GCP) Video

```go
func (g *GCP) Video() *video.Video
```
Video returns a client used for torrenting(non-gcp), GCP video intelligence and
GCP computer vision

#### func (*GCP) XML

```go
func (g *GCP) XML(obj interface{}) []byte
```
XML formats an object and turns it into XML bytes

#### func (*GCP) YAML

```go
func (g *GCP) YAML(obj interface{}) []byte
```
YAML formats an object and turns it into YAML bytes

#### type HandlerFunc

```go
type HandlerFunc func(g *GCP) error
```

HandlerFunc is used to run a function using a GCP object (see GCP.Execute)
Creating a HandlerFunc is easy...

    func NewHandlerFunc() HandlerFunc {
    	return func(g *GCP) error {

    	this is similar to http.HandlerFunc...

    	return nil
    }}
