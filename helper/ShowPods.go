package helper

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodInfo struct {
	Name      string
	Namespace string
	Status    string
}

func ShowPod(clientset *kubernetes.Clientset, namespace string) ([]PodInfo, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podInfoList []PodInfo
	for _, pod := range pods.Items {
		reason := getPodReason(pod)
		if reason == "" {
			status := pod.Status.Phase
			podInfo := PodInfo{
				Name:      pod.Name,
				Namespace: pod.Namespace,
				Status:    string(status),
			}
			podInfoList = append(podInfoList, podInfo)
		} else {
			podInfo := PodInfo{
				Name:      pod.Name,
				Namespace: pod.Namespace,
				Status:    string(reason),
			}
			podInfoList = append(podInfoList, podInfo)
		}
	}

	return podInfoList, nil
}

func getPodReason(pod v1.Pod) string {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		state := containerStatus.State
		if state.Waiting != nil {
			return state.Waiting.Reason
		} else if state.Terminated != nil {
			return state.Terminated.Reason
		}
	}
	return ""
}
