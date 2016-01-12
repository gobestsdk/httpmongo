# httpDB
http方式的增删查改

##适用类型
###.1.对于小型web来说，httpdb使mongodb支持前端页面ajax直接以json形式进行操作
###.2.对于大中型web来说，httpdb把数据库封装成一个高性能http读写接口，可以当作中间件，隐藏在内网，与队列组建结合使用

##httpmongo

httpmongo借鉴了coachdb的rest api

路由设计如下：
localhost://DBname/Collectionname

###header新增字段
"dbcmd":""
```javascript

 <script type="text/javascript" language="javascript">
        $(function() {

            $("#test").click(function() {
                $.ajax({
                    type: "GET",
                    url: "/mongo/DBname/Collectionname",
                    beforeSend: function(request) {
                        request.setRequestHeader("dbcmd", "find");
                    },
                    success: function(result) {
                        alert(result);
                    }
                });
            });
        });
    </script>

```

### get方法支持以下dbcmd:
#### Session.DatabaseNames
获取当前回话中的所有数据库名

#### DB.CollectionNames
获取数据库中的所有集合名

#### C.Count
获取集合中的元素数目
这也是默认方法

#### C.Find.All
查询该集合所有数据

### post方法支持以下dbcmd

#### C.Insert
一次可以插入多个对象

#### C.Find.One
查询单条数据

#### C.Find
查询符合条件的所有数据

#### C.Remove
删除符合条件的1个数据

#### C.RemoveAll
删除符合条件的所有数据

#### C.Update
更新符合条件的1个数据

#### C.UpdateAll
更新符合条件的所有数据

#### C.Upsert
更新插入符合条件的1个数据
