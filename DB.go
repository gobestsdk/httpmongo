package httpmongo

import (
	"log"
	"net/http"
	"strings"

	"github.com/golangframework/moeregexp"
)

func dbo_Mongo_DB(w http.ResponseWriter, r *http.Request) {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_DB_func_path, r.URL.Path) {
		DB, cmd,err:= Mongo_DB_parse(r.URL.Path)
		if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		log.Print("匹配:" + DB + "\t" + cmd)
		if cmd == "show collections" {
			s, err := MgoSession()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			Cs, _ := s.DB(DB).CollectionNames()
			w.Write([]byte(strings.Join(Cs, "\n")))
			return
		}
		w.Write([]byte("请求命令不支持"))
		return
	} else {
		w.Write([]byte("请求不匹配" + cmd))
	}
}
