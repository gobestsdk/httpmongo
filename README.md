# httpDB
http方式的增删查改

##适用类型
###.1.对于小型web来说，httpdb使mongodb支持前端页面ajax直接以json形式进行操作
###.2.对于大中型web来说，httpdb把数据库封装成一个高性能http读写接口，可以当作中间件，隐藏在内网，与队列组建结合使用

##httpmongo

httpmongo借鉴了coachdb的rest api

路由设计如下：
localhost://DBname/Collectionname

###insert
http.method:post

``` javascript
ajax.

```


###query
http.method:get

###delete
http.method:delete

###update
http.method:update
