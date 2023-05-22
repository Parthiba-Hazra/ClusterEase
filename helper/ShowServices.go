package helper

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceInfo struct {
	Name      string
	Namespace string
	Type      string
	IPs       string
	Ports     string
	Age       string
}

func ShowServices(clientset *kubernetes.Clientset, namespace string) ([]ServiceInfo, error) {
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var ServiceInfoList []ServiceInfo
	for _, service := range services.Items {

		var ports []string
		serviceCreationTime := service.CreationTimestamp
		age := time.Since(serviceCreationTime.Time).Round(time.Second)

		for _, port := range service.Spec.Ports {
			portStr := strconv.Itoa(int(port.Port))
			ports = append(ports, portStr)
		}
		ips := fmt.Sprintf("%v | %v", service.Spec.ClusterIP, service.Spec.ExternalIPs)
		portsStr := strings.Join(ports, ", ")

		seviceInfo := ServiceInfo{
			Name:      service.Name,
			Namespace: service.Namespace,
			Type:      string(service.Spec.Type),
			IPs:       ips,
			Ports:     portsStr,
			Age:       age.String(),
		}
		ServiceInfoList = append(ServiceInfoList, seviceInfo)
	}
	return ServiceInfoList, nil
}
