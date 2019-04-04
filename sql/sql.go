package sql

import (
	"context"
	"google.golang.org/api/option"
)

type SQL struct {
}

func New(ctx context.Context, opts ...option.ClientOption) (*SQL, error) {
	return &SQL{}, nil
}

func (v *SQL) Close() {

}

func (v *SQL) Client() {

}
