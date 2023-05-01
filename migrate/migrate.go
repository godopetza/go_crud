package main

import (
	"github.com/godopetza/go-crud/initializers"
	"github.com/godopetza/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
