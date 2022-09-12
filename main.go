package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/go-crud/controllers"
	"github.com/johnsoncwb/go-crud/initializers"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.ConnectToDB()
}

func main() {
	route := gin.Default()

	route.POST("/posts", controllers.PostsCreate)
	route.GET("/posts", controllers.PostsIndex)
	route.GET("posts/:id", controllers.PostsShow)
	route.PUT("posts/:id", controllers.PostsUpdate)
	route.DELETE("posts/:id", controllers.PostDelete)
	route.GET("/alive", controllers.Alive)

	err := route.Run()

	if err != nil {
		log.Fatal(err)
	}
}
