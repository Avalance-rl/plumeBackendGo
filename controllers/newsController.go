package controllers

import (
	"github.com/gin-gonic/gin"
	"plume-backend/initilializers"
	"plume-backend/models"
)

func NewsWireGet(c *gin.Context) ([]models.Article, error) {
	// структура для представления новостей
	type News struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	var newSList []News
	err := initilializers.DBnews.Find(&newSList).Error
	if err != nil {
		panic("ошибка при парсинге БД news")
	}
	var article []models.Article
	result := initilializers.DBsec.Find(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return article, nil
}
