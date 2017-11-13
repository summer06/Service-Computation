# Simple Web Server
---
##### Requirement: implement a simple web server

##### Develop Tool: Atom 1.21.0

----

### net/http package

http 是典型的 C/S 架构，客户端向服务端发送请求（request），服务端做出应答（response）。

golang 的标准库 net/http 提供了 http 编程有关的接口，封装了内部TCP连接和报文解析的复杂琐碎的细节，使用者只需要和 http.request 和 http.ResponseWriter 两个对象交互就行。也就是说，我们只要写一个 handler，请求会通过参数传递进来，而它要做的就是根据请求的数据做处理，把结果写到 Response 中。

```
package main

import (
	"io"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":12345", nil)
}
```


上面的代码没有什么问题，但是有一个不便：每次写 Handler 的时候，都要定义一个类型，然后编写对应的 ServeHTTP 方法，这个步骤对于所有 Handler 都是一样的。重复的工作总是可以抽象出来，net/http 也正这么做了，它提供了 http.HandleFunc 方法，允许直接把特定类型的函数作为 handler。
```
package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":12345", nil)
}
```
##### **参考上面的做法，我对我的代码进行了优化，通过ab测试，效果确实好了不少。**
可以说，golang自带的net／http包的确给我们提供了很多便利。

### Usage
在命令行输入
`$ ./main -p 8000`
服务器即在8000端口运行。
端口可以自己使用-p参数来进行指定，默认值为8000.

### Result
- 使用curl进行测试

`$ curl -v http://localhost:8000`

```
  $ curl -v http://localhost:8000

  * Rebuilt URL to: http://localhost:8000/
  *   Trying ::1...
  * TCP_NODELAY set
  * Connected to localhost (::1) port 8000 (#0)
  > GET / HTTP/1.1
  > Host: localhost:8000
  > User-Agent: curl/7.54.0
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Date: Mon, 13 Nov 2017 16:51:43 GMT
  < Content-Length: 11
  < Content-Type: text/plain; charset=utf-8
  <
  * Connection #0 to host localhost left intact
```

- 使用ab测试

`$ ab -n 1000 -c 100 http://127.0.0.1:8000/`

```
  $ ab -n 1000 -c 100 http://127.0.0.1:8000/

  This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
  Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
  Licensed to The Apache Software Foundation, http://www.apache.org/

  Benchmarking 127.0.0.1 (be patient)
  Completed 100 requests
  Completed 200 requests
  Completed 300 requests
  Completed 400 requests
  Completed 500 requests
  Completed 600 requests
  Completed 700 requests
  Completed 800 requests
  Completed 900 requests
  Completed 1000 requests
  Finished 1000 requests


  Server Software:        
  Server Hostname:        127.0.0.1
  Server Port:            8000

  Document Path:          /
  Document Length:        11 bytes

  Concurrency Level:      100
  Time taken for tests:   0.059 seconds
  Complete requests:      1000
  Failed requests:        0
  Total transferred:      128000 bytes
  HTML transferred:       11000 bytes
  Requests per second:    16808.70 [#/sec] (mean)
  Time per request:       5.949 [ms] (mean)
  Time per request:       0.059 [ms] (mean, across all concurrent requests)
  Transfer rate:          2101.09 [Kbytes/sec] received

  Connection Times (ms)
                min  mean[+/-sd] median   max
  Connect:        1    3   0.5      3       4
  Processing:     2    3   0.5      3       5
  Waiting:        1    3   0.5      3       4
  Total:          3    6   0.6      6       8

  Percentage of the requests served within a certain time (ms)
    50%      6
    66%      6
    75%      6
    80%      6
    90%      6
    95%      7
    98%      7
    99%      7
   100%      8 (longest request)
```
