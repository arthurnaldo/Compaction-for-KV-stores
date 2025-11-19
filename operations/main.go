package main

import (
	"time"

	"net/http"
	"operations/compact"
	"operations/get"
	"operations/hash"
	"operations/set"
	"operations/state"

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

	set.Set(k, v, state.KeyDict)

}

// subprocess that listens to get requests, and sends them to get module
func get_request(c *gin.Context) {
	key, key_err := c.Get("key")
	if !key_err {
		c.JSON(http.StatusBadRequest, "no key or value provided")
	}

	k, ok := key.(string)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "key must be a string",
		})
		return
	}

	value, err := get.Get(k)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": value})

}

// runs compact asyncronously every 10 seconds
func run_compact() {
	time.Sleep(10000)
	compact.Compact()
}

func main() {
	state.KeyDict = hash.CreateHashMap(1000)
	router := gin.Default()
	router.GET("/value", get_request)
	router.POST("keyvalue", set_request)
	router.Run("localhost:8080")

}
