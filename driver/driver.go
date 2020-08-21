package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SetupModels() *gorm.DB {
	user := getEnv("PG_USER", "ppvhvfwd")
	password := getEnv("PG_PASSWORD", "jbcaUVTtUIa_WqALGc7F5DZgQArj_UVa")
	host := getEnv("PG_HOST", "john.db.elephantsql.com")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "ppvhvfwd")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	db.SingularTable(true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection established")
	//db.Table("quiz-data").AutoMigrate(&models.Quiz{})
	return db
}
