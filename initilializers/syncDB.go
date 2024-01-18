package initilializers

import "plume-backend/models"

func SyncDatabase() {
	DBsec.AutoMigrate(&models.User{})
	DBnews.AutoMigrate(&models.Article{})
}
