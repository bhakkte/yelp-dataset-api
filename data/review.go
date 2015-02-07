package data

type YelpReview struct {
	Id         string         `json:"review_id" bson:"_id,omitempty"`
	BusinessId string         `json:"business_id" bson:"business_id"`
	UserId     string         `json:"user_id" bson:"user_id"`
	Date       string         `json:"date" bson:"date"`
	Text       string         `json:"text" bson:"text"`
	Stars      float32        `json:"stars" bson:"stars"`
	Votes      map[string]int `json:"votes" bson:"votes"`
}

func (u YelpReview) TableName() string {
	return "reviews"
}

func (u YelpReview) GetId() interface{} {
	return u.Id
}
