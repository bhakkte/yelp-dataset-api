package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo/bson"
)

func handleUsers(c *gin.Context) {
	userIds := []string{
		"a2RDAvce5iepRZjBVEHk-g",
	}
	users := []data.YelpUser{}

	err := dbSession.DB("").C("users").Find(bson.M{
		"_id": bson.M{"$in": userIds},
	}).All(&users)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"users": users})
	}
}

func handleUser(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{"id": id})
}
