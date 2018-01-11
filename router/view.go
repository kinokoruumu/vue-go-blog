package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func viewRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
