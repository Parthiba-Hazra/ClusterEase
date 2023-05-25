package helper

import (
	"context"
	"fmt"
	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func CreateResourcesFromYAML(clientset *kubernetes.Clientset, namespace string, filePath string) error {
	// Read the YAML file contents
	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create the decoder to decode YAML content into unstructured objects
	decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDeserializer()

	// Decode the YAML content into an unstructured object
	obj, _, err := decoder.Decode(yamlContent, nil, nil)
	if err != nil {
		return err
	}

	// Convert the unstructured object to a typed object
	typedObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return err
	}

	// Create the appropriate resource based on the kind
	switch kind := typedObj["kind"]; kind {
	case "Deployment":
		// Convert the typed object to a Deployment object
		deployment := &appsv1.Deployment{}
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(typedObj, deployment)
		if err != nil {
			return err
		}

		// Create the deployment in the cluster
		_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}

		fmt.Println("Deployment created:", deployment.GetName())

	case "Service":
		// Convert the typed object to a Service object
		service := &corev1.Service{}
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(typedObj, service)
		if err != nil {
			return err
		}

		// Create the service in the cluster
		_, err = clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
		if err != nil {
			return err
		}

		fmt.Println("Service created:", service.GetName())

	// Add more cases for other resource kinds as needed

	default:
		return fmt.Errorf("unsupported kind: %s", kind)
	}

	return nil
}
