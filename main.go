package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := gin.Default()

	app.LoadHTMLFiles(path.Join("templates", "*.html"))

	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Homepage",
		})
	})

	api := app.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
