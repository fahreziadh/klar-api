package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

type DAO struct {
	Server   string
	Database string
}

func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
