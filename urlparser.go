package httpmongo

import (
	"errors"
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
)
const (
	defname               = "[^:?&*/.]+"
	funcs				   = "[A-Za-z0-9_.]+"
	Mongo_path="^/mongo/"
	Mongo_func_path            = "^/mongo/.+"
	Mongo_DB_func_path         = "^/mongo/" + defname + "$"
	Mongo_DB_C_func_path = "^/mongo/" + defname + "/" + defname + "?.+$"
)


func Mongo_parse(urlpath string) (string,error){
	/*	support list
	
	/mongo/show dbs	

	*/
	if !moeregexp.IsMatch(Mongo_func_path,urlpath) {
		err:="urlpath:\n" + urlpath + "不匹配正则" + Mongo_func_path
		panic(err)
		return "",errors.New(err)
	}
	funcs:=urlpath[7:]
	return funcs,nil
}
func Mongo_DB_parse(urlpath string) (string,string,error) {
	/*	support list
	
	/mongo/DB.show collections
	
	*/
	if !moeregexp.IsMatch(Mongo_DB_func_path, urlpath) {
		err:="urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_func_path
		panic(err)
		return "","",errors.New(err)
	}
	DBname := urlpath[7:xstring.Index(urlpath,".")]	
	funcs:=urlpath[xstring.Index(urlpath,"."):]
	return DBname,funcs,nil
}

func Mongo_DB_C_parse(urlpath string) (string, string,string,error) {
	//    localohst:/mongo/wfw.eff/李鹏
	if !moeregexp.IsMatch(Mongo_DB_C_func_path, urlpath) {
		err:="urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_C_func_path
		panic(err)
		return "","","",errors.New(err)
	}
	DB_C := xstring.Split(xstring.Substring(urlpath, 7), "/")
	if len(DB_C) != 2 {
		err:="奇怪的错误"
		panic(err)
		return "","","",errors.New(err)
	}
	DB := DB_C[0]
	C := DB_C[1]
	funcs:=urlpath[xstring.Index(urlpath,"."):]
	return DB, C,funcs,nil
}

