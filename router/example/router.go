package example

import (
	"k8sVisualization/api"

	"github.com/gin-gonic/gin"
)

type ExampleRouter struct{}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	group := r.Group("/example")
	exampleApiGroup := api.AllApiGroup.ExampleApiGroup
	group.GET("/ping", exampleApiGroup.ExampleTest)
}
