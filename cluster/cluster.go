package cluster

import (
	"fmt"
	errs "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod\n")
		}

		time.Sleep(10 * time.Second)
	}
}
