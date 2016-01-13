package httpmongo

import (
	"errors"
	"log"
	"net/http"

	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
)

func dbo_Mongo_DB(w http.ResponseWriter, r *http.Request) error {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_DB_func_path, r.URL.Path) {

		DB, cmd, _ := Mongo_DB_parse(r.URL.Path)
		log.Print("匹配:" + DB + "\t" + cmd)
		if cmd == "show collections" {

			Cs, _ := MgoDatabase(DB).CollectionNames()
			w.Write([]byte(xstring.Join(Cs, "\n")))
			return nil
		}
		return errors.New("请求命令不支持")

	} else {
		w.Write([]byte("请求不匹配" + cmd))
		return errors.New("请求不匹配")
	}
}
