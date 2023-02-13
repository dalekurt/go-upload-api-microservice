package main

import (
	"log"
    _ "github.com/joho/godotenv/autoload" // load .env file automatically
    "upload_api/config"
	"upload_api/routers"
)

func main() {

	minioUpload.MinioConnection()

	router := routers.SetupRouter()
	err := router.Run(":8080") //TODO Use environment variable PORT

	if err != nil {
		log.Fatalf("Server error: %s", err)
	}  
}
