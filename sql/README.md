# sql
--
    import "github.com/autom8ter/gcloud/sql"


## Usage

#### type SQL

```go
type SQL struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*SQL, error)
```

#### func (*SQL) Client

```go
func (v *SQL) Client()
```

#### func (*SQL) Close

```go
func (v *SQL) Close()
```
