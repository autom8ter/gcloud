package gcloud_test

import (
	"context"
	"github.com/autom8ter/gcloud"
	"testing"
)

var ctx = context.Background()

func Test(t *testing.T) {
	_, err := gcloud.New(ctx, &gcloud.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
}
