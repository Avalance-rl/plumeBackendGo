package initilializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DBsec *gorm.DB
var DBnews *gorm.DB

func ConnectToDB() {
	var err error
	dsn1 := os.Getenv("DBsec")
	dsn2 := os.Getenv("DBnews")

	DBsec, err = gorm.Open(postgres.Open(dsn1), &gorm.Config{})
	if err != nil {
		panic("Ошибка при подключении к БД пользователи")
	}
	DBnews, err = gorm.Open(postgres.Open(dsn2), &gorm.Config{})
	if err != nil {
		panic("Ошибка при подключении к БД новости")
	}
}
