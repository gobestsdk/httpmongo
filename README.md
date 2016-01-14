# httpDB
http方式的增删查改

##适用类型
###.1.对于小型web
httpdb使mongodb支持前端页面ajax直接以json形式进行操作
###.2.对于大中型web
httpdb把数据库封装成一个高性能http读写接口，可以当作中间件，隐藏在内网，与队列组建结合使用

##httpmongo


```javascript

```

### &httprequestbody
&httprequestbody表示http.request.body体中的byte[]数据,该数据应该可以被解析为json
应当用&data来表示比较长的json数据，较短的json数据直接放在url中即可

### 以下函数名，与mongo shell保持一致命名

官方文档
[https://docs.mongodb.org/manual/reference/method/](https://docs.mongodb.org/manual/reference/method/)
``` shell
/mongo.show dbs	

/mongo/DB.show collections

/mongo/DB/C.count() 	查询集合元素数量
/mongo/DB/C.find() 	查询所有文档
/mongo/DB/C.findOne()	查询并返回一个对象。如果没有找到则返回 null
/mongo/DB/C.findcount() 	返回匹配该查询的对象总数
/mongo/DB/C.insert()		向聚集中插入对象。不会检查该对象是否已经存在聚集中
/mongo/DB/C.insertmany()		向聚集中批量插入对象。
/mongo/DB/C.remove()    从聚集里删除匹配的对象
/mongo/DB/C.save()  在聚集中保存对象，如果已经存在的话则替换它
/mongo/DB/C.update()    在聚集中更新对象。update() 有许多参数
```
其余重复命令，暂不支持

