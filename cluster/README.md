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
func New() (*Cluster, error)
```
"GCLOUD_CLUSTER" environmental variable to "in" or "In" for in cluster
kubernetes client set "GCLOUD_CLUSTER_MASTER" to set the kubernetes master url

#### func (*Cluster) Client

```go
func (c *Cluster) Client() *kubernetes.Clientset
```

#### func (*Cluster) Close

```go
func (c *Cluster) Close()
```
