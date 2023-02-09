package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"upload_api/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy"})    
	  })
	
	router.POST("/", controller.UploadHandler)
	
	return router
}
