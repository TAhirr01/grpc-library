package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := "root:root@tcp(127.0.0.1:3306)/grpc-library?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db baglantisi alinmadi: ", err)
	}
	DB = database
}
