package k8s

import (
	"context"
	"fmt"
	"k8sVisualization/global"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodApi struct{}

func (*PodApi) GetPodListApi(c *gin.Context) {
	ctx := context.TODO()

	list, err := global.K8SClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	for _, pod := range list.Items {
		fmt.Println(pod.Name, " ", pod.Namespace)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
	})
}
