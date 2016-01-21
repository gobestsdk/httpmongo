package httpmongo

import (
	"net/http"
	"strings"

	"github.com/golangframework/moeregexp"
)

func dbo_Mongo(w http.ResponseWriter, r *http.Request) {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_func_path, r.URL.Path) {
		cmd, err:= Mongo_parse(r.URL.Path)
		if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		if cmd == "show dbs" {
			s, err := MgoSession()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			dbs, _ := s.DatabaseNames()
			w.Write([]byte(strings.Join(dbs, "\n")))
			return
		}
		w.Write([]byte("请求命令不支持"))
		return
	} else {
		w.Write([]byte("请求不匹配" + cmd))
	}
}
