package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //this is to get the dialects
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
	user := getEnv("PG_USER", "ckuhwdif")
	password := getEnv("PG_PASSWORD", "vI0SVzF5_h8C0GPI-mJ9KS-Gtyf1IVUp")
	host := getEnv("PG_HOST", "lallah.db.elephantsql.com")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "ckuhwdif")

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

// CREATE TABLE links (btnNo int, btnName VARCHAR(100), webLink VARCHAR(100));
// INSERT INTO links(btnNo,btnName,webLink)
// VALUES(1,'Dashboard Tracking Alerts','https://xd.adobe.com/view/f36a8abc-fbbb-4be5-8f22-8a1e6648d3b8-a639/');
// INSERT INTO links(btnNo,btnName,webLink)
// VALUES(2,'Parts Tools Tags','https://xd.adobe.com/view/f36a8abc-fbbb-4be5-8f22-8a1e6648d3b8-a639/');
// INSERT INTO links(btnNo,btnName,webLink)
// VALUES(3,'Forecast Orders Reports','https://xd.adobe.com/view/f36a8abc-fbbb-4be5-8f22-8a1e6648d3b8-a639/');
// INSERT INTO links(btnNo,btnName,webLink)
// VALUES(4,'Mobile App','https://xd.adobe.com/view/f36a8abc-fbbb-4be5-8f22-8a1e6648d3b8-a639/');
// INSERT INTO links(btnNo,btnName,webLink)
// VALUES(5,'iPad App','https://xd.adobe.com/view/f36a8abc-fbbb-4be5-8f22-8a1e6648d3b8-a639/');
