package data

type YelpUser struct {
	Id           string  `json:"user_id"`
	Name         string  `json:"name"`
	ReviewCount  int     `json:"review_count"`
	YelpingSince string  `json:"yelping_since"`
	Fans         int     `json:"fans"`
	AverageStars float32 `json:"average_stars"`
	Elite        []int   `json:"elite"`
}
