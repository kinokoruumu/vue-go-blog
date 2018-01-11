package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kinokoruumu/vue-go-blog/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/articles", controller.GetArticleAll)
	api.GET("/articles/:id", controller.GetArticle)
}
