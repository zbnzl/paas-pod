package dao

import (
	"github.com/zbnzl/paas-pod/internal/conf"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewK8sClient(data *conf.Data) *kubernetes.Clientset {
	// 获取kubeconfig文件路径
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	if data.GetK8S().GetConfigPath() != "" {
		kubeconfig = data.GetK8S().GetConfigPath()
	}
	// 根据kubeconfig文件创建Kubernetes客户端配置
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建Kubernetes客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
