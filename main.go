package main

import (
	"log"
	"os"

	"github.com/bijoyko/globaldemo/controller"
	"github.com/bijoyko/globaldemo/driver"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*.html")
	db := driver.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", controller.MainPage)
	router.GET("/admin-demo", controller.LoadAdmin)
	router.POST("/adminlogin", controller.VerifyAdmin)
	router.POST("/updatequestions", controller.UpdateLinks)
	router.Run(":" + port)
}
