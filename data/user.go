package data

type YelpUser struct {
	Id           string  `json:"user_id" bson:"_id,omitempty"`
	Name         string  `json:"name" bson:"name"`
	ReviewCount  int     `json:"review_count" bson:"review_count"`
	YelpingSince string  `json:"yelping_since" bson:"yelping_since"`
	Fans         int     `json:"fans" bson:"fans"`
	AverageStars float32 `json:"average_stars" bson:"average_stars"`
	Elite        []int   `json:"elite" bson:"elite"`
}

func (u YelpUser) TableName() string {
	return "users"
}

func (u YelpUser) GetId() interface{} {
	return u.Id
}
