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
	"net/http"
)

// DefaultClient returns an authenticated http client with the specified scopes
func DefaultClient(ctx context.Context, scopes []string) (*http.Client, error) {
	return google.DefaultClient(ctx, scopes...)
}

func toJSON(v interface{}) []byte {
	output, _ := json.MarshalIndent(v, "", "  ")
	return output
}

func toProto(msg proto.Message) []byte {
	output, _ := proto.Marshal(msg)
	return output
}

func toYAML(v interface{}) []byte {
	output, _ := yaml.Marshal(v)
	return output
}

func toXML(v interface{}) []byte {
	output, _ := xml.Marshal(v)
	return output
}

func render(text string, data interface{}, w io.Writer) error {
	t, err := template.New("").Funcs(sprig.GenericFuncMap()).Parse(text)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
