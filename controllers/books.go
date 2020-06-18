package controllers

import (
	"fmt"
	"net/http"

	"API_BASIC_LEARN/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListBook(c *gin.Context) {
	DB := c.MustGet("DB").(*gorm.DB)
	var book []models.Book
	err := models.GetAllBook(DB, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}

func AddNewBook(c *gin.Context) {
	DB := c.MustGet("DB").(*gorm.DB)
	var book models.Book
	var input CreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book = models.Book{
		Author:   input.Author,
		Name:     input.Name,
		Category: input.Category,
		Price:    input.Price,
	}
	err := models.AddNewBook(DB, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"result": book})
	}
}

func GetOneBook(c *gin.Context) {
	DB := c.MustGet("DB").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	err := models.GetOneBook(DB, &book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"result": book})
	}
}

func PutOneBook(c *gin.Context) {
	DB := c.MustGet("DB").(*gorm.DB)
	var book models.Book
	var input CreateRequest
	id := c.Params.ByName("id")
	err := models.GetOneBook(DB, &book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.Author = input.Author
	book.Name = input.Name
	book.Category = input.Category
	book.Price = input.Price

	models.PutOneBook(DB, &book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}

func DeleteBook(c *gin.Context) {
	DB := c.MustGet("DB").(*gorm.DB)
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(DB, &book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}
