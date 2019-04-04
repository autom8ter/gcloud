# trace
--
    import "github.com/autom8ter/gcloud/trace"


## Usage

#### type Trace

```go
type Trace struct {
}
```


#### func  New

```go
func New(opts ...option.ClientOption) (*Trace, error)
```

#### func (*Trace) Exporter

```go
func (t *Trace) Exporter() *stackdriver.Exporter
```

#### func (*Trace) Flush

```go
func (t *Trace) Flush()
```

#### func (*Trace) Register

```go
func (t *Trace) Register(views ...*view.View) error
```
