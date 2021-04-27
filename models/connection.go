package models

import (
	"github.com/globalsign/mgo"
)

func GetDB() (error, *mgo.Database) {

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return err, nil
	}
	d := session.DB("Tokenized")
	return nil, d
}
