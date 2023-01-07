package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseId(c *gin.Context) (id int64) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	return
}

func getRequests(c *gin.Context) {
	id := parseId(c)
	requests := GetRequests(id)
	c.IndentedJSON(http.StatusOK, requests)
}

func incRequests(c *gin.Context) {
	id := parseId(c)
	IncRequests(id)
}

func main() {
	router := gin.Default()
	router.GET("/requests", getRequests)
	router.POST("/requests", incRequests)

	router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
