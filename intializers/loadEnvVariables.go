package intializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	err:= godotenv.Load()
	if err!=nil{
		log.Fatal("Error loginf .env")
	}
}