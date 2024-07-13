package initialize

import (
	"k8sVisualization/global"
	"k8sVisualization/router"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()

	{
		exampleGroup := router.AllRouterGroup.ExampleRouterGroup
		exampleGroup.InitExample(r)
	}
	{
		k8sGroup := router.AllRouterGroup.K8sRouterGroup
		k8sGroup.InitK8sRouter(r)
	}

	r.Run(":" + global.CONF.System.Port) // listen and serve on 0.0.0.0:8880 (for windows "localhost:8080")
}
