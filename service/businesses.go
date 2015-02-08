package main

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiasaki/yelp-dataset-api/data"
	"labix.org/v2/mgo/bson"
)

func handleMontrealReviewWords(c *gin.Context) {
	businessIds := []string{
		"cOuf4T33pbwejU6BxqtVzQ",
		"cOuf4T33pbwejU6BxqtVzQ",
		"GSwUJUgJ-0cnhiO4lsuSOA",
		"ipIeFooftvKsbQq67YX9kw",
		"nbiyY-01-DyUkZ8c9fVNrA",
		"pG8nRNktGCqq2M-apvL9Ig",
		"9U7xp1MGuOVOBzw3N9AOVA",
		"qbBhrTb6XVAcRNM__i3SPw",
		"pA659_V23xpRPYnQdh0vDA",
		"IFoZEWCfCfWPAv_2eZiFrg",
		"anYjyT9mLZthbJxXk-spfw",
		"a6Y7l1GRJxlX7_Yavnh8bA",
		"KEopXGnfgw3ZxtsUmAfYrA",
		"di4J7RmeNSmBN6qGj9rqng",
		"QlsCTIQcGKm6m9BQzgy30w",
		"j9mALdpaWQwFno4FSsevIQ",
		"HtLC-ETA8mTYUWDboDiJTQ",
		"yhcR9XFnNtAmZo5lHfK7Eg",
		"EqrOAyElnB4vn5sX4GAksA",
		"DJI-nhBjADCP0eahvVbZKA",
		"U-pOhk942gmfTO6veMBfHA",
		"VaIofXGibAa2ibEomErtVA",
		"flYvcbsLvI5CoL9iFgPInw",
		"XPNJt0zir6OmimaSk83tCw",
		"G28z7q4EF6JaBEyAQisHrA",
		"Eq09MP-yeL7qkGwJgnPA0Q",
		"5s2e_PCD5qshfBAFoLKqLg",
		"Ds9cKpbfPiXb11tjrMDvLw",
		"-bgbw5LcKJ97kumwljc8kw",
		"zca-aUYe9Q_yEIXeNCW7Ww",
		"7yosOEWC4Q9-yx81lQxH-Q",
		"xeysfxV3Y9rWe_smQQBwnA",
		"0Z3dCshAfz2GvRwUVkMOSQ",
		"acmbXFQfH-Fnue2wq77nJw",
		"GYqknb_itWF-ZzYZqwQC4A",
		"z0qdeUIZhZLG9BCBFV_I1A",
		"BWaCdx0LgVmIGNkwiCBMEg",
		"TSqnLFfqvHvAzrh4qtUBwA",
		"C9zCQ04M7HddK_UvHPF5zA",
		"KxampCnWLJOgUmdLsSzVQA",
		"HSEBGQeKZQbvLpmGgq1QyQ",
		"Sx96d65M32CG6rNX5LHWcQ",
		"iG5AOfpR0-e3-bdZLpKDMg",
		"H-q09bPKnXaBW0NGuhb09A",
		"kClmBhnrxUqXQHQ9TOIuaQ",
		"6FegfipLbt_w9BrqLpgj6A",
		"VRGJoI77kh07XhPR5ccKng",
		"JWepmB8IGhNfSM93_SmmtQ",
		"4nx8W1xZHIbcyQb5UJSiiw",
		"wMBy-fn3jkZ1VabtikmQDQ",
		"_Tdi18bgcFY-QmaiwOSiLg",
		"5LhcUep64k4OyawKzXk4SA",
		"ViKTmf7xQ7RCqhHs7HW5aQ",
		"mRvfPqj-Irp78JFRPup3wQ",
		"ZVFC5WlNcTzWKk59r34Scg",
		"IFfx9sxgcBuWoe3pAcbLvQ",
		"sRtMupSMALp5Ov4JQ4OE1A",
		"4Xk9Ve1EkN6REtOma4CNqQ",
		"1Lt_ZeJPZOINqLmW-muCgw",
		"VA_vQ6aA0WmXY-bUbYXbuw",
		"i381TLhzlhilOuX0jE4-Og",
		"QQz_NrM7bcrcyPDMw7Fgpw",
		"y0nP7sQUufEDMV0lARp8Iw",
		"efUQKzgsSe3ttZJDLRzKMQ",
		"qz7ybMMUQAvFmMMPInePuA",
		"o-Y5Ois5awDNJHPNBVtp5w",
		"qazVQsiFSaiLlgO_6fAFrA",
		"uin8gE-DiRkf6ZdBb_XamQ",
		"MpMzyvRQfi9Syo_kvnMyWg",
		"fVawkEVn8uZr35qWzeldjQ",
		"GALenlN_Xg4M9W8NjwSC-g",
		"lDPAiRLOqcF0zr_HzsMVSA",
		"5qF7hy3DkZhHR-41599gQQ",
		"L_tDRiPXR0HwTcveZ_5RIw",
		"d7yEPVrczwKe-Un9tSywbA",
		"WbJ6PsiZuFBPCGxrKeCHrw",
		"Fa2bzZPBFi-seY_GIGgldg",
		"FHAEXGKsY4HEK1A9mx52eg",
		"yQEfxrdi_hoFo71aprvqBQ",
		"_uVa_LnA_TU4nfbaM9-6bQ",
		"2gL-Mhq2zpesKfgEilD5BA",
		"SVzMkTYIEjK9gSGnOBVvyw",
		"zb01gN65HaDO1TfRHqVoIg",
		"euBHu4Rr8w67mYxS5B9ivQ",
		"RgJqstm9YmwVZ2wOYoWwEw",
		"Ft4hmzhrA2fPDaAa1855Iw",
		"6KmRNO-wq3yGQOZogToJxQ",
		"AIskp-s-bN3lM-ytMG9R4w",
		"Sm8ErbRr1fLDyrD4HYVQCA",
		"6Iwfpp_4lyGx7-JudsaeBg",
		"2NtSJQvwB4e7RK0Zm931Rw",
		"27E6qzMutpc9ZtIlKR70ig",
		"nW0WnA957mjFz3_i77To0A",
		"OEwDk8-qpDlA-5mDwWFXoQ",
		"fbOVOFBGcbxi0OklbCe2rQ",
		"IDkuywWnYNPHtw5ZuGgMPg",
		"O5sdmXByyoXYTFODjn7dAw",
		"gv8H080nC_9wtUf9QN9-Vg",
		"yRJDlRxE9UQRHVLzAzaPcg",
		"Fat5ff3edpyWl8Q8_udQWA",
		"-tiT0PFDzzjkKIXOemmOpw",
		"GSwUJUgJ-0cnhiO4lsuSOA",
		"ipIeFooftvKsbQq67YX9kw",
		"nbiyY-01-DyUkZ8c9fVNrA",
		"pG8nRNktGCqq2M-apvL9Ig",
		"9U7xp1MGuOVOBzw3N9AOVA",
		"qbBhrTb6XVAcRNM__i3SPw",
		"pA659_V23xpRPYnQdh0vDA",
		"IFoZEWCfCfWPAv_2eZiFrg",
		"anYjyT9mLZthbJxXk-spfw",
		"a6Y7l1GRJxlX7_Yavnh8bA",
		"KEopXGnfgw3ZxtsUmAfYrA",
		"di4J7RmeNSmBN6qGj9rqng",
		"QlsCTIQcGKm6m9BQzgy30w",
		"j9mALdpaWQwFno4FSsevIQ",
		"HtLC-ETA8mTYUWDboDiJTQ",
		"yhcR9XFnNtAmZo5lHfK7Eg",
		"EqrOAyElnB4vn5sX4GAksA",
		"DJI-nhBjADCP0eahvVbZKA",
		"U-pOhk942gmfTO6veMBfHA",
		"VaIofXGibAa2ibEomErtVA",
		"flYvcbsLvI5CoL9iFgPInw",
		"XPNJt0zir6OmimaSk83tCw",
		"G28z7q4EF6JaBEyAQisHrA",
		"Eq09MP-yeL7qkGwJgnPA0Q",
		"5s2e_PCD5qshfBAFoLKqLg",
		"Ds9cKpbfPiXb11tjrMDvLw",
		"-bgbw5LcKJ97kumwljc8kw",
		"zca-aUYe9Q_yEIXeNCW7Ww",
		"7yosOEWC4Q9-yx81lQxH-Q",
		"xeysfxV3Y9rWe_smQQBwnA",
		"0Z3dCshAfz2GvRwUVkMOSQ",
		"acmbXFQfH-Fnue2wq77nJw",
		"GYqknb_itWF-ZzYZqwQC4A",
		"z0qdeUIZhZLG9BCBFV_I1A",
		"BWaCdx0LgVmIGNkwiCBMEg",
		"TSqnLFfqvHvAzrh4qtUBwA",
	}
	reviews := []data.YelpReview{}

	err := dbSession.DB("").C("reviews").Find(bson.M{
		"business_id": bson.M{"$in": businessIds},
	}).All(&reviews)

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

	c.JSON(200, gin.H{"words": sortMapByValue(words)})
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
