# httpmongo
http方式的增删查改

## 适用类型
###.1.对于小型web
httpmongo使mongodb支持前端页面ajax直接以json形式进行操作

### 1.2风险：如果不对路由进行二次封装，直接引用httpmongo，会使得mongo的增删改查，前端可以直接调用，会丧失对数据库的保护。
###.2.对于大中型web
httpdb把数据库封装成一个高性能http读写接口，可以当作中间件，隐藏在内网，与队列组件结合使用

##httpmongo原理
![img](/exampleimg/httpmongo.jpg)
### &httprequestbody
&httprequestbody表示http.request.body体中的byte[]数据,该数据应该可以被解析为json
应当用&data来表示比较长的json数据，较短的json数据直接放在url中即可

### 以下函数名，与mongo shell保持一致命名

mongo shell官方文档
[https://docs.mongodb.org/manual/reference/method/](https://docs.mongodb.org/manual/reference/method/)
``` shell
/mongo.show dbs	
//无需参数
/mongo/DB.show collections
//无需参数
/mongo/DB/C.count() 	查询集合元素数量
//无需参数
```
``` javascript

/mongo/DB/C.find() 	查询所有文档
var filter={"pid":"123456"}
 $.post("/mongo/DB/C.find(&httprequestbody)",JSON.stringify(filter),function(data,status){
    if(data==null||data=="[]"||data=="")
    {
    	var bds=JSON.parse(data)
    }
})
//当没有参数时，返回所有元素

/mongo/DB/C.findOne()	查询并返回一个对象。如果没有找到则返回 null
/mongo/DB/C.findcount() 	返回匹配该查询的对象总数

这里直接返回数字，而不是jsonstring

/mongo/DB/C.insert()		向聚集中插入对象。不会检查该对象是否已经存在聚集中
var inserter={"pid":"123456"}
 $.post("/mongo/DB/C.insert(&httprequestbody)",JSON.stringify(inserter),function(data,status){
    if(data==null)
    {
    	console.log(data)
    	//如果出错，则不会返回
    	//一般插入成功，会返回{"nInsert":1}
    	
    }
})
/mongo/DB/C.insertmany()		向聚集中批量插入对象。
javascript jquery范例用法
var inserters='{"pid":"123456"},{"pid":"23"},{"name":"lipeng"}'
 $.post("/mongo/DB/C.insertmany(&httprequestbody)",inserters,function(data,status){
    if(data==null)
    {
    	console.log(data)
    	//如果出错，则不会返回
    	//一般插入成功，会返回{"nInsert":3}
    }
})
/mongo/DB/C.remove()    从聚集里删除匹配的对象
javascript jquery范例用法
var filter={"pid":"123456"}
 $.post("/mongo/DB/C.remove(&httprequestbody)",JSON.stringify(filter),function(data,status){
    if(data==null)
    {
    	console.log(data)
    	//如果出错，则不会返回
    	//一般插入成功，会返回{"nRemove":2}
    }
})
/mongo/DB/C.save()  在聚集中保存对象，如果已经存在的话则替换它
save支持1种用法

/mongo/DB/C.update()    在聚集中更新对象。
update应当传入2个json
javascript jquery范例用法
var inserter={"pid":"123456"}
var updater={"$set":{"name":"lipeng"}}
 $.post("/mongo/DB/C.insertmany(&httprequestbody)",JSON.stringfly(inserter)+','+JSON.string(updater),function(data,status){
    if(data==null)
    {
    	console.log(data)
    	//如果出错，则不会返回
    	//一般插入成功，会返回{"nUpdate":2}
    }
})

注意，updater必须填写mongo的操作符哦
```
其余重复命令，暂不支持

### 用法示例
``` go
//
package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/golangframework/httpmongo"
	"github.com/golangframework/moeregexp"
)

var (
	root    = ""
	mongodb = "127.0.0.1:27017"
)

func main() {
	//检查根目录
	root, _ = os.Getwd()
	var mux = http.NewServeMux()
	mux.HandleFunc("/", router)
	err := http.ListenAndServe(":8090", mux)
	log.Println("http.ListenAndServe(:8090)")
	if err != nil {
		log.Fatal("http.ListenAndServe:", err.Error())
	}
}
func router(w http.ResponseWriter, r *http.Request) {
	urlpath := r.URL.Path
  //路由匹配正则 "^/mongo.+"
	if moeregexp.IsMatch(httpmongo.Mongo_path, urlpath) {
	  //调用handler_mongo，处理 /mongo路由下的所有请求
		httpmongo.Httphandler_mongo(mongodb, w, r)
	}
}

```
