package helper

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodInfo struct {
	Name      string
	Namespace string
	Status    string
	Ready     string
	Restart   string
}

func ShowPod(clientset *kubernetes.Clientset, namespace string) ([]PodInfo, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podInfoList []PodInfo
	for _, pod := range pods.Items {
		var containerRestarts int32
		var containerReady int
		var totalContainers int

		for container := range pod.Spec.Containers {
			containerRestarts += pod.Status.ContainerStatuses[container].RestartCount
			if pod.Status.ContainerStatuses[container].Ready {
				containerReady++
			}
			totalContainers++
		}

		ready := fmt.Sprintf("%v/%v", containerReady, totalContainers)
		restarts := fmt.Sprintf("%v", containerRestarts)
		reason := getPodReason(pod)
		if reason == "" {
			status := pod.Status.Phase
			podInfo := PodInfo{
				Name:      pod.Name,
				Namespace: pod.Namespace,
				Status:    string(status),
				Ready:     ready,
				Restart:   restarts,
			}
			podInfoList = append(podInfoList, podInfo)
		} else {
			podInfo := PodInfo{
				Name:      pod.Name,
				Namespace: pod.Namespace,
				Status:    string(reason),
				Ready:     ready,
				Restart:   restarts,
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
