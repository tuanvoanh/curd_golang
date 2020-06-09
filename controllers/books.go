package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}

func AddNewBook(c *gin.Context) {
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
	err := models.AddNewBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"result": book})
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := models.GetOneBook(&book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"result": book})
	}
}

func PutOneBook(c *gin.Context) {
	var book models.Book
	var input CreateRequest
	id := c.Params.ByName("id")
	err := models.GetOneBook(&book, id)
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

	models.PutOneBook(&book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(&book, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": book})
	}
}
