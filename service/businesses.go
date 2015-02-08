package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	ksort "github.com/kiasaki/batbelt/sort"
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
	}).Select(bson.M{"business_id": 1, "text": 1}).All(&reviews)

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
	words := map[string]Word{}

	for _, review := range reviews {
		r := strings.NewReplacer(",", "", ".", "", ":", "", ";", "", "!", "",
			"that", "", "this", "", "with", "", "have", "", "they", "", "here", "",
			"just", "", "where", "", "their", "", "it's", "", "were", "", "them", "",
			"from", "", "than", "", "this", "", "don't", "", "didn't", "", "i've", "",
			"it's", "")
		reviewWords := strings.Split(r.Replace(strings.ToLower(review.Text)), " ")
		for _, reviewWord := range reviewWords {
			if len(reviewWord) > 3 {
				val, ok := words[reviewWord]
				if !ok {
					val = Word{reviewWord, 0, map[string]int{}, []ksort.Pair{}}
				}
				val.Occurences++
				val.Businesses[review.BusinessId]++
				words[reviewWord] = val
			}
		}
	}

	// Sort words by descending occurences
	sortedWords := sortMapByValue(words)
	// Cap words returned at 200 (so we dont get a dictionary of 36 000 w)
	if len(sortedWords) >= 200 {
		sortedWords = sortedWords[:200]
	}

	// Sort businesses in those top 200 words
	for i, word := range sortedWords {
		word.SortedBusinesses = ksort.SortMapByValue(word.Businesses)
		if len(word.SortedBusinesses) > 5 {
			word.SortedBusinesses = word.SortedBusinesses[:5]
		}
		word.Businesses = nil
		sortedWords[i] = word
	}

	c.JSON(200, gin.H{"words": sortedWords})
}

type Word struct {
	Word             string         `json:"word"`
	Occurences       int            `json:"occurences"`
	Businesses       map[string]int `json:"businesses,omitempty"`
	SortedBusinesses []ksort.Pair   `json:"sorted_businesses"`
}

type WordList []Word

func (p WordList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p WordList) Len() int           { return len(p) }
func (p WordList) Less(i, j int) bool { return p[i].Occurences > p[j].Occurences }

func sortMapByValue(m map[string]Word) WordList {
	p := make(WordList, len(m))
	i := 0
	for _, v := range m {
		p[i] = v
		i++
	}
	sort.Sort(&p)
	return p
}
