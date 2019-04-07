package clients

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func NewKubernetesClientSet(inCluster bool) (*kubernetes.Clientset, error) {
	if inCluster {
		// creates the in-cluster config
		cfg, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		// creates the clientset
		cli, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			return nil, err
		}
		return cli, nil
	} else {
		path := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags(os.Getenv("GCLOUD_CLUSTER_MASTER"), path)
		if err != nil {
			return nil, err
		}
		// create the clientset
		cli, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		return cli, nil
	}
}
