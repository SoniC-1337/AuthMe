package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadHandler(c *gin.Context) {
	filePath := "file.bin"
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to open file",
		})
		return
	}
	defer file.Close()

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=file.bin")

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to seek file",
		})
		return
	}
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to copy file",
		})
	}
}
