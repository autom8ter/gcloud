# storage
--
    import "github.com/autom8ter/gcloud/storage"


## Usage

#### type Blob

```go
type Blob struct {
}
```


#### func  NewBlob

```go
func NewBlob(ctx context.Context, opts ...option.ClientOption) (*Blob, error)
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

#### type Document

```go
type Document struct {
}
```


#### func  NewDocument

```go
func NewDocument(ctx context.Context, opts ...option.ClientOption) (*Document, error)
```

#### func (*Document) Client

```go
func (v *Document) Client() *firestore.Client
```

#### func (*Document) Close

```go
func (v *Document) Close()
```

#### type SQL

```go
type SQL struct {
}
```


#### func  NewSQL

```go
func NewSQL(ctx context.Context, opts ...option.ClientOption) (*SQL, error)
```
Must set ""GCLOUD_SPANNER_DB"" in environmental variables

#### func (*SQL) Admin

```go
func (s *SQL) Admin() *database.DatabaseAdminClient
```

#### func (*SQL) Close

```go
func (s *SQL) Close()
```

#### func (*SQL) CreateDatabase

```go
func (s *SQL) CreateDatabase(ctx context.Context, opts ...SQLCreate) (*database.CreateDatabaseOperation, error)
```

#### func (*SQL) InsertStructs

```go
func (s *SQL) InsertStructs(ctx context.Context, tables ...TableStruct) (time.Time, error)
```

#### func (*SQL) InsertTables

```go
func (s *SQL) InsertTables(ctx context.Context, tables ...TableMap) (time.Time, error)
```

#### func (*SQL) Spanner

```go
func (s *SQL) Spanner() *spanner.Client
```

#### type SQLCreate

```go
type SQLCreate func(request *adminpb.CreateDatabaseRequest)
```


#### type SignedUrlFunc

```go
type SignedUrlFunc func(options *storage.SignedURLOptions)
```


#### type Storage

```go
type Storage struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Storage, error)
```

#### func (*Storage) Blob

```go
func (s *Storage) Blob() *Blob
```

#### func (*Storage) Close

```go
func (s *Storage) Close()
```

#### func (*Storage) Document

```go
func (s *Storage) Document() *Document
```

#### func (*Storage) SQL

```go
func (s *Storage) SQL() *SQL
```

#### type TableMap

```go
type TableMap map[string]map[string]interface{}
```


#### type TableStruct

```go
type TableStruct map[string]interface{}
```
