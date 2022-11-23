# hey

```
Usage: hey [options...] <url>
Options:
  // 指定运行的总请求数。默认值为200。
  -n  Number of requests to run. Default is 200.
  // 客户端并发执行的请求数，默认为50。总请求数不能小于并发数。
  -c  Number of workers to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  // 客户端发送请求的速度限制，以每秒响应数QPS为单位，默认没有限制。
  -q  Rate limit, in queries per second (QPS) per worker. Default is no rate limit.
  // 发送请求的持续时长，超时后程序停止并退出。若指定了持续时间，则忽略总请求数(-n)，例如-z 10s，-z 3m
  -z  Duration of application to send requests. When duration is reached,
      application stops and exits. If duration is specified, n is ignored.
      Examples: -z 10s -z 3m.
  // 输出类型。若没有提供，则打印摘要。CSV是唯一支持的格式，结果以逗号分隔各个指标项。
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.
  // HTTP方法，例如GET，POST，PUT，DELETE，HEAD，OPTIONS方法
  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  // HTTP请求头，可以用-H连续添加多个请求头。
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  // 每个请求的超时时间（以秒为单位）。默认值为20s，数值0代表永不超时。
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  // HTTP响应头
  -A  HTTP Accept header.
  // HTTP请求正文
  -d  HTTP request body.
  // HTTP请求正文文件
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  // 内容类型，默认为“ text / html”。
  -T  Content-type, defaults to "text/html".
  // 基本身份验证，用户名：密码。
  -a  Basic authentication, username:password.
  // HTTP代理地址作, 使用host:port格式。
  -x  HTTP Proxy address as host:port.
  // 启用HTTP / 2
  -h2 Enable HTTP/2.
  // HTTP主机头
  -host HTTP Host header.
  // 禁用压缩。
  -disable-compression  Disable compression.
  // 禁用保持活动状态，防止重新使用不同的HTTP请求之间的TCP连接。
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  // 禁用HTTP重定向                     
  -disable-redirects    Disable following of HTTP redirects
  // 使用的cpu内核数。当前计算机的默认值为8核。
  -cpus                 Number of used cpu cores.
                        (default for current machine is 8 cores)
```

```
# 并发400 持续120
hey -z 120s -c 400 http://localhost:8080/school/list


Summary:
  Total:	122.0965 secs
  Slowest:	18.8614 secs
  Fastest:	0.0005 secs
  Average:	0.2101 secs
  Requests/sec:	1685.4458
  
  Total data:	10038549 bytes
  Size/request:	48 bytes

Response time histogram:
  0.001 [1]	|
  1.887 [202711]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  3.773 [1460]	|
  5.659 [406]	|
  7.545 [211]	|
  9.431 [164]	|
  11.317 [22]	|
  13.203 [0]	|
  15.089 [23]	|
  16.975 [62]	|
  18.861 [21]	|


Latency distribution:
  10% in 0.0018 secs
  25% in 0.0026 secs
  50% in 0.0050 secs
  75% in 0.1091 secs
  90% in 0.5956 secs
  95% in 1.1385 secs
  99% in 2.3687 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0005 secs, 18.8614 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0209 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0213 secs
  resp wait:	0.2099 secs, 0.0005 secs, 18.8613 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0498 secs

Status code distribution:
  [200]	205081 responses

Error distribution:
  [252]	Get "http://localhost:8080/school/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
  [1]	Get "http://localhost:8080/school/list": read tcp [::1]:58551->[::1]:8080: read: connection reset by peer
```


```
# 总数2000 并发50
hey -n 2000 -c 50  http://localhost:8080/school/list

### 输出

Summary:
  Total:	1.0323 secs
  Slowest:	0.3833 secs
  Fastest:	0.0007 secs
  Average:	0.0220 secs
  Requests/sec:	1937.3976
  
  Total data:	98000 bytes
  Size/request:	49 bytes

Response time histogram:
  0.001 [1]	|
  0.039 [1670]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.077 [81]	|■■
  0.115 [148]	|■■■■
  0.154 [43]	|■
  0.192 [13]	|
  0.230 [18]	|
  0.268 [19]	|
  0.307 [4]	|
  0.345 [1]	|
  0.383 [2]	|


Latency distribution:
  10% in 0.0014 secs
  25% in 0.0019 secs
  50% in 0.0030 secs
  75% in 0.0071 secs
  90% in 0.0914 secs
  95% in 0.1157 secs
  99% in 0.2475 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.0007 secs, 0.3833 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.0033 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:	0.0218 secs, 0.0006 secs, 0.3778 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	2000 responses
```


