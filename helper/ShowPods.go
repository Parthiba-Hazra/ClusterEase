package helper

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodInfo struct {
	Name      string
	Namespace string
}

func ShowPod(clientset *kubernetes.Clientset) ([]PodInfo, error) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podInfoList []PodInfo
	for _, pod := range pods.Items {
		podInfo := PodInfo{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		}

		podInfoList = append(podInfoList, podInfo)
	}

	return podInfoList, nil
}
