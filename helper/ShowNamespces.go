package helper

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceInfo struct {
	Name   string
	Status string
	Age    string
}

func ShowNameSpaces(clientset *kubernetes.Clientset) ([]NamespaceInfo, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaceInfoList []NamespaceInfo
	for _, ns := range namespaces.Items {

		namespaceCreatorTImestamp := ns.GetCreationTimestamp()
		age := time.Since(namespaceCreatorTImestamp.Time).Round(time.Second)

		namespaceInfo := NamespaceInfo{
			Name:   ns.Name,
			Status: string(ns.Status.Phase),
			Age:    age.String(),
		}
		namespaceInfoList = append(namespaceInfoList, namespaceInfo)
	}
	return namespaceInfoList, nil
}
