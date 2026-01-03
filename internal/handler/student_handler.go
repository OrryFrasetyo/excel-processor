package handler

import (
	"excel-processor/internal/worker"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	processor *worker.Processor
}

func NewStudentHandler(processor *worker.Processor) *StudentHandler {
	return &StudentHandler{processor: processor}
}

func (h *StudentHandler) UploadFile(c *gin.Context) {
	filename := "large_students.csv"

	go h.processor.Start(filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "Permintaan diterima! Data sedang diproses di backgroud.",
		"file":    filename,
		"status":  "Check terminal logs for progress",
	})
}
