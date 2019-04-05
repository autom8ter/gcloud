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

#### func (*GCP) Render

```go
func (g *GCP) Render(text string, data interface{}, w io.Writer) error
```
Render uses html/template along with the sprig funcmap functions to render a
strings to an io writer ref: https://github.com/Masterminds/sprig

#### func (*GCP) ToChinese

```go
func (g *GCP) ToChinese(ctx context.Context, content []string) ([]string, error)
```
ToSpanish returns the provided content in Chinese

#### func (*GCP) ToFrench

```go
func (g *GCP) ToFrench(ctx context.Context, content []string) ([]string, error)
```
ToFrench returns the provided content in French

#### func (*GCP) ToGerman

```go
func (g *GCP) ToGerman(ctx context.Context, content []string) ([]string, error)
```
ToGerman returns the provided content in German

#### func (*GCP) ToItalian

```go
func (g *GCP) ToItalian(ctx context.Context, content []string) ([]string, error)
```
ToItalian returns the provided content in Italian

#### func (*GCP) ToRussian

```go
func (g *GCP) ToRussian(ctx context.Context, content []string) ([]string, error)
```
ToRussian returns the provided content in Russian

#### func (*GCP) ToSpanish

```go
func (g *GCP) ToSpanish(ctx context.Context, content []string) ([]string, error)
```
ToSpanish returns the provided content in Spanish

#### func (*GCP) WriteAudioTranscript

```go
func (g *GCP) WriteAudioTranscript(ctx context.Context, audioURL string, w io.Writer) error
```

#### func (*GCP) WriteTextToSpeechMP3

```go
func (g *GCP) WriteTextToSpeechMP3(ctx context.Context, text string, w io.Writer) error
```

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
