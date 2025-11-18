package main

import (
	"operations/get"
	"operations/set"

	"github.com/gin-gonic/gin"
)

// subprocess that listens to set requests, and sends them to set module
func set_request(c *gin.Context) {
	key, key_err := c.Get("key")
	value, val_err := c.Get("value")
	if key_err || val_err {
		panic("no key or value provided")
	}

	k := key.(string)
	v := value.(string)

	set.Set(k, v)

}

// subprocess that listens to get requests, and sends them to get module
func get_request(c *gin.Context) {
	key, key_err := c.Get("key")
	if key_err {
		panic("no key or value provided")
	}

	k := key.(string)

	println(get.Get(k))

}

func main() {

	router := gin.Default()
	router.GET("/value", get_request)
	router.POST("keyvalue", set_request)
	router.Run("localhost:8080")

}
