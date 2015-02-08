package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo/bson"
)

func getQueryFloat(c *gin.Context, key string) (float32, error) {
	rawFloat, err := strconv.ParseFloat(c.Request.URL.Query().Get(key), 32)
	return float32(rawFloat), err
}

func handleLocationWords(c *gin.Context) {
	// Start by fetching businesses id's
	var lat, lng float32
	var err error
	if lat, err = getQueryFloat(c, "lat"); err != nil {
		c.JSON(400, gin.H{"message": "Can't parse latitude"})
		return
	}
	if lng, err = getQueryFloat(c, "lng"); err != nil {
		c.JSON(400, gin.H{"message": "Can't parse longitude"})
		return
	}

	businesses := []data.YelpBusiness{}
	err = dbSession.DB("").C("businesses").Find(bson.M{
		"loc": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float32{lng, lat},
				},
				"$maxDistance": 5000,
			},
		},
	}).Select(bson.M{"_id": 1}).Limit(300).All(&businesses)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	businessIds := make([]string, len(businesses))
	for _, b := range businesses {
		businessIds = append(businessIds, b.Id)
	}

	// Now get reviews for those
	reviews := []data.YelpReview{}
	err = dbSession.DB("").C("reviews").Find(bson.M{
		"business_id": bson.M{"$in": businessIds},
	}).Select(bson.M{"text": 1}).All(&reviews)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	} else {
		figureOutPopularWords(c, reviews)
	}
}

func handleBusinessWords(c *gin.Context) {
	id := c.Params.ByName("id")
	reviews := []data.YelpReview{}

	err := dbSession.DB("").C("reviews").Find(bson.M{
		"business_id": id,
	}).All(&reviews)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	} else {
		figureOutPopularWords(c, reviews)
	}
}

func figureOutPopularWords(c *gin.Context, reviews []data.YelpReview) {
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

	sortedWords := sortMapByValue(words)
	// Cap words returned at 200 (so we dont get a dictionary of 36 000 w)
	if len(sortedWords) >= 200 {
		sortedWords = sortedWords[:200]
	}

	c.JSON(200, gin.H{"words": sortedWords})
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
