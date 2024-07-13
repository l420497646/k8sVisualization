package api

import "k8sVisualization/api/example"

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
}

var AllApiGroup = new(ApiGroup)
