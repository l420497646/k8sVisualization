package initialize

import (
	"k8sVisualization/router"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()

	{
		exampleGroup := router.AllRouterGroup.ExampleRouterGroup
		exampleGroup.InitExample(r)
	}

	r.Run(":8889") // listen and serve on 0.0.0.0:8880 (for windows "localhost:8080")
}
