package httpmongo

import (
	"errors"
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

func MgoSession() (*mgo.Session, error) {
	if mgoSession == nil {
		mgoSession, err = mgo.Dial(MONGO_URL)
		if err != nil {
			log.Println("数据库连接失败")
			return nil, errors.New("数据库连接失败")
		}
	}
	mgoSession.SetMode(mgo.Monotonic, true)
	return mgoSession.Clone(), nil
}

//MgoDatabase 切换数据库
func MgoDatabase(DB string) *mgo.Database {
	session, _ := MgoSession()
	if database == nil {
		database = session.DB(DB)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return database
}

//MgoDataCollect 同时切换数据库和集合
func MgoDataCollect(DB, c string) *mgo.Collection {
	session, _ := MgoSession()
	if collection == nil {
		collection = session.DB(DB).C(c)
	}
	return collection
}
func CloseSession() {
	mgoSession.Close()
}
