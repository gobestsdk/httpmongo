package httpmongo

import (
	"github.com/golangframework/moeregexp"
	"github.com/golangframework/xstring"
)

func Mongo_DB_parse(urlpath string) string {
	//    localohst:/mongo/wfw.eff
	if !moeregexp.IsMatch(Mongo_DB_path, urlpath) {
		panic("urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_path)
	}
	DBname := xstring.Substring(urlpath, 7)
	return DBname
}

func Mongo_DB_C_parse(urlpath string) (string, string) {
	//    localohst:/mongo/wfw.eff/李鹏
	if !moeregexp.IsMatch(Mongo_DB_C_path, urlpath) {
		panic("urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_C_path)
	}
	DB_C := xstring.Split(xstring.Substring(urlpath, 7), "/")
	if len(DB_C) != 2 {
		panic("奇怪的错误，切分 DB和C会出错？")
	}
	DB := DB_C[0]
	C := DB_C[1]
	return DB, C
}

func Mongo_DB_C_Query_parse(urlpath string) (string, string, string) {
	//   get localohst:/mongo/wfw.eff/user?d\{"name":"lipeng"}
	if !moeregexp.IsMatch(Mongo_DB_C_Query_path, urlpath) {
		panic("urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_C_Query_path)
	}
	DB_C_str := urlpath[7:xstring.Index(urlpath, "?")]
	if len(DB_C_str) != 2 {
		panic("奇怪的错误，切分 DB和C会出错？")
	}
	DB_C := xstring.Split(DB_C_str, "\\")
	DB := DB_C[0]
	C := DB_C[1]
	Query := xstring.Substring(urlpath, xstring.Index(urlpath, "?")+1)
	return DB, C, Query
}
