package controllers

import (
	"github.com/gin-gonic/gin"
	"plume-backend/initilializers"
)

func NewsWireGet(c *gin.Context) {
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

}
