package main

import (
	"log"

	"github.com/GURUAKASH-MUTHURAJAN/swagger/controllers"
	_ "github.com/GURUAKASH-MUTHURAJAN/swagger/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API (Your API Title)
// @version 1
// @Description Sample description

// @contact.name Guru Akash
// @contact.url https://github.com/GURUAKASH-MUTHURAJAN
// @contact.email guuakashsm@gmail.com


// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	user := v1.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
