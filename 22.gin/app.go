package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.MaxMultipartMemory = 8 << 20 // 8 MiB
	server.StaticFile("/", "./public")
	server.POST("/upload", uploadFile)

	server.Run(":8080")

}

func uploadFile(context *gin.Context) {
	name := context.PostForm("name")
	email := context.PostForm("email")

	// source
	file, err := context.FormFile("file")
	log.Printf("filename %s", file.Filename)

	if err != nil {
		context.String(http.StatusBadRequest, "got form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	log.Printf("filenam %s", file.Filename)
	if err := context.SaveUploadedFile(file, filename); err != nil {
		context.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	context.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
