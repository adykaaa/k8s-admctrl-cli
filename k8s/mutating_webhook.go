package k8s

import (
	"context"

	"k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/rest"
)

func deployMutatingWebhookConfiguration(namespace string) error {

	// create the MutatingWebhookConfiguration object
	mwc := &v1beta1.MutatingWebhookConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-mutating-webhook-configuration",
		},
		Webhooks: []v1beta1.MutatingWebhook{
			// add your webhooks here
		},
	}

	// create the MutatingWebhookConfiguration
	_, err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Create(context.TODO(), mwc, metav1.CreateOptions{})
	return err
}
