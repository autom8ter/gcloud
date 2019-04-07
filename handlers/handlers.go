package handlers

import (
	"github.com/autom8ter/gcloud"
	"github.com/pkg/errors"
	"log"
)

func Noop() gcloud.HandlerFunc {
	return func(gcp *gcloud.GCP) error {
		if gcp == nil {
			log.Println("noop handler registered nil gcp instance")
			return errors.New("nil gcp instance detected")
		}
		log.Println("noop handler")
		return nil
	}
}
