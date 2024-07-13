package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleApi struct{}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
