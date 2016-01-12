package httpmongo

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	MONGO_URL = "127.0.0.1:27017"
)

var database *mgo.Database
var mgoSession *mgo.Session
var collection *mgo.Collection
var err error

func MgoSession() *mgo.Session {
	if mgoSession == nil {
		mgoSession, err = mgo.Dial(MONGO_URL)
		if err != nil {
			log.Println(err)
		}
	}
	mgoSession.SetMode(mgo.Monotonic, true)
	return mgoSession.Clone()
}

//MgoDatabase 切换数据库
func MgoDatabase(name string) *mgo.Database {
	session := MgoSession()
	if database == nil {
		database = session.DB(name)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return database
}

//MgoDataCollect 同时切换数据库和集合
func MgoDataCollect(data, c string) *mgo.Collection {
	s := MgoSession()
	if collection == nil {
		collection = s.DB(data).C(c)
	}
	return collection
}
func CloseSession() {
	mgoSession.Close()
}
