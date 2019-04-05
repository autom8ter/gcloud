package gcloud_test

import (
	"context"
	"github.com/autom8ter/gcloud"
	"google.golang.org/api/option"
	"testing"
	"time"
)

var ctx = context.Background()

type User struct {
	ID          string
	DisplayName string
	Email       string
	Phone       string
	Password    string
	Token       string
	Annotations map[string]interface{}
}

func Test(t *testing.T) {
	ctx = context.WithValue(ctx, "start", time.Now())
	g, err := gcloud.New(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		t.Fatal(err.Error())
	}
	if g == nil {
		t.Fatal("nil gcloud")
	}
}
