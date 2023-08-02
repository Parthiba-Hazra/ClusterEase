package helper

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentDetails struct {
	Name              string
	Namespace         string
	CreationTime      time.Time
	Replicas          int32
	AvailableReplicas int32
	ReadyReplicas     int32
	UpdatedReplicas   int32
	Strategy          string
	Selector          string
	Containers        []ContainerDetails
}

func GetDeploymentDetails(clientset *kubernetes.Clientset, namespace string, deploymentName string) ([]DeploymentDetails, error) {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var deploymentDetailsList []DeploymentDetails

	deploy := deployment
	deploymentDetails := DeploymentDetails{
		Name:              deploy.Name,
		Namespace:         deploy.Namespace,
		CreationTime:      deploy.CreationTimestamp.Time,
		Replicas:          *deploy.Spec.Replicas,
		AvailableReplicas: deploy.Status.AvailableReplicas,
		ReadyReplicas:     deploy.Status.ReadyReplicas,
		UpdatedReplicas:   deploy.Status.UpdatedReplicas,
		Strategy:          string(deploy.Spec.Strategy.Type),
		Selector:          getLabelSelector(deploy.Spec.Selector),
		Containers:        make([]ContainerDetails, len(deploy.Spec.Template.Spec.Containers)),
	}

	for i, container := range deploy.Spec.Template.Spec.Containers {
		containerDetails := ContainerDetails{
			ContainerName: container.Name,
			Ports:         make([]PortDetails, len(container.Ports)),
		}
		for j, port := range container.Ports {
			containerDetails.Ports[j] = PortDetails{
				PortName:      port.Name,
				Protocol:      string(port.Protocol),
				ContainerPort: port.ContainerPort,
			}
		}
		deploymentDetails.Containers[i] = containerDetails
	}

	deploymentDetailsList = append(deploymentDetailsList, deploymentDetails)

	return deploymentDetailsList, nil
}

func getLabelSelector(selector *v1.LabelSelector) string {
	if selector == nil || len(selector.MatchLabels) == 0 {
		return ""
	}

	labelSelector := ""
	for key, value := range selector.MatchLabels {
		labelSelector += key + "=" + value + ","
	}

	// Remove the trailing comma
	labelSelector = labelSelector[:len(labelSelector)-1]

	return labelSelector
}
