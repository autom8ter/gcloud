# blob
--
    import "github.com/autom8ter/gcloud/blob"


## Usage

#### type Blob

```go
type Blob struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Blob, error)
```

#### func (*Blob) AddBucketACL

```go
func (b *Blob) AddBucketACL(ctx context.Context, bucket string, e storage.ACLEntity, role storage.ACLRole) error
```

#### func (*Blob) BucketAttributes

```go
func (b *Blob) BucketAttributes(ctx context.Context, name string, attributes storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
```

#### func (*Blob) BucketRules

```go
func (b *Blob) BucketRules(ctx context.Context, bucket string) ([]storage.ACLRule, error)
```

#### func (*Blob) Client

```go
func (b *Blob) Client() *storage.Client
```

#### func (*Blob) Close

```go
func (b *Blob) Close()
```

#### func (*Blob) CreateBucket

```go
func (b *Blob) CreateBucket(ctx context.Context, name, project string, attributes *storage.BucketAttrs) error
```

#### func (*Blob) DeleteBucket

```go
func (b *Blob) DeleteBucket(ctx context.Context, name string) error
```

#### func (*Blob) ObjectURL

```go
func (b *Blob) ObjectURL(objAttrs *storage.ObjectAttrs) string
```

#### func (*Blob) SignedURL

```go
func (b *Blob) SignedURL(bucket, object string, opts ...SignedUrlFunc) (string, error)
```

#### func (*Blob) UpdateBucket

```go
func (b *Blob) UpdateBucket(ctx context.Context, name string, attributes storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
```

#### func (*Blob) UploadObject

```go
func (b *Blob) UploadObject(ctx context.Context, r io.Reader, projectID, bucket, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error)
```

#### type SignedUrlFunc

```go
type SignedUrlFunc func(options *storage.SignedURLOptions)
```
