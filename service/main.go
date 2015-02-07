package main

import (
	"flag"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo"
)

var dbSession *mgo.Session

var port = flag.Int("port", 8080, "Port to listen on")
var dbUrl = flag.String("mongo-url", "mongodb://localhost:27017/yelp-dataset-api", "MongoDB url")

func main() {
	flag.Parse()

	dbSession = dialMongo(*dbUrl)
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

func dialMongo(url string) *mgo.Session {
	dbSession, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	data.Index(dbSession.DB(""))
	return dbSession
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
