package data

import (
	"errors"

	"labix.org/v2/mgo"
)

type Model struct {
	Id string
}

func (m *Model) TableName() string {
	panic(errors.New("TableName() not implemented on model"))
}

func (m *Model) Save(db *mgo.Database) error {
	var err error

	if m.Id == "" {
		err = db.C(m.TableName()).Insert(m)
	} else {
		_, err = db.C(m.TableName()).UpsertId(m.Id, m)
	}

	return err
}
