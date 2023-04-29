package main

import (
	"go-crud/m/intializers"
	"go-crud/m/models"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()
}

func main() {
	intializers.DB.AutoMigrate(&models.Post{})
}
