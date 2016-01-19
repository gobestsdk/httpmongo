package httpmongo

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	MONGO_URL = ""
)

var session *mgo.Session

func MgoSession() (*mgo.Session, error) {
	log.Print("当前session")
	log.Print(session)
	if session != nil {
		return session, nil
	} else {
		log.Print("session需要重建")
		s, err := mgo.Dial(MONGO_URL)
		session = s
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			return nil, errors.New("无法连接mgo服务器:" + MONGO_URL)
		} else {
			return session, nil
		}
	}
}
