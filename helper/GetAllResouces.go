package helper

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ResourceInfo struct {
	Kind      string
	Name      string
	Namespace string
	CreatedAt time.Time
}

func GetAllResources(clientset *kubernetes.Clientset, namespace string) ([]ResourceInfo, error) {
	var resources []ResourceInfo

	// Fetch Deployments
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		resources = append(resources, ResourceInfo{
			Kind:      "Deployment",
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			CreatedAt: deployment.CreationTimestamp.Time,
		})
	}

	// Fetch Services
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, service := range services.Items {
		resources = append(resources, ResourceInfo{
			Kind:      "Service",
			Name:      service.Name,
			Namespace: service.Namespace,
			CreatedAt: service.CreationTimestamp.Time,
		})
	}

	// Add other resource types (StatefulSets, DaemonSets, Jobs, CronJobs, etc.) similarly.

	return resources, nil
}
