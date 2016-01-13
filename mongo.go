package httpmongo

import (
	"errors"
	"net/http"

	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
)

func dbo_Mongo(w http.ResponseWriter, r *http.Request) error {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_func_path, r.URL.Path) {
		cmd, _ := Mongo_parse(r.URL.Path)
		if cmd == "show dbs" {
			s, err := MgoSession()
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			dbs, _ := s.DatabaseNames()
			w.Write([]byte(xstring.Join(dbs, "\n")))
			return nil
		}
		return errors.New("请求命令不支持")

	} else {
		w.Write([]byte("请求不匹配" + cmd))
		return errors.New("请求不匹配")
	}
}
