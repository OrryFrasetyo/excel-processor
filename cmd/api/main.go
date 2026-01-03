package main

import (
	"excel-processor/internal/handler"
	"excel-processor/internal/repository"
	"excel-processor/internal/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB()

	studentRepo := repository.NewStudentRepository(repository.DB)
	processor := worker.NewProcessor(studentRepo)
	studentHandler := handler.NewStudentHandler(processor)

	r := gin.Default()

	r.POST("/process-data", studentHandler.UploadFile)

	// check health server
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
