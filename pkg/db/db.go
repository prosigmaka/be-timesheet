package db

import (
	"be-timesheet/pkg/entity"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	host := viper.GetString("Database.Host")
	port := viper.GetInt("Database.Port")
	dbname := viper.GetString("Database.DBName")

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta`, host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Timesheet{})

	return db

}
