package gcloud_test

import (
	"context"
	"github.com/autom8ter/gcloud"
	"google.golang.org/api/option"
	"testing"
)

var ctx = context.Background()

func TestNew(t *testing.T) {
	g, err := gcloud.New(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		t.Fatal(err.Error())
	}
	if g == nil {
		t.Fatal("nil gcloud")
	}
}
