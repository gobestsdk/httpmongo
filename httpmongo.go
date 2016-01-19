package httpmongo

import (
	"net/http"

	"github.com/golangframework/moeregexp"
)

func Httphandler_mongo(DBaddr string, w http.ResponseWriter, r *http.Request) {
	MONGO_URL = DBaddr

	checkrequest(w, r)
	if moeregexp.IsMatch(Mongo_DB_C_func_path, r.URL.Path) {
		dbo_Mongo_DB_C(w, r)
	} else {
		if moeregexp.IsMatch(Mongo_DB_func_path, r.URL.Path) {
			dbo_Mongo_DB(w, r)
		} else {
			if moeregexp.IsMatch(Mongo_func_path, r.URL.Path) {
				dbo_Mongo(w, r)
			}
		}
	}

}

func checkrequest(w http.ResponseWriter, r *http.Request) bool {
	if moeregexp.IsMatch(Mongo_path, r.URL.Path) {
		return true
	}
	panic("非法请求")
	return false
}
