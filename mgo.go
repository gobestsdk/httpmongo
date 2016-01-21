package httpmongo

import (
	"time"
	"errors"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	MONGO_URL = ""
)

var session *mgo.Session

func MgoSession() (*mgo.Session, error) {
	if session != nil {
		return session, nil
	} else {
		log.Print("正在尝试新建一个session")
		
		s, err := mgo.DialWithTimeout(MONGO_URL,time.Second*5)
		session = s
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			return nil, errors.New("无法连接mgo服务器:" + MONGO_URL)
		} else {
			log.Print("session建立成功")
			return session, nil
		}
	}
}
