## getTemp

模拟curl命令发送http GET|POST|PUT|DELETE 请求, 这里只实现了GET功能,

 默认是从http://localhost:8080/api/cpu/T 读取数据,配合[pi-db](https://github.com/andy-y-li/pi-db)一起使用。

+ ***编译***

```
$ go build
```



+ ***运行***

  

```
$  ./getTemp -h         
Usage of ./getTemp:
  -method string
    	request method (default "GET|POST|PUT|DELETE")
  -reqBody string
    	json data (default "{a:x}")
  -url string
    	request url (default "http://localhost:8080/api/cpu/T")
    	
$ ./getTemp -method GET 
2020-04-21T13:47:03Z,49.7,63.3%
2020-04-21T13:52:03Z,49.5,23.3%
2020-04-21T13:57:04Z,48.7,63.3%
2020-04-21T14:02:04Z,49.5,70.0%
2020-04-21T14:07:05Z,49.9,25.0%
2020-04-21T14:12:05Z,48  ,71.7%
2020-04-21T14:17:05Z,49.8,56.7%
2020-04-21T14:22:06Z,49  ,26.7%
2020-04-21T14:27:06Z,48.6,88.3%

# 数据可以重定向到一个csv文件中：
$ ./getTemp -method GET > rpi.csv

```



