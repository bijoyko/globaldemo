package main

import (
	"log"
	"os"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }
	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("view/*.html")
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", controllers.MainPage)
	router.POST("/next", controllers.ValidateNames)
	router.GET("/next/:id1/:id2/:id3", controllers.LoadForm)
	router.POST("/submit", controllers.Form)
	router.POST("/viewCorrectAnswers", controllers.CorrectAnswers)
	router.POST("/viewTopScorers", controllers.ScoreView)
	router.GET("master-quiztable", controllers.Quiztable)
	router.GET("/admin-bijoy", controllers.LoadAdmin)
	router.POST("/adminlogin", controllers.VerifyAdmin)
	router.POST("/updatequestions", controllers.UpdateQuestions)
	// router.Run(":" + port)
	router.Run(":8080")
}
