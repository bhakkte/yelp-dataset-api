package main

import (
	"github.com/gin-gonic/gin"
)

func handleUsers(c *gin.Context) {
	userIds := []string{"a2RDAvce5iepRZjBVEHk-g"}
	c.JSON(200, gin.H{"ids": userIds})
}

func handleUser(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{"id": id})
}
