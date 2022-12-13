package main

import (
	"context"
	"fmt"
	"log"

	k8s "github.com/adykaaa/k8s-admctrl-cli/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientset, err := k8s.NewDefaultClientSet()
	if err != nil {
		log.Fatalf("error during clientset initialization : %v", err)
	}

	webhookConfigurations, err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("There are no validating webhooks in the cluster!")
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
