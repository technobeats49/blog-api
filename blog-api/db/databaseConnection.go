package db

import (
	"gopkg.in/mgo.v2"
)

const (
	MONGODB_URL = "mongodb://localhost:27017"
	DB_NAME     = "blogs-go"
	COLL_NAME   = "blogs"
)

func getSession() *mgo.Session {

	s, err := mgo.Dial(MONGODB_URL)
	if err != nil {
		panic(err)
	}
	return s
}

// db session object
var DBS *mgo.Session = getSession()
