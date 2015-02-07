package data

import (
	"labix.org/v2/mgo"
)

// Ensure database indexes are respected for given mongo database
func Index(db *mgo.Database) {
	return
	//if err := db.C("users").EnsureIndex(mgo.Index{
	//	Key:    []string{"id"},
	//	Unique: true,
	//}); err != nil {
	//	panic(err)
	//}
}
