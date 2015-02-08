package data

import (
	"labix.org/v2/mgo"
)

// Ensure database indexes are respected for given mongo database
func Index(db *mgo.Database) {
	if err := db.C("reviews").EnsureIndex(mgo.Index{
		Key: []string{"user_id"},
	}); err != nil {
		panic(err)
	}

	if err := db.C("reviews").EnsureIndex(mgo.Index{
		Key: []string{"business_id"},
	}); err != nil {
		panic(err)
	}

	if err := db.C("reviews").EnsureIndex(mgo.Index{
		Key: []string{"user_id", "business_id"},
	}); err != nil {
		panic(err)
	}

	if err := db.C("businesses").EnsureIndex(mgo.Index{
		Key: []string{"$2dsphere:loc"},
	}); err != nil {
		panic(err)
	}
}
