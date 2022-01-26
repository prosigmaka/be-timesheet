package db

import (
	"log"
	"timesheet-app/pkg/models"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// username := viper.GetString("Database.Username")
	// password := viper.GetString("Database.Password")
	// host := viper.GetString("Database.Host")
	// port := viper.GetInt("Database.Port")
	// dbname := viper.GetString("Database.DBName")

	// dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/jakarta`, host, username, password, dbname, port)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	return nil, err
	// }

	// // if err := db.DB().Ping(); err != nil {
	// // 	return nil, err
	// // }

	// return db, nil

	dbURL := "postgres://postgres:farhan5497@localhost:5432/timesheet-be"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Project{})

    return db



}