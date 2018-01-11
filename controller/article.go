package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kinokoruumu/vue-go-blog/model"
	"github.com/kinokoruumu/vue-go-blog/service"
)

type ArticleRes struct {
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateArticle(c *gin.Context) {
	article, ok := createArticleReq(c)
	if !ok {
		return
	}
	article = service.Article.Create(article)
	c.JSON(http.StatusCreated, article)
}

func createArticleReq(c *gin.Context) (model.Article, bool) {
	article := model.Article{
		Title: c.PostForm("title"),
		Text:  c.PostForm("text"),
	}
	// ここでバリデーション

	return article, true
}

func createArticleRes(m []model.Article) []ArticleRes {
	var res []ArticleRes
	for _, value := range m {
		res = append(res, ArticleRes{
			Title:     value.Title,
			Text:      value.Text,
			CreatedAt: value.CreatedAt,
		})
	}
	return res
}

func GetArticleAll(c *gin.Context) {
	ArticleAll := service.Article.All()
	c.JSON(http.StatusOK, gin.H{
		"articles": createArticleRes(ArticleAll),
	})
}

func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "不正なリクエストです",
		})
		return
	}
	article := service.Article.Find(uint(articleID))
	if article == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, article)
}
