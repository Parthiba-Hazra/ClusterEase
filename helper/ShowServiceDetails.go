package helper

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceDetails struct {
	Name            string
	Namespace       string
	CreationTime    time.Time
	Labels          map[string]string
	Type            string
	ClusterIP       string
	ExternalIPs     []string
	LoadBalancerIP  string
	Ports           []ServicePortDetails
	Selector        map[string]string
	SessionAffinity string
}

type ServicePortDetails struct {
	Name       string
	Protocol   string
	Port       int32
	TargetPort string
	NodePort   int32
}

func ShowServiceDetails(clientset *kubernetes.Clientset, namespace string, serviceName string) ([]ServiceDetails, error) {
	service, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var serviceDetailsList []ServiceDetails

	serviceDetails := ServiceDetails{
		Name:            service.Name,
		Namespace:       service.Namespace,
		CreationTime:    service.CreationTimestamp.Time,
		Labels:          service.Labels,
		Type:            string(service.Spec.Type),
		ClusterIP:       service.Spec.ClusterIP,
		ExternalIPs:     service.Spec.ExternalIPs,
		LoadBalancerIP:  service.Spec.LoadBalancerIP,
		Ports:           make([]ServicePortDetails, len(service.Spec.Ports)),
		Selector:        service.Spec.Selector,
		SessionAffinity: string(service.Spec.SessionAffinity),
	}

	for i, port := range service.Spec.Ports {
		servicePortDetails := ServicePortDetails{
			Name:       port.Name,
			Protocol:   string(port.Protocol),
			Port:       port.Port,
			TargetPort: port.TargetPort.String(),
			NodePort:   port.NodePort,
		}
		serviceDetails.Ports[i] = servicePortDetails
	}

	serviceDetailsList = append(serviceDetailsList, serviceDetails)

	return serviceDetailsList, nil
}
