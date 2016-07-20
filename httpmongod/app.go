package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/golangframework/httpmongo"
	"github.com/golangframework/moeregexp"
)
var (
	port    = flag.String("port", "8090", "Use -port <port>")
	mongodb = flag.String("mongodb", "127.0.0.1:27017", "Use -mongodb <mongourl>")
)
func main() {
	flag.Parse()
	var mux = http.NewServeMux()
	mux.HandleFunc("/", router)
	log.Println("httpmongod -port " + *port + " -mongodb " + *mongodb + "")
	err := http.ListenAndServe(":"+*port, mux)
	if err != nil {
		log.Fatal("监听端口:", err.Error())
	}
}
func router(w http.ResponseWriter, r *http.Request) {
	urlpath := r.URL.Path
	//路由匹配正则 "^/mongo.+"
	if moeregexp.IsMatch(httpmongo.Mongo_path, urlpath) {
		//调用handler_mongo，处理 /mongo路由下的所有请求
		httpmongo.Httphandler_mongo(*mongodb, w, r)
	}
}