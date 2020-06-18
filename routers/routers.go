package routers

import (
	"API_BASIC_LEARN/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter ...
func SetupRouter(r *gin.Engine, DB *gorm.DB) {
	r.Use(Database(DB))
	v1 := r.Group("/v1")
	{
		// super simple code
		v1.GET("hahaha", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "posted",
			})
		})
		// code thiet nam o day
		v1.GET("book", controllers.ListBook)
		v1.POST("book", controllers.AddNewBook)
		v1.GET("book/:id", controllers.GetOneBook)
		v1.PUT("book/:id", controllers.PutOneBook)
		v1.DELETE("book/:id", controllers.DeleteBook)
	}
}

func Database(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", DB)
		c.Next()
	}
}
