package k8s

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClientSet() *kubernetes.Clientset {
	// read the kubeconfig environment variable
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// create a kubernetes client
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatalf("Could not create a K8s client configuration: ", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Could not initiate the HTTP client for K8s: ", err)
	}

	return clientset
}
