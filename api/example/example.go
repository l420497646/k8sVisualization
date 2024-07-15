package example

import (
	"k8sVisualization/model/common/response"

	"github.com/gin-gonic/gin"
)

type ExampleApi struct{}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	response.SuccessWithDetailed(c, "请求数据成功", map[string]string{
		"message": "pong",
	})
}
