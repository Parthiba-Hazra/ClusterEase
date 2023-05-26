package helper

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteResource(clientset *kubernetes.Clientset, kind string, name string, namespace string) error {
	kind = strings.ToLower(kind)

	switch kind {
	case "deployment":
		err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "service":
		err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "namespace":
		err := clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "configmap":
		err := clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "secret":
		err := clientset.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "statefulset":
		err := clientset.AppsV1().StatefulSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "daemonset":
		err := clientset.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "job":
		err := clientset.BatchV1().Jobs(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}

	case "cronjob":
		err := clientset.BatchV1beta1().CronJobs(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported resource kind: %s", kind)
	}

	fmt.Printf("Resource deleted: kind=%s, name=%s, namespace=%s\n", kind, name, namespace)
	return nil
}
