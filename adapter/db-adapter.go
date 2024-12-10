package adapter

import (
	"fmt"

	"github.com/Quanghh2233/blogs/config"
	"github.com/Quanghh2233/blogs/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() *gorm.DB {
	dbConfig := config.GetDbConf()
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Port,
		dbConfig.Sslmode,
		dbConfig.Timezone,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection with database")
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Post{}, &model.Comment{})
	return db
}

func CloseDbConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection with Database")
	}
	dbSql.Close()
}
