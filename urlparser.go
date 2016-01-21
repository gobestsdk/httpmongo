package httpmongo

import (
	"errors"
	"strings"

	"github.com/golangframework/moeregexp"
)

const (
	defname              = "[^:?&*/.]+"
	funcs                = "[A-Za-z0-9_.]+"
	Mongo_path           = "^/mongo.+"
	Mongo_func_path      = "^/mongo..+"
	Mongo_DB_func_path   = "^/mongo/" + defname + "..+"
	Mongo_DB_C_func_path = "^/mongo/" + defname + "/" + defname + "..+"
)

func Mongo_parse(urlpath string) (string, error) {
	/*	support list
	
		/mongo.show dbs

	*/
	if !moeregexp.IsMatch(Mongo_func_path, urlpath) {
		err := "urlpath:\n" + urlpath + "不匹配正则" + Mongo_func_path
		return "", errors.New(err)
	}
	funcs := urlpath[7:]
	return funcs, nil
}
func Mongo_DB_parse(urlpath string) (string, string, error) {
	/*	support list

		/mongo/DB.show collections

	*/
	if !moeregexp.IsMatch(Mongo_DB_func_path, urlpath) {
		err := "urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_func_path
		return "", "", errors.New(err)
	}
	DBname := urlpath[7:strings.Index(urlpath, ".")]
	funcs := urlpath[strings.Index(urlpath, ".")+1:]
	return DBname, funcs, nil
}

func Mongo_DB_C_parse(urlpath string) (string, string, string, error) {
	/*

		/mongo/DB/C.count()     查询集合元素数量
		/mongo/DB/C.find()  查询所有文档
		/mongo/DB/C.findOne()   查询并返回一个对象。如果没有找到则返回 null
		/mongo/DB/C.find().count()  返回匹配该查询的对象总数
		/mongo/DB/C.insert()        向聚集中插入对象。不会检查该对象是否已经存在聚集中
		/mongo/DB/C.remove()    从聚集里删除匹配的对象
		/mongo/DB/C.save()  在聚集中保存对象，如果已经存在的话则替换它
		/mongo/DB/C.update()    在聚集中更新对象。update() 有许多参数

	*/
	if !moeregexp.IsMatch(Mongo_DB_C_func_path, urlpath) {
		err := "urlpath:\n" + urlpath + "不匹配正则" + Mongo_DB_C_func_path
		return "", "", "", errors.New(err)
	}
	DB_C := strings.Split(urlpath[7:strings.Index(urlpath, ".")], "/")
	if len(DB_C) != 2 {
		err := "奇怪的错误"
		return "", "", "", errors.New(err)
	}
	DB := DB_C[0]
	C := DB_C[1]
	funcs := urlpath[strings.Index(urlpath, ".")+1:]
	return DB, C, funcs, nil
}
