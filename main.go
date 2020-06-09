package main

import (
	"fmt"

	"./configs"
	"./models"
	"./routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// var db *gorm.DB
	db, err := configs.ConnectDatabase("mysql", "root:Tuan4120@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	fmt.Println("database connected")
	defer db.Close()
	db.AutoMigrate(&models.Book{})

	r := gin.Default()
	routers.SetupRouter(r)
	fmt.Println(r.Run(":3000"))

}
