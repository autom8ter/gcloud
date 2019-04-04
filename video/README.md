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
func (v *Intelligence) Client() *video.Client
```

#### func (*Intelligence) Close

```go
func (v *Intelligence) Close()
```

#### func (*Intelligence) DetectAll

```go
func (i *Intelligence) DetectAll(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) DetectExplicitContent

```go
func (i *Intelligence) DetectExplicitContent(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) DetectFaces

```go
func (i *Intelligence) DetectFaces(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) DetectLabel

```go
func (i *Intelligence) DetectLabel(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) DetectText

```go
func (i *Intelligence) DetectText(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) TrackObjects

```go
func (i *Intelligence) TrackObjects(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```
objectTracking analyzes a video and extracts entities with their bounding boxes.

#### func (*Intelligence) TrackObjectsFromStorage

```go
func (i *Intelligence) TrackObjectsFromStorage(ctx context.Context, gcsURI string, w io.Writer) (*videopb.AnnotateVideoResponse, error)
```

#### func (*Intelligence) TranscribeSpeech

```go
func (i *Intelligence) TranscribeSpeech(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error)
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

#### func (*Vision) Annotator

```go
func (v *Vision) Annotator() *vision.ImageAnnotatorClient
```

#### func (*Vision) Close

```go
func (v *Vision) Close()
```

#### func (*Vision) ProductSearch

```go
func (v *Vision) ProductSearch() *vision.ProductSearchClient
```
