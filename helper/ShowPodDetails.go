package helper

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodDetails struct {
	Name             string
	Namespace        string
	CreationTime     time.Time
	Phase            string
	Conditions       []PodCondition
	IP               string
	ContainerDetails []ContainerDetails
}

type PodCondition struct {
	Type               string
	Status             string
	LastTransitionTime time.Time
	Reason             string
	Message            string
}

type ContainerDetails struct {
	ContainerName string
	Ports         []PortDetails
}

type PortDetails struct {
	PortName      string
	Protocol      string
	ContainerPort int32
	HostPort      int32
}

func GetPodDetails(clientset *kubernetes.Clientset, namespace string, podName string) ([]PodDetails, error) {
	poddetail, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var podDetailsList []PodDetails

	pod := poddetail
	podDetails := PodDetails{
		Name:             pod.Name,
		Namespace:        pod.Namespace,
		CreationTime:     pod.CreationTimestamp.Time,
		Phase:            string(pod.Status.Phase),
		Conditions:       make([]PodCondition, len(pod.Status.Conditions)),
		IP:               pod.Status.PodIP,
		ContainerDetails: make([]ContainerDetails, len(pod.Spec.Containers)),
	}

	for i, condition := range pod.Status.Conditions {
		podDetails.Conditions[i] = PodCondition{
			Type:               string(condition.Type),
			Status:             string(condition.Status),
			LastTransitionTime: condition.LastTransitionTime.Time,
			Reason:             condition.Reason,
			Message:            condition.Message,
		}
	}

	for i, container := range pod.Spec.Containers {
		containerDetails := ContainerDetails{
			ContainerName: container.Name,
			Ports:         make([]PortDetails, len(container.Ports)),
		}
		for j, port := range container.Ports {
			containerDetails.Ports[j] = PortDetails{
				PortName:      port.Name,
				Protocol:      string(port.Protocol),
				ContainerPort: port.ContainerPort,
				HostPort:      port.HostPort,
			}
		}
		podDetails.ContainerDetails[i] = containerDetails
	}

	podDetailsList = append(podDetailsList, podDetails)

	return podDetailsList, nil
}
