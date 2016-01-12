package httpmongo

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golangframework/moeregexp"
)

func PUT(w http.ResponseWriter, r *http.Request) {
	if moeregexp.IsMatch(Mongo_DB_C_path, r.URL.Path) == false {
		log.Print(Mongo_DB_C_path)
		log.Print(r.URL.Path)
		panic("路径格式不正确,应该为:/DB/C")
	}
	DB, C := Mongo_DB_C_parse(r.URL.Path)
	bytes, _ := ioutil.ReadAll(r.Body)
	Data := string(bytes)

	w.Write([]byte(DB + "\t" + C))
}
