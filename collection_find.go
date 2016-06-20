package httpmongo

import (
	"encoding/json"

	"net/http"
	"strings"

	"github.com/golangframework/JSON"
	"github.com/golangframework/Object"
)


func find(DB string, C string, args string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	js := []JSON.JSON{} //结果集合

	var filter JSON.JSON
	if args == "" {
		args = "{}"
	}
	err = json.Unmarshal([]byte(args), &filter)
	if err != nil {
		w.Write([]byte("find 条件，无法 序列化为 json"))
		return
	}
	err = c.Find(&filter).All(&js)
	if err != nil {
		w.Write([]byte("find失败"))
		return
	}
	jsonlist := JSON.ToJsonstringArray(js) //结果json字符串集合
	out := "[" + strings.Join(jsonlist, ",") + "]"
	w.Write([]byte(out))
}
func findcount(DB string, C string, args string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	var filter JSON.JSON
	err = json.Unmarshal([]byte(args), &filter)
	if err != nil {
		w.Write([]byte("无法 序列化为 json"))
		return
	}

	count_, err := c.Find(&filter).Count()
	if err != nil {
		w.Write([]byte("find.count失败"))
		return
	}
	out := Object.Tostring(count_)
	w.Write([]byte(out))
}
func findOne(DB string, C string, args string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	js := JSON.JSON{} //结果

	var filter JSON.JSON
	err = json.Unmarshal([]byte(args), &filter)
	if err != nil {
		w.Write([]byte("无法 序列化为 json"))
		return
	}

	err = c.Find(&filter).One(&js)
	if err != nil {
		w.Write([]byte("findOne失败"))
		return
	}
	out := js.ToJsonstring()
	w.Write([]byte(out))
}