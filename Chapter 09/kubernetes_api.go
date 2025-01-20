package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    v1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
    // Need to have kube config setup
    config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
    if err != nil {
        panic(err)
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    pod := &v1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name: "demo-pod",
        },
        Spec: v1.PodSpec{
            Containers: []v1.Container{
                {
                    Name:  "busybox",
                    Image: "busybox",
                    Command: []string{
                        "sleep", "3600",
                    },
                },
            },
        },
    }

    pod, err = clientset.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
    if err != nil {
        panic(err)
    }

    fmt.Printf("Pod created: %s\n", pod.ObjectMeta.Name)


    pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err)
    }

    for _, pod := range pods.Items {
      fmt.Printf("Pod Name: %s, Status: %s\n", pod.Name, pod.Status.Phase)
    }

    logOptions := &v1.PodLogOptions{
        Container: "my-container-name", // Specify container if the pod has multiple containers
        Follow:    true,    // Stream logs
    }
    req := clientset.CoreV1().Pods("default").GetLogs("pod-name", logOptions)
    podLogs, err := req.Stream(context.TODO())

    if err != nil {
        panic(err)
    }
    defer podLogs.Close()

    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, podLogs)
    if err != nil {
        panic(err)
    }

    fmt.Println("Logs:")
    fmt.Println(buf.String())
}
