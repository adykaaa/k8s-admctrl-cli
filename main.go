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
	config, err := clientcmd.BuildConfigFromFlags("", "/home/adykaaa/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	webhookConfigurations, err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, webhookConfiguration := range webhookConfigurations.Items {
		fmt.Println(webhookConfiguration.Name)
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
