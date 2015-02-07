package data

import (
	"labix.org/v2/mgo"
)

type Model interface {
	TableName() string
	GetId() interface{}
}

func Save(db *mgo.Database, m Model) error {
	var err error

	//if m.GetId() == "" {
	err = db.C(m.TableName()).Insert(m)
	//} else {
	//	_, err = db.C(m.TableName()).UpsertId(m.GetId(), m)
	//}

	return err
}
