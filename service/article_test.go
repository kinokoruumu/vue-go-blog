package service

import (
	"testing"

	"fmt"

	"github.com/kinokoruumu/vue-go-blog/model"
)

func TestArticleImpl_Create(t *testing.T) {
	Article.Create(model.Article{
		Title: "What is Lorem Ipsum?",
		Text:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	})
}

func TestArticleImpl_All(t *testing.T) {
	articles := Article.All()
	fmt.Println(articles)
}

func TestArticleImpl_Find(t *testing.T) {
	article := Article.Find(1)
	fmt.Println(article)
}
