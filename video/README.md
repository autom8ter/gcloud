# video
--
    import "github.com/autom8ter/gcloud/video"


## Usage

#### type Intelligence

```go
type Intelligence struct {
}
```


#### func  NewIntelligence

```go
func NewIntelligence(ctx context.Context, opts ...option.ClientOption) (*Intelligence, error)
```

#### func (*Intelligence) Client

```go
func (v *Intelligence) Client()
```

#### func (*Intelligence) Close

```go
func (v *Intelligence) Close()
```

#### type Video

```go
type Video struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Video, error)
```

#### func (*Video) Close

```go
func (v *Video) Close()
```

#### func (*Video) Intelligence

```go
func (v *Video) Intelligence() *Intelligence
```

#### func (*Video) Vision

```go
func (v *Video) Vision() *Vision
```

#### type Vision

```go
type Vision struct {
}
```


#### func  NewVision

```go
func NewVision(ctx context.Context, opts ...option.ClientOption) (*Vision, error)
```

#### func (*Vision) Client

```go
func (v *Vision) Client()
```

#### func (*Vision) Close

```go
func (v *Vision) Close()
```
