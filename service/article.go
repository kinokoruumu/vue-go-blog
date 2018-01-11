package service

import "github.com/kinokoruumu/vue-go-blog/model"

var Article = articleImpl{}

type articleImpl struct {
}

func (self articleImpl) Create(article model.Article) model.Article {
	err := db.Create(&article).Error
	if err != nil {
		panic(err)
	}
	return article
}

func (self articleImpl) All() []model.Article {
	var articles []model.Article
	db.Find(&articles)
	return articles
}

func (self articleImpl) Find(id uint) *model.Article {
	var article []model.Article
	db.Find(&article, id)
	if len(article) == 0 {
		return nil
	}
	return &article[0]
}
