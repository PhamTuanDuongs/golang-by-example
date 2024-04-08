package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleReadingExcelFile(c *gin.Context) {
    fmt.Println("File upload Endpoint Hit")

    // Parse multipart form
    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
        fmt.Println("Failed to parse form:", err)
        return
    }

    // Get the file
    file, err := c.FormFile("coverPage")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form"})
        fmt.Println("Failed to get file from form:", err)
        return
    }

    // Print file details
    fmt.Println("Received file:", file.Filename)
    fmt.Println("File header:", file.Header)

    // Send response
    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}


func main() {
  r := gin.Default()
 r.Use(cors.Default())
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.POST("/uploadFile", HandleReadingExcelFile)
  r.Run(":8889") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}