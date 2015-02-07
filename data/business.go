package data

type YelpBusiness struct {
	Id          string                 `json:"business_id" bson:"_id,omitempty"`
	Name        string                 `json:"name" bson:"name"`
	FullAddress string                 `json:"full_address" bson:"full_address"`
	Open        bool                   `json:"open" bson:"open"`
	Categories  []string               `json:"categories" bson:"categories"`
	ReviewCount int                    `json:"review_count" bson:"review_count"`
	Latitude    float32                `json:"latitude" bson:"latitude"`
	Longitude   float32                `json:"longitude" bson:"longitude"`
	Stars       float32                `json:"stars" bson:"stars"`
	Attributes  map[string]interface{} `json:"attributes" bson:"attributes"`
}

func (u YelpBusiness) TableName() string {
	return "businesses"
}

func (u YelpBusiness) GetId() interface{} {
	return u.Id
}
