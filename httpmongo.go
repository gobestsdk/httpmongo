package httpmongo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golangframework/moeregexp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var DBaddr = ""

const (
	Mongourlpath = "^/mongo/[^:?&*/]+/[^:?&*/]+?.+$"
	DBorC_name   = "/[^:?&*/]+"
)

type JsonDocument map[interface{}]interface{}

func Httphandler_mongo_put(dbaddr string, w http.ResponseWriter, r *http.Request) {
	DBaddr = dbaddr
	urlpath := r.URL.Path
	if !moeregexp.IsMatch(Mongourlpath, urlpath) {
		panic("urlpath:\n" + urlpath + " \nnot match\n" + Mongourlpath)
	}

	DBname := moeregexp.GetMatchCollection(DBorC_name, urlpath)[0]
	DBname = DBname[1:len(DBname)]
	Cname := moeregexp.GetMatchCollection(DBorC_name, urlpath)[1]
	Cname = Cname[1:len(Cname)]
	var Urlquerydocument bson.D
	prestr := 1 + 5 + 1 + len(DBname) + 1 + len(Cname) + 1
	if len(urlpath) > prestr {
		Urlquerydocument = bson.Unmarshal([]byte(urlpath[prestr:]), Urlquerydocument)
		log.Print(Urlquerydocument)
	}
	out := ""
	switch r.Method {
	case "GET":
		{
			log.Printf("GET ")

			for j := range find_filterstring(DBname, Cname, "{}") {
				bs, _ := json.Marshal(j)
				out += string(bs)
			}
		}
	}
	w.Write([]byte(out))
}
func connectDB(dbaddr string) mgo.Session {
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

func insert_jsonstring(DB string, C string, jsonstr string) error {
	return insert_jsonbytes(DB, C, []byte(jsonstr))
}

func insert_jsonbytes(DB string, C string, jsonbytes []byte) error {
	session := connectDB(DBaddr)
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

func find_filterstring(DB string, C string, filterstr string) []JsonDocument {
	session := connectDB(DBaddr)
	c := session.DB(DB).C(C)
	var filter interface{}
	err := json.Unmarshal([]byte(filterstr), &filter)
	if err != nil {
		panic(err)
	}
	log.Print(filter)
	query := c.Find(nil)

	result := []JsonDocument{}

	query.All(&result)
	return result
}
