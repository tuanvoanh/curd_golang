package main

import (
	"fmt"
	"os"

	"API_BASIC_LEARN/configs"
	"API_BASIC_LEARN/models"
	"API_BASIC_LEARN/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// var db *gorm.DB
	fmt.Println(os.Getenv("DB_URL"))
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	db, err := configs.ConnectDatabase("mysql", os.Getenv("DB_URL"))
	fmt.Println(os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("statuse: ", err)
	} else {
		fmt.Println("database connected")
	}

	defer db.Close()
	db.AutoMigrate(&models.Book{})

	r := gin.Default()
	routers.SetupRouter(r)
	fmt.Println(r.Run(":" + os.Getenv("PORT")))

}
