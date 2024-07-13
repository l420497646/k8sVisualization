package api

import (
	"k8sVisualization/api/example"
	"k8sVisualization/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8sApiGroup     k8s.ApiGroup
}

var AllApiGroup = new(ApiGroup)
