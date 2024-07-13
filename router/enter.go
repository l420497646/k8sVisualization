package router

import (
	"k8sVisualization/router/example"
)

type RouterGroup struct {
	ExampleRouterGroup example.RouterGroup
}

var AllRouterGroup = new(RouterGroup)
