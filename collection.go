package httpmongo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golangframework/JSON"
	"github.com/golangframework/Object"
	"github.com/golangframework/moeregexp"
	"gopkg.in/mgo.v2/bson"
)

const (
	httpRequestBody = "&httprequestbody"
)

func dbo_Mongo_DB_C(w http.ResponseWriter, r *http.Request) {
	var cmd = ""
	if moeregexp.IsMatch(Mongo_DB_C_func_path, r.URL.Path) {
		DB, C, cmd,err:= Mongo_DB_C_parse(r.URL.Path)
		if err != nil {
				w.Write([]byte(err.Error()))
				return
		}
		funcname := cmd[0:strings.Index(cmd, "(")]
		args := cmd[strings.Index(cmd, "(")+1 : len(cmd)-1]
		if strings.Contains(args, httpRequestBody) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte("没有 接收到 post数据。url不应该有" + httpRequestBody))
				return
			}
			log.Print(string(body))
			args = strings.Replace(args, httpRequestBody, string(body), -1)
		}

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
		case "remove":
			remove(DB, C, args, w)
		case "save":
			save(DB, C, args, w)
		case "update":
			update(DB, C, args, w)
		default:
			w.Write([]byte("请求函数名未知"))
			return
		}
	} else {
		w.Write([]byte("请求不匹配" + cmd))
	}
}
func count(DB string, C string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	count_, _ := c.Count()

	out := Object.Tostring(count_)
	w.Write([]byte(out))
}

func insert(DB string, C string, args string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	var inserter JSON.JSON
	err = json.Unmarshal([]byte(args), &inserter)
	if err != nil {
		w.Write([]byte("无法 序列化为 json"))
		return
	}
	err = c.Insert(inserter)

	if err != nil {
		w.Write([]byte("insert失败"))
		return
	}
	out := "{\"nInsert\":1}"
	w.Write([]byte(out))
}
func insertmany(DB string, C string, args string, w http.ResponseWriter) {
	s, err := MgoSession()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := s.DB(DB).C(C)
	/*args
	[

	{"name":"lipeng"},{"name":"test"}

	]
	*/
	var inserterlist []interface{}
	var ds = "[" + args + "]"
	log.Print(ds)
	err = json.Unmarshal([]byte(ds), &inserterlist)
	if err != nil {
		w.Write([]byte("无法 序列化为 []json"))
		return
	}

	err = c.Insert(inserterlist...)

	if err != nil {
		w.Write([]byte("insertmany失败"))
		return
	}

	out := "{\"nInsert\":" + Object.Tostring(len(inserterlist)) + "}"
	w.Write([]byte(out))
}
func remove(DB string, C string, args string, w http.ResponseWriter) {
	if args == "" {
		out := "{\"nRemoved\":0}"
		w.Write([]byte(out))
	} else {
		s, err := MgoSession()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		c := s.DB(DB).C(C)
		var filter JSON.JSON
		err = json.Unmarshal([]byte(args), &filter)
		if err != nil {
			w.Write([]byte("条件无法 序列化为 json"))
			return
		}
		removeinfo, err := c.RemoveAll(filter)

		if err != nil {
			w.Write([]byte("remove失败"))
			return
		}
		out := "{\"nRemove\":" + Object.Tostring(removeinfo.Removed) + "}"
		w.Write([]byte(out))
	}
}
func save(DB string, C string, args string, w http.ResponseWriter) {
	if args == "" {
		out := "{\"nInsert\":0,\"nUpdate\":0}"
		w.Write([]byte(out))
	} else {
		s, err := MgoSession()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		c := s.DB(DB).C(C)
		var saver bson.M
		err = json.Unmarshal([]byte(args), &saver)
		out := ""
		if err != nil {
			w.Write([]byte("条件无法 序列化为 json"))
			return
		}

		if saver["_id"] == nil {
			err = c.Insert(saver)

			if err != nil {
				w.Write([]byte("插入失败"))
				return
			}
			out = "{\"nInsert\":1}"
		} else {
			if bson.IsObjectIdHex(Object.Tostring(saver["_id"])) == false {

				w.Write([]byte("saver 中 _id 不正确"))
				return
			} else {
				filter := bson.M{"_id": bson.ObjectIdHex(Object.Tostring(saver["_id"]))}

				count, _ := c.Find(filter).Count()
				saver["_id"] = bson.ObjectIdHex(Object.Tostring(saver["_id"]))

				if count >= 1 {
					rinfo, err := c.RemoveAll(filter)
					if err != nil {
						w.Write([]byte("删除失败" + JSON.ToJsonstring(rinfo)))
						return
					}
					err = c.Insert(saver)

					if err != nil {
						w.Write([]byte("插入失败"))
						return
					}
					out = "{\"nUpdate\":1}"
				} else {
					err = c.Insert(saver)

					if err != nil {
						w.Write([]byte("插入失败"))
						return
					}
					out = "{\"nInsert\":1}"
				}
			}
		}

		w.Write([]byte(out))
	}
}
func update(DB string, C string, args string, w http.ResponseWriter) {
	if args == "" {
		out := "{\"nUpdate\":0}"
		w.Write([]byte(out))
	} else {
		/*args
		[

		{"name":"lipeng"},{"name":"test"}

		]
		*/
		var updatearg []JSON.JSON
		err := json.Unmarshal([]byte("["+args+"]"), &updatearg)
		if err != nil || len(updatearg) != 2 {
			w.Write([]byte("条件无法序列化为 2个json"))
			return
		}
		s, err := MgoSession()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		c := s.DB(DB).C(C)
		log.Print(updatearg[0], updatearg[1])
		updateinfo, err := c.UpdateAll(updatearg[0], updatearg[1])

		if err != nil {
			w.Write([]byte("更新失败"))
			return
		}

		out := "{\"nUpdate\":" + Object.Tostring(updateinfo.Updated) + "}"
		w.Write([]byte(out))
	}
}
