/**
 * @author yantao
 * @date 2024/3/29
 * @description 获取容器日志
 */
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 创建 Kubernetes 客户端实例
func createKubeClient() (*kubernetes.Clientset, error) {
	// 如果在集群内部运行，可以直接使用 in-cluster 配置
	config, err := rest.InClusterConfig()
	if err != nil {
		// 如果不在集群内部，使用 kubeconfig 文件配置
		config, err = clientcmd.BuildConfigFromFlags("", "/home/yantao/.kube/config31")
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// 获取指定容器的日志
func getPodLogs(namespace, podName, containerName string) (string, error) {
	clientset, err := createKubeClient()
	if err != nil {
		return "", fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	logOpts := &v1.PodLogOptions{
		Container: containerName,
		Follow:    false,            // 根据需求设置是否实时追加日志
		TailLines: &[]int64{100}[0], // 可选，只获取最近的100行日志
		//SinceSeconds: &[]int64{60}[0],  // 可选，获取过去60秒内的日志
		Timestamps: true,
	}

	request := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOpts)
	logStream, err := request.Stream(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get logs for pod %s/%s: %w", namespace, podName, err)
	}
	defer logStream.Close()

	logData, err := ioutil.ReadAll(logStream)
	if err != nil {
		return "", fmt.Errorf("failed to read logs: %w", err)
	}

	return string(logData), nil
}

func main() {
	namespace := "mxsche"
	podName := "mxschejob-lxhhm-fgsgk"
	containerName := "mxsche-container-646cca66"

	logs, err := getPodLogs(namespace, podName, containerName)
	if err != nil {
		fmt.Println("Error fetching logs:", err)
	} else {
		fmt.Println(string(logs))
	}
}
