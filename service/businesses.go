package main

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo/bson"
)

func handleBusinessWords(c *gin.Context) {
	id := c.Params.ByName("id")
	reviews := []data.YelpReview{}

	err := dbSession.DB("").C("reviews").Find(bson.M{
		"business_id": id,
	}).All(&reviews)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	} else {
		words := map[string]int{}

		for _, review := range reviews {
			r := strings.NewReplacer(",", "", ".", "", ":", "", ";", "", "!", "",
				"that", "", "this", "", "with", "", "have", "", "they", "", "here", "",
				"just", "", "where", "", "their", "", "it's", "", "were", "", "them", "",
				"from", "")
			reviewWords := strings.Split(r.Replace(review.Text), " ")
			for _, reviewWord := range reviewWords {
				if len(reviewWord) > 3 {
					words[reviewWord]++
				}
			}
		}

		c.JSON(200, gin.H{"business_id": id, "words": sortMapByValue(words)})
	}
}

func sortWords(m map[string]int) map[string]int {
	pairList := sortMapByValue(m)
	sortedMap := make(map[string]int, len(pairList))
	for _, val := range pairList {
		sortedMap[val.Key] = val.Value
	}
	return sortedMap
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(&p)
	return p
}
