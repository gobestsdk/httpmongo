package httpmongo

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/golangframework/moejson"
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
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
		case "count":
			count(DB, C, w)
		case "find":
			find(DB, C, args, w)
		case "findcount":
			findcount(DB, C, args, w)
		case "findOne":
			findOne(DB, C, args, w)
		case "insert":
			insert(DB, C, args, w)
		case "insertmany":
			insertmany(DB, C, args, w)
		default:
			return errors.New("请求函数名未知")
		}
		return errors.New("请求命令不支持")

	} else {
		w.Write([]byte("请求不匹配" + cmd))
		return errors.New("请求不匹配")
	}
}
func count(DB string, C string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)
	count_, _ := c.Count()
	out := xstring.Tostring(count_)
	w.Write([]byte(out))
}
func find(DB string, C string, args string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)
	js := []moejson.Mjson{} //结果集合

	var filter moejson.Mjson
	err := json.Unmarshal([]byte(args), &filter)
	if err != nil {
		panic("无法 序列化为 json")
	}

	err = c.Find(&filter).All(&js)

	jsonlist := moejson.ToJsonarraystring(js) //结果json字符串集合
	out := "[" + xstring.Join(jsonlist, ",") + "]"
	w.Write([]byte(out))
}
func findcount(DB string, C string, args string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)

	var filter moejson.Mjson
	err := json.Unmarshal([]byte(args), &filter)
	if err != nil {
		panic("无法 序列化为 json")
	}

	count_, err := c.Find(&filter).Count()
	out := xstring.Tostring(count_)
	w.Write([]byte(out))
}
func findOne(DB string, C string, args string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)
	js := moejson.Mjson{} //结果

	var filter moejson.Mjson
	err := json.Unmarshal([]byte(args), &filter)
	if err != nil {
		panic("无法 序列化为 json")
	}

	err = c.Find(&filter).One(&js)

	out := js.ToJsonstring()
	w.Write([]byte(out))
}
func insert(DB string, C string, args string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)

	var inserter moejson.Mjson
	err := json.Unmarshal([]byte(args), &inserter)
	if err != nil {
		panic("无法序列化为 json")
	}
	err = c.Insert(inserter)
	if err != nil {
		panic("插入失败")
	}
	out := "resultn:1"
	w.Write([]byte(out))
}
func insertmany(DB string, C string, args string, w http.ResponseWriter) {
	c := MgoDataCollect(DB, C)
	/*args
	{"list":[

	{"name":"lipeng"},{"name":"test"}

	]}
	*/
	var header = "{\"list\":["
	var footer = "]}"
	var inserterlist moejson.Mjson
	var ds = header + args + footer
	log.Print(ds)
	err := json.Unmarshal([]byte(ds), &inserterlist)
	if err != nil {
		panic("无法 序列化为 []json")
	}
	log.Print(inserterlist)
	inserter := [2]moejson.Mjson{moejson.Mjson{"name": "m1"}, moejson.Mjson{"name": "m2"}}
	err = c.Insert(inserter)
	if err != nil {
		panic("插入失败")
	}
	out := "resultn:" + xstring.Tostring(2)
	w.Write([]byte(out))
}
