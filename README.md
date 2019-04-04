# gcloud
--
    import "github.com/autom8ter/gcloud"


## Usage

#### type Func

```go
type Func func(g *GCP) error
```

Func is used to run a function using a GCP object (see GCP.Execute)

#### type GCP

```go
type GCP struct {
}
```

GCP holds Google Cloud Platform Clients and carries some utility functions

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

#### func (*GCP) Execute

```go
func (g *GCP) Execute(fns ...Func) error
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

#### func (*GCP) Text

```go
func (g *GCP) Text() *text.Text
```
Text returns a client used for common text operations: GCP text2speech,
translation, and speech services

#### func (*GCP) Video

```go
func (g *GCP) Video() *video.Video
```
PubSub returns a client used for GCP video intelligence and computer vision

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
