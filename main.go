package main

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Time struct {
	TimeFormat	string `json:"time_format"`
	CurrentTime string `json:"current_time"`
}

func main() {
	// Gin Router
	router := gin.Default()
	router.GET("/time", getTime)

	// starting server
	router.Run(fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("PORT")))
}

func getTime(c *gin.Context) {
	currentTime := []Time{
		{
			TimeFormat: http.TimeFormat,
			CurrentTime: time.Now().Format(time.RFC3339),
		},
	}

	c.IndentedJSON(http.StatusOK, currentTime)
}

