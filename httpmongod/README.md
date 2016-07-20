# httpmongod

### 启动命令

>httpmongod -port :8090 -mongodb 127.0.0.1:27017

### 为什么要做httpmongo这个工具

http方式的mongod服务,提供增删改查的仿mongoshell命令接口

启动mongod服务，再启动httpmongod

然后在你的浏览器地址栏试试输入
``` 
http://localhost:8090/mongo.show dbs

```
这样就可以查看到本机mongo所有的数据库

![](https://github.com/golangdeveloper/httpmongod/blob/master/img/example.png?raw=true)

httpmongo把shell的交互，变换到了http交互，你可以便捷的使用chrome的postman工具来进行数据库读写

## 引用库来源

[https://github.com/golangframework/httpmongo](https://github.com/golangframework/httpmongo)

## benchmark

     
使用ab插入100万条数据
```
timeloveboy@timeloveboy-os:~$ cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c 
     12  Intel(R) Xeon(R) CPU           X5650  @ 2.67GHz

  
timeloveboy@timeloveboy-os:~$ ab -k -n 1000000 -c 100 'http://localhost:8090/mongo/MYDB/myTABLE.insert({"key":"timeloveboy"})'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100000 requests
Completed 200000 requests
Completed 300000 requests
Completed 400000 requests
Completed 500000 requests
Completed 600000 requests
Completed 700000 requests
Completed 800000 requests
Completed 900000 requests
Completed 1000000 requests
Finished 1000000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8090

Document Path:          /mongo/MYDB/myTABLE.insert({"key":"timeloveboy"})
Document Length:        13 bytes

Concurrency Level:      100
Time taken for tests:   196.809 seconds
Complete requests:      1000000
Failed requests:        0
Keep-Alive requests:    1000000
Total transferred:      154000000 bytes
HTML transferred:       13000000 bytes
Requests per second:    5081.07 [#/sec] (mean)
Time per request:       19.681 [ms] (mean)
Time per request:       0.197 [ms] (mean, across all concurrent requests)
Transfer rate:          764.15 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       5
Processing:     0   20   8.8     19     177
Waiting:        0   20   8.8     19     177
Total:          0   20   8.8     19     177

Percentage of the requests served within a certain time (ms)
  50%     19
  66%     22
  75%     24
  80%     25
  90%     30
  95%     34
  98%     41
  99%     49
 100%    177 (longest request)

```
然后打开连接
```
http://localhost:8090/mongo/MYDB/myTABLE.findcount({"key":"timeloveboy"})

返回1000000，说明100万条数据全部插入成功
```


10万条
Requests per second:    5151.70 [#/sec] (mean)

1万条
Requests per second:    8938.03 [#/sec] (mean)

1000条
Requests per second:    8392.29 [#/sec] (mean)
