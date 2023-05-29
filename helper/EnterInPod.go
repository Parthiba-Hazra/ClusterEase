package helper

import (
	"context"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/remotecommand"
)

func EnterInPod(clientset *kubernetes.Clientset, podName, namespace, command string) error {
	// Get the pod details
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		return err
	}

	// Set up the command to execute within the pod
	cmd := []string{
		"/bin/sh",
		"-c",
		command,
	}

	// Create an executor for running the command
	executor, err := remotecommand.NewSPDYExecutor(clientset.Config, "POST", fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/exec", namespace, podName), []string{"sh"}, os.Stdout, os.Stderr)

	// Create the exec request
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Command: cmd,
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
		}, v1.ParameterCodec)

	// Prepare the request to execute the command
	executor.StreamOptions.Tty = false
	executor.StreamOptions.Stdin = nil

	// Execute the command within the pod
	err = executor.Stream(remotecommand.StreamOptions{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		return err
	}

	return nil
}
