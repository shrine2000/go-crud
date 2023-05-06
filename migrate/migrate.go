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
	err := intializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return 
	}
}
