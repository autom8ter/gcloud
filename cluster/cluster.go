package cluster

import (
	errs "github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
)

type Cluster struct {
	cli *kubernetes.Clientset
}

//"GCLOUD_CLUSTER" environmental variable to "in" or "In" for in cluster kubernetes client set
// "GCLOUD_CLUSTER_MASTER" to set the kubernetes master url
func New() (*Cluster, error) {
	c := &Cluster{}
	var err error
	var newErr error
	if strings.Contains(os.Getenv("GCLOUD_CLUSTER"), "in") || strings.Contains(os.Getenv("GCLOUD_CLUSTER"), "In") {
		// creates the in-cluster config
		cfg, err := rest.InClusterConfig()
		if err != nil {
			return c, err
		}

		// creates the clientset
		c.cli, newErr = kubernetes.NewForConfig(cfg)
		if newErr != nil {
			err = errs.Wrap(err, newErr.Error())
		}
	} else {
		path := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags(os.Getenv("GCLOUD_CLUSTER_MASTER"), path)
		if err != nil {
			return c, err
		}
		// create the clientset
		c.cli, newErr = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
	}
	return c, err
}

func (c *Cluster) Close() {

}

func (c *Cluster) Client() *kubernetes.Clientset {
	return c.cli
}
