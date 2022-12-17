package k8s

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Load the kubeconfig file
	kubeconfig := "/path/to/kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	// Create the Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Read the deployment YAML file
	deploymentYAML, err := ioutil.ReadFile("/path/to/deployment.yaml")
	if err != nil {
		panic(err)
	}

	// Create the deployment
	deployment := &appsv1.Deployment{}
	err = yaml.Unmarshal(deploymentYAML, deployment)
	if err != nil {
		panic(err)
	}

	// Deploy the deployment
	result, err := clientset.AppsV1().Deployments("default").Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployment created: %q\n", result.GetObjectMeta().GetName())
}
