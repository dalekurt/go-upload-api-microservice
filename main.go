package main

import (
    _ "github.com/joho/godotenv/autoload" // load .env file automatically
    "upload_api/config"
	"upload_api/routers"
)

func main() {

	minioUpload.MinioConnection()
	
	router := routers.SetupRouter()
	err := router.Run()

	if err != nil {
		return
	}  
}
