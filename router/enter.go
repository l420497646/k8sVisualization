package router

import (
	"k8sVisualization/router/example"
	"k8sVisualization/router/k8s"
)

type RouterGroup struct {
	ExampleRouterGroup example.RouterGroup
	K8sRouterGroup     k8s.RouterGroup
}

var AllRouterGroup = new(RouterGroup)
