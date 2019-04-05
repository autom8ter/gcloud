package gcloud_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/gcloud"
	"google.golang.org/api/option"
	"testing"
	"time"
)

var ctx = context.Background()
var toTranslate = []string{"Hello World!"}
var translated = map[string][]string{}

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
	{
		chi, err := g.ToChinese(ctx, toTranslate)
		if err != nil {
			t.Fatal(err.Error())
		}
		translated["chinese"] = chi
	}
	{
		spa, err := g.ToSpanish(ctx, toTranslate)
		if err != nil {
			t.Fatal(err.Error())
		}
		translated["spanish"] = spa
	}
	{
		fr, err := g.ToFrench(ctx, toTranslate)
		if err != nil {
			t.Fatal(err.Error())
		}
		translated["french"] = fr
	}
	fmt.Println(string(g.JSON(translated)))
}
