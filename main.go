package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Create a new Kubernetes config object
	config, err := clientcmd.BuildConfigFromFlags("", "/home/adykaaa/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Use the clientset to do something, such as list all of the pods in the current namespace
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d pods in the current namespace\n", len(pods.Items))
}
