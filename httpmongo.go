package httpmongo

import (
	"log"
	"net/http"
)

const (
	defname               = "[^:?&*/]+"
	Mongo_path            = "^/mongo/.+"
	Mongo_DB_path         = "^/mongo/" + defname + "$"
	Mongo_DB_C_path       = "^/mongo/" + defname + "/" + defname + "$"
	Mongo_DB_C_Query_path = "^/mongo/" + defname + "/" + defname + "?.+$"
)

type JsonDocument map[interface{}]interface{}

func Httphandler_mongo(DBaddr string, w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("func") {
	case "PUT":
		log.Printf("PUT")
		PUT(w, r)
	case "POST":
		log.Print("POST")

	case "DELETE":
		log.Print("DELETE")

	case "PATCH":
		log.Print("PATCH")

	case "GET":
		fallthrough
	default:
		log.Printf("GET ")
		w.Write([]byte(" "))
	}
}
