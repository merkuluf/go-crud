package main

import (
	"github/merkuluf/go-crud/initializers"
	"github/merkuluf/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