```
# 并发数1000持续120s
Documents$  hey -z 120s -c 1000  http://localhost:8080/school/list

Summary:
  Total:	139.0173 secs
  Slowest:	19.4089 secs
  Fastest:	0.0007 secs
  Average:	0.2934 secs
  Requests/sec:	1089.4541  [QPS是1089]
  
  Total data:	7159272 bytes
  Size/request:	48 bytes

Response time histogram:
  0.001 [1]	|
  1.941 [141946]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  3.882 [2203]	|■
  5.823 [881]	|
  7.764 [320]	|
  9.705 [401]	|
  11.646 [93]	|
  13.586 [0]	|
  15.527 [28]	|
  17.468 [211]	|
  19.409 [30]	|


Latency distribution:
  10% in 0.0016 secs
  25% in 0.0023 secs
  50% in 0.0041 secs
  75% in 0.0189 secs
  90% in 0.6920 secs
  95% in 1.4974 secs
  99% in 4.6482 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0003 secs, 0.0007 secs, 19.4089 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.4354 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0319 secs
  resp wait:	0.2928 secs, 0.0006 secs, 19.4088 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0245 secs

Status code distribution:
  [200]	146114 responses

Error distribution:
  [4013]	Get "http://localhost:8080/school/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

### 看看Redis的QPS
```
# 并发1000的情况下持续120s
hey -z 120s -c 1000  http://localhost:8080/user/list 

Summary:
  Total:	120.0890 secs
  Slowest:	19.9609 secs
  Fastest:	0.0000 secs
  Average:	0.1286 secs
  Requests/sec:	7764.5443 [QPS是7764]
  
  Total data:	49388119 bytes
  Size/request:	53 bytes

Response time histogram:
  0.000 [1]	|
  1.996 [930296]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  3.992 [1]	|
  5.988 [2]	|
  7.984 [3]	|
  9.980 [0]	|
  11.977 [0]	|
  13.973 [0]	|
  15.969 [2]	|
  17.965 [1]	|
  19.961 [2]	|


Latency distribution:
  10% in 0.0333 secs
  25% in 0.0843 secs
  50% in 0.1072 secs
  75% in 0.1208 secs
  90% in 0.2757 secs
  95% in 0.3815 secs
  99% in 0.5086 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 19.9609 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.1190 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0884 secs
  resp wait:	0.1284 secs, 0.0000 secs, 19.9609 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0747 secs

Status code distribution:
  [200]	930308 responses

Error distribution:
  [15]	Get "http://localhost:8080/user/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
  [1]	Get "http://localhost:8080/user/list": read tcp [::1]:52074->[::1]:8080: read: connection reset by peer
```

### QPS会因为并发数变化

```
# 一万的并发下
hey -z 20s -c 10000  http://localhost:8080/user/list 

Summary:
  Total:	40.0010 secs
  Slowest:	19.9858 secs
  Fastest:	0.1706 secs
  Average:	11.6929 secs
  Requests/sec:	47539.5388 [QPS 47539]
  
  Total data:	440880 bytes
  Size/request:	57 bytes

Response time histogram:
  0.171 [1]	|
  2.152 [580]	|■■■■■■■■■■■■■■■■■■
  4.134 [183]	|■■■■■■
  6.115 [371]	|■■■■■■■■■■■■
  8.097 [476]	|■■■■■■■■■■■■■■■
  10.078 [928]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  12.060 [1165]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  14.041 [1267]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  16.023 [963]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  18.004 [1136]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  19.986 [626]	|■■■■■■■■■■■■■■■■■■■■


Latency distribution:
  10% in 4.3353 secs
  25% in 8.6970 secs
  50% in 12.1812 secs
  75% in 15.9408 secs
  90% in 17.7296 secs
  95% in 18.7592 secs
  99% in 19.5151 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0416 secs, 0.1706 secs, 19.9858 secs
  DNS-lookup:	0.0764 secs, 0.0000 secs, 0.5305 secs
  req write:	0.0013 secs, 0.0000 secs, 0.4417 secs
  resp wait:	5.7211 secs, 0.0835 secs, 18.8410 secs
  resp read:	0.0041 secs, 0.0000 secs, 0.6897 secs

Status code distribution:
  [200]	7696 responses

Error distribution:
  [2037]	Get "http://localhost:8080/user/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

### 看看MySQL的QPS

