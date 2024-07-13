package k8s

import (
	"k8sVisualization/api"

	"github.com/gin-gonic/gin"
)

type K8sRouter struct{}

func (*K8sRouter) InitK8sRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	k8sApiGroup := api.AllApiGroup.K8sApiGroup
	group.GET("/getPodList", k8sApiGroup.GetPodListApi)
}
