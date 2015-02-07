package main

import (
	"flag"
	"strconv"

	"github.com/gin-gonic/gin"
)

var port = flag.Int("port", 8080, "Port to listen on")

func main() {
	flag.Parse()

	router := gin.Default()

	router.GET("/", handleDefaultVersionRedirect)
	v1 := router.Group("/v1")
	{
		v1.GET("/", handleHome)
		v1.GET("/users", handleUsers)
		v1.GET("/users/:id", handleUser)
	}

	router.Run(":" + strconv.Itoa(*port))
}

func handleDefaultVersionRedirect(c *gin.Context) {
	c.Redirect(301, "/v1")
}

func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Index of Yelp Dataset Api",
		"endpoints": gin.H{
			"/users":             "Lists few selected interesting users",
			"/users/:id":         "Shows one user's profile",
			"/users/:id/reviews": "Lists one users reviews including associated business info",
		},
	})
}
