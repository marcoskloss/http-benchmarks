package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"log"
)

func main() {
  r := gin.New()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "ok": "true",
    })
  })

	log.Println("go/gin listening on port 3004");
  r.Run(":3004")
}

// Run using `GIN_MODE=release go run gin.go`
