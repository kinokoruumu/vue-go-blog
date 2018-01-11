package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kinokoruumu/vue-go-blog/controller"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")
	//ウェルカムページ
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api")

	api.GET("/articles", controller.GetArticleAll)
	api.GET("/articles/:id", controller.GetArticle)

	r.Run(":3000")
}
