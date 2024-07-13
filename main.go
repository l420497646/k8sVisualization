package main

import (
	"k8sVisualization/api"

	"github.com/gin-gonic/gin"
)

// 项目启动入口
func main() {
	r := gin.Default()
	allApiGroup := api.AllApiGroup
	exampleApiGroup := allApiGroup.ExampleApiGroup

	r.GET("/ping", exampleApiGroup.ExampleTest)

	r.Run(":8889") // listen and serve on 0.0.0.0:8880 (for windows "localhost:8080")
}
