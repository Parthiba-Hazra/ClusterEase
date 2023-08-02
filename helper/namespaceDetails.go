package helper

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceDetails struct {
	Name          string
	CreationTime  time.Time
	Status        string
	Labels        map[string]string
	Annotations   map[string]string
	ResourceQuota string
}

func GetNamespaceDetails(clientset *kubernetes.Clientset, namespace string) ([]NamespaceDetails, error) {
	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var nsDetailsList []NamespaceDetails

	nsDetails := NamespaceDetails{
		Name:          ns.Name,
		CreationTime:  ns.CreationTimestamp.Time,
		Status:        string(ns.Status.Phase),
		Labels:        ns.Labels,
		Annotations:   ns.Annotations,
		ResourceQuota: getResourceQuota(clientset, namespace),
	}

	nsDetailsList = append(nsDetailsList, nsDetails)

	return nsDetailsList, nil
}

func getResourceQuota(clientset *kubernetes.Clientset, namespace string) string {
	quota, err := clientset.CoreV1().ResourceQuotas(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return ""
	}

	if len(quota.Items) > 0 {
		return quota.Items[0].Name
	}

	return ""
}
