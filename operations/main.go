package main

import (
	"operations/compact"

	"github.com/gin-gonic/gin"
)

// subprocess that listens to set requests, and sends them to set module
func set_request(c *gin.Context) {

}

// subprocess that listens to get requests, and sends them to get module
func get_request(c *gin.Context) {

}

func main() {
	compact.Compact()

	router := gin.Default()
	router.GET("/value", get_request)
	router.POST("keyvalue", set_request)
	router.Run("localhost:8080")

}
