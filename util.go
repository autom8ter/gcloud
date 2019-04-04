package gcloud

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/Masterminds/sprig"
	"github.com/golang/protobuf/proto"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// DefaultClient returns an authenticated http client with the specified scopes
func Client(ctx context.Context, scopes []string) (*http.Client, error) {
	return google.DefaultClient(ctx, scopes...)
}

func JSON(v interface{}) []byte {
	output, _ := json.MarshalIndent(v, "", "  ")
	return output
}

func Proto(msg proto.Message) []byte {
	output, _ := proto.Marshal(msg)
	return output
}

func YAML(v interface{}) []byte {
	output, _ := yaml.Marshal(v)
	return output
}

func MustGetEnv(envKey, defaultValue string) string {
	val := os.Getenv(envKey)
	if val == "" {
		val = defaultValue
	}
	if val == "" {
		log.Fatalf("%q should be set", envKey)
	}
	return val
}

func XML(v interface{}) []byte {
	output, _ := xml.Marshal(v)
	return output
}

func Render(text string, data interface{}, w io.Writer) error {
	t, err := template.New("").Funcs(sprig.GenericFuncMap()).Parse(text)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
