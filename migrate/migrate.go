package main

import (
	"github.com/johnsoncwb/go-crud/initializers"
	"github.com/johnsoncwb/go-crud/models"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
