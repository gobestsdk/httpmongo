package httpmongo

import (
	"errors"
	"log"
	"net/http"

	"github.com/golangframework/moejson"
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
	"gopkg.in/mgo.v2/bson"
)

func dbo_Mongo_DB_C(w http.ResponseWriter, r *http.Request) error {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_DB_C_func_path, r.URL.Path) {

		DB, C, cmd, _ := Mongo_DB_C_parse(r.URL.Path)
		log.Print("dbo_Mongo_DB_C匹配:" + DB + "\t" + C + "\t" + cmd)

		funcname := cmd[0:xstring.Index(cmd, "(")]
		args := cmd[xstring.Index(cmd, "(")+1 : len(cmd)-1]

		log.Print("funcname:" + funcname + "\t" + args)
		switch funcname {
		case "find":
			find(DB, C, args, w)
		default:
			return errors.New("请求函数名未知")
		}
		return errors.New("请求命令不支持")

	} else {
		w.Write([]byte("请求不匹配" + cmd))
		return errors.New("请求不匹配")
	}
}

func find(DB string, C string, args string, w http.ResponseWriter) {
	if args == "" {
		c := MgoDataCollect(DB, C)
		js := []moejson.Mjson{}
		err = c.Find(&bson.M{}).All(&js)
		jsonlist := moejson.ToJsonarraystring(js)
		out := "[" + xstring.Join(jsonlist, ",") + "]"
		w.Write([]byte(out))
	}
}
