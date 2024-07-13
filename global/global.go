package global

import (
	"k8sVisualization/config"

	"k8s.io/client-go/kubernetes"
)

var (
	CONF         config.Config
	K8SClientSet *kubernetes.Clientset
)
