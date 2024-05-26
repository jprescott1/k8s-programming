package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)

	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodeList.Items {
		fmt.Println(n.Name)
	}

	newPod := corev1.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-pod",
			Labels: map[string]string{
				"foo": "bar",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: "busybox", Image: "busybox:latest", Command: []string{"sleep", "100000"}},
			},
		},
	}

	fmt.Println(newPod)

	Pod, _ := clientset.CoreV1().Pods("default").Create(context.Background(), &newPod, v1.CreateOptions{})

	fmt.Println(Pod)
}
