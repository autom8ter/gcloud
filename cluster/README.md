# cluster
--
    import "github.com/autom8ter/gcloud/cluster"


## Usage

#### type Cluster

```go
type Cluster struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Cluster, error)
```

#### func (*Cluster) Close

```go
func (c *Cluster) Close()
```
