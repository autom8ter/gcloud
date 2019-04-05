package cluster

import (
	"context"
	"google.golang.org/api/option"
)

type Cluster struct {

}


func New(ctx context.Context, opts ...option.ClientOption) (*Cluster, error) {

	return &Cluster{

	}, nil
}

func (c *Cluster) Close() {

}