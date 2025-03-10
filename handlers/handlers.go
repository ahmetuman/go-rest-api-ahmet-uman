package handlers

import (
	"my-gin-api/database"
	"my-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	database.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find author"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find author"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&author)
	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find author"})
		return
	}
	database.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Preload("Author").Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Preload("Author").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var author models.Author
	if err := database.DB.First(&author, book.AuthorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find author for this id"})
		return
	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hata"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find book"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find book"})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func GetReviews(c *gin.Context) {
	var reviews []models.Review
	database.DB.Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

func GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := database.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find review"})
		return
	}
	c.JSON(http.StatusOK, review)
}

func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := database.DB.First(&book, review.BookID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find book for this ID"})
		return
	}

	database.DB.Create(&review)
	c.JSON(http.StatusCreated, review)
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := database.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find review"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&review)
	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := database.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find review"})
		return
	}
	database.DB.Delete(&review)
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}
