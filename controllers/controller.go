package controller

import (
	"context"
    "log"
    "os"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"upload_api/config"
)


func UploadHandler(c *gin.Context){
	ctx := context.Background()
    bucketName := os.Getenv("MINIO_BUCKET")
    file, err := c.FormFile("files")

	if err != nil {
		return
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return
	}
	defer buffer.Close()

	minioClient, err := minioUpload.MinioConnection()
    if err != nil {
		return
	}

	objectName := file.Filename
    fileBuffer := buffer
    contentType := file.Header["Content-Type"][0]
    fileSize := file.Size

    // Upload the zip file with PutObject
    info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

    if err != nil {
		return
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}


