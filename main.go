package main

import (
	"my-gin-api/database"
	"my-gin-api/handlers"
	"my-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "my-gin-api/docs"
)

// @title Go API with Gin
// @version 1.0
// @description This is a simple API using Gin and PostgreSQL
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	// Tables
	database.DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Review{})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Docker ile test et ?

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/authors", handlers.GetAuthors)
	r.GET("/authors/:id", handlers.GetAuthorByID)
	r.POST("/authors", handlers.CreateAuthor)
	r.PUT("/authors/:id", handlers.UpdateAuthor)
	r.DELETE("/authors/:id", handlers.DeleteAuthor)

	r.GET("/reviews", handlers.GetReviews)
	r.GET("/reviews/:id", handlers.GetReviewByID)
	r.POST("/reviews", handlers.CreateReview)
	r.PUT("/reviews/:id", handlers.UpdateReview)
	r.DELETE("/reviews/:id", handlers.DeleteReview)

	r.Run(":8080")
}