```
### 1000的并发 [还很正常]
hey -z 120s -c 1000  http://localhost:8080/school/list 

Summary:
  Total:	138.9990 secs
  Slowest:	19.9999 secs
  Fastest:	0.0007 secs
  Average:	0.4394 secs
  Requests/sec:	1261.2181
  
  Total data:	8402677 bytes
  Size/request:	48 bytes

Response time histogram:
  0.001 [1]	|
  2.001 [167147]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  4.001 [1739]	|
  6.000 [822]	|
  8.000 [1172]	|
  10.000 [382]	|
  12.000 [111]	|
  14.000 [129]	|
  16.000 [141]	|
  18.000 [39]	|
  20.000 [199]	|


Latency distribution:
  10% in 0.0022 secs
  25% in 0.0036 secs
  50% in 0.0098 secs
  75% in 0.3932 secs
  90% in 1.3295 secs
  95% in 1.5090 secs
  99% in 7.2367 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0004 secs, 0.0007 secs, 19.9999 secs
  DNS-lookup:	0.0002 secs, 0.0000 secs, 0.3546 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0726 secs
  resp wait:	0.4388 secs, 0.0003 secs, 19.9999 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0396 secs

Status code distribution:
  [200]	171882 responses

Error distribution:
  [2439]	Get "http://localhost:8080/school/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

```
### 5000的并发 [MySQL直接炸了]
hey -z 120s -c 5000  http://localhost:8080/school/list 
```

```
### 2000的并发 QPS就到达了1500了
hey -z 120s -c 2000  http://localhost:8080/school/list

Summary:
  Total:	139.6770 secs
  Slowest:	20.6940 secs
  Fastest:	0.0001 secs
  Average:	0.5845 secs
  Requests/sec:	1513.0905  [之前的1000并发QPS是1200现在并发达到2000就变成了1500 但是响应时长从 2.00变2.06]
  
  Total data:	18433105 bytes
  Size/request:	91 bytes

Response time histogram:
  0.000 [1]	|
  2.069 [186852]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  4.139 [8793]	|■■
  6.208 [1347]	|
  8.278 [1827]	|
  10.347 [945]	|
  12.416 [471]	|
  14.486 [291]	|
  16.555 [274]	|
  18.625 [235]	|
  20.694 [147]	|


Latency distribution:
  10% in 0.0030 secs
  25% in 0.0077 secs
  50% in 0.0338 secs
  75% in 0.3293 secs
  90% in 1.6945 secs
  95% in 2.6071 secs
  99% in 8.9285 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0009 secs, 0.0001 secs, 20.6940 secs
  DNS-lookup:	0.0004 secs, 0.0000 secs, 1.0833 secs
  req write:	0.0001 secs, 0.0000 secs, 0.1896 secs
  resp wait:	0.5792 secs, 0.0001 secs, 20.5821 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.1527 secs

Status code distribution:
  [200]	201183 responses

Error distribution:
  [6653]	Get "http://localhost:8080/school/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

### 看看3000的并发下 很多请求失败

```
hey -z 120s -c 3000  http://localhost:8080/school/list

Summary:
  Total:	134.8539 secs
  Slowest:	22.9676 secs
  Fastest:	0.0277 secs
  Average:	2.4938 secs
  Requests/sec:	919.3209 [QPS降低到了919]
  
  Total data:	13965726 bytes
  Size/request:	130 bytes

Response time histogram:
  0.028 [1]	|
  2.322 [76778]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  4.616 [13518]	|■■■■■■■
  6.910 [5932]	|■■■
  9.204 [3324]	|■■
  11.498 [2871]	|■
  13.792 [1538]	|■
  16.086 [1440]	|■
  18.380 [1006]	|■
  20.674 [523]	|
  22.968 [1]	|


Latency distribution:
  10% in 0.2889 secs
  25% in 0.4177 secs
  50% in 1.1410 secs
  75% in 2.7342 secs
  90% in 6.9127 secs
  95% in 10.8025 secs
  99% in 16.9526 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0059 secs, 0.0277 secs, 22.9676 secs
  DNS-lookup:	0.0129 secs, 0.0000 secs, 2.8672 secs
  req write:	0.0007 secs, 0.0000 secs, 0.9943 secs
  resp wait:	0.9801 secs, 0.0001 secs, 19.7405 secs
  resp read:	0.0022 secs, 0.0000 secs, 3.7197 secs

Status code distribution:
  [200]	106932 responses

Error distribution:
  [8]	Get "http://localhost:8080/school/list": EOF
  [4606]	Get "http://localhost:8080/school/list": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

