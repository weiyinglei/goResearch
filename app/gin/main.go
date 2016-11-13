package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()


	r.GET("/", func(c *gin.Context){
		c.JSON(200,gin.H{
			"message": "index",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})


	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//r.Run()
	r.Run(":80")
	// router.Run(":3000") for a hard coded port
}