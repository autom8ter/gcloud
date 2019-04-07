package gcloud_test

import (
	"context"
	"github.com/autom8ter/gcloud"
	"github.com/autom8ter/gcloud/handlers"
	"testing"
)

var ctx = context.Background()

func Test(t *testing.T) {
	g, err := gcloud.New(ctx, &gcloud.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := g.Execute(handlers.Noop()); err != nil {
		t.Fatal(err.Error())
	}
}
