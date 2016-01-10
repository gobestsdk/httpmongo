package httpDB

import(
	"gopkg.in/mgo.v2"
	"encoding/json"
)
func connectDB(dbaddr string) mgo.Session{
	//以下部分代码抄写自官网教程
	session, err := mgo.Dial(dbaddr)
    if err != nil {
            panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.单调模式
    session.SetMode(mgo.Monotonic, true)
	return *session
}

func insert_jsonstring(DB string,C string,jsonstr string) error {	
    return insert_jsonbytes(DB,C,[]byte(jsonstr))
}

func insert_jsonbytes(DB string,C string,jsonbytes []byte) error {
	session:=connectDB("")
    c := session.DB(DB).C(C)
	
	var document interface{}
	err := json.Unmarshal(jsonbytes, &document)
	if err != nil {
	    return err
	}
	err = c.Insert(document)
	if err != nil {
	    return err
	}
	return nil
}