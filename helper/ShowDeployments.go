package helper

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentInfo struct {
	Name      string
	Namespace string
	Ready     string
	Age       string
}

func ShowDeployments(clientset *kubernetes.Clientset, namespace string) ([]DeploymentInfo, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var deploymentList []DeploymentInfo
	for _, deployment := range deployments.Items {

		replicaReady := *deployment.Spec.Replicas
		totalReplica := deployment.Status.ReadyReplicas
		deploymentCreatorTimeStamp := deployment.CreationTimestamp
		age := time.Since(deploymentCreatorTimeStamp.Time).Round(time.Second)

		ready := fmt.Sprintf("%v/%v", replicaReady, totalReplica)

		deploymentInfo := DeploymentInfo{
			Name:      deployment.Name,
			Namespace: string(deployment.Namespace),
			Ready:     ready,
			Age:       age.String(),
		}
		deploymentList = append(deploymentList, deploymentInfo)
	}
	return deploymentList, nil
}
