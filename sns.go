package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/stor-inwinstack/kaoliang/pkg/config"
	"gitlab.com/stor-inwinstack/kaoliang/pkg/controllers"
	"gitlab.com/stor-inwinstack/kaoliang/pkg/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	config.SetServerConfig()
	models.SetDB()
	models.Migrate()
	models.SetCache()
}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		action := c.PostForm("Action")
		switch action {
		case "CreateTopic":
			controllers.CreateTopic(c)
		case "ListTopics":
			controllers.ListTopics(c)
		case "DeleteTopic":
			controllers.DeleteTopic(c)
		case "Subscribe":
			controllers.Subscribe(c)
		}
	})

	r.Run()
}
