package k8s

// Create the MutatingWebhookConfiguration object
webhook := &v1beta1.MutatingWebhookConfiguration{
  ObjectMeta: metav1.ObjectMeta{
    Name: "example-webhook",
  },
  Webhooks: []v1beta1.Webhook{
    {
      Name: "example-webhook.example.com",
      Rules: []v1beta1.RuleWithOperations{
        {
          Operations: []v1beta1.OperationType{v1beta1.Create, v1beta1.Update},
          Rule: v1beta1.Rule{
            APIGroups:   []string{""},
            APIVersions: []string{"v1"},
            Resources:   []string{"pods"},
          },
        },
      },
      ClientConfig: v1beta1.WebhookClientConfig{
        Service: &v1beta1.ServiceReference{
          Namespace: "default",
          Name:      "example-webhook-svc",
        },
        // ...
      },
    },
  },
}

// Implement the webhook logic
func mutatePods(pod *v1.Pod) (*v1.Pod, error) {
  // mutate the pod here
  return pod, nil
}

// Register the webhook with the cluster
kubeClient, err := kubernetes.NewForConfig(config)
if err != nil {
  // handle error
}

err = kubeClient.AdmissionregistrationV1beta1().Mutating

