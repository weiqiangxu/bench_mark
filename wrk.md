# WRK

```
wrk -c 400 -t 10 -d 120s -H "Authorization: Bearer pwd" --latency http://localhost:8080/user/list
```

### 用法

```
连接数（c）与 QPS（q），请求响应时间毫秒（t）的关系大概可理解为：q = 1000/t * c
如果连接数是10，响应时间100毫秒，那么QPS就是 1000/100 * 10 = 100 (10个连接数，每个连接数每秒钟处理10个当然就是100啦)

wrk 4.2.0

使用方法: wrk <选项> <被测HTTP服务的URL>                            
  Options:                                            
    -c, --connections <N>  跟服务器建立并保持的TCP连接数量|连接数（connection）可以理解为并发数一般在测试过程中，这个值需要使用者不断向上调试，直至 QPS 达到一个临界点，便可认为此时的并发数为系统所能承受的最大并发量
    -d, --duration    <T>  压测时间           
    -t, --threads     <N>  使用多少个线程进行压测   一般是CPU核数，最大不要超过 CPUx2 核数，否则会带来额外的上下文切换将线程数设置为CPU核数主要是为了WRK 能最大化利用 CPU，使结果更准确
                                                      
    -s, --script      <S>  指定Lua脚本路径       
    -H, --header      <H>  为每一个HTTP请求添加HTTP头      
        --latency          在压测结束后，打印延迟统计信息   
        --timeout     <T>  超时时间     
    -v, --version          打印正在使用的wrk的详细版本信息
                                                      
  <N>代表数字参数，支持国际单位 (1k, 1M, 1G)
  <T>代表时间参数，支持时间单位 (2s, 2m, 2h)

```

### 结果解析

```
Running 30s test @ www.bing.com （压测时间30s）
  8 threads and 200 connections （共8个测试线程，200个连接）
  Thread Stats   Avg      Stdev     Max   +/- Stdev
              （平均值） （标准差）（最大值）（正负一个标准差所占比例）
    Latency    46.67ms  215.38ms   1.67s    95.59%
    （延迟）
    Req/Sec     7.91k     1.15k   10.26k    70.77%
    （处理中的请求数）
  Latency Distribution （延迟分布）
     50%    2.93ms
     75%    3.78ms
     90%    4.73ms
     99%    1.35s （99分位的延迟）
  1790465 requests in 30.01s, 684.08MB read （30.01秒内共处理完成了1790465个请求，读取了684.08MB数据）
Requests/sec:  59658.29 （平均每秒处理完成59658.29个请求）
Transfer/sec:     22.79MB （平均每秒读取数据22.79MB）
```

### 测试本机程序
```
wrk -c 400 -t 10 -d 120 （此时redis pool是10）

Running 2m test @ http://localhost:8080/user/list
  10 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   177.92ms  173.57ms   1.92s    95.73%
    Req/Sec   165.03    194.85     1.59k    89.84%
  Latency Distribution
     50%  200.67ms
     75%  213.85ms
     90%  239.46ms
     99%    1.07s 
  177071 requests in 2.00m, 29.24MB read
  Socket errors: connect 159, read 245, write 0, timeout 0
Requests/sec:   1474.36
Transfer/sec:    249.33KB
```

```
wrk -c 400 -t 10 -d 120 （此时redis pool是1000）

Running 2m test @ http://localhost:8080/user/list
  10 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   179.75ms  187.92ms   1.82s    95.75%
    Req/Sec   164.29    175.63     1.45k    90.18%
  Latency Distribution
     50%  198.88ms
     75%  207.98ms
     90%  224.15ms
     99%    1.23s 
  180130 requests in 2.00m, 29.79MB read
  Socket errors: connect 159, read 104, write 0, timeout 0
Requests/sec:   1499.81
Transfer/sec:    253.98KB
```


### 压测MySQL查询

```
wrk -c 400 -t 10 -d 120s -H "Authorization: Bearer pwd" --latency http://localhost:8080/school/list

Running 2m test @ http://localhost:8080/school/list
  10 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   601.26ms  289.60ms   1.96s    71.23%
    Req/Sec    46.36     45.06   370.00     88.90%
  Latency Distribution
     50%  568.25ms
     75%  759.34ms
     90%    1.02s 
     99%    1.33s 
  47928 requests in 2.00m, 4.31MB read
  Socket errors: connect 159, read 107, write 0, timeout 1
  Non-2xx or 3xx responses: 47782
Requests/sec:    399.11
Transfer/sec:     36.73KB


优化连接以后



Running 2m test @ http://localhost:8080/school/list
  10 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   189.92ms  300.86ms   1.82s    86.78%
    Req/Sec   214.11    144.93     1.06k    67.94%
  Latency Distribution
     50%    8.54ms
     75%  312.77ms
     90%  607.06ms
     99%    1.26s 
  227579 requests in 2.00m, 37.32MB read
  Socket errors: connect 159, read 101, write 0, timeout 525
Requests/sec:   1895.17
Transfer/sec:    318.28KB
```

### 高并发下

```
wrk -c 1000 -t 10 -d 120s -H "Authorization: Bearer pwd" --latency http://localhost:8080/school/list

Running 2m test @ http://localhost:8080/school/list
  10 threads and 1000 connections
^[[C  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   125.15ms  287.81ms   2.00s    85.88%
    Req/Sec   190.15    139.06     1.09k    67.21%
  Latency Distribution
     50%    3.75ms
     75%   13.21ms
     90%  502.79ms
     99%    1.47s 
  210312 requests in 2.00m, 34.50MB read
  Socket errors: connect 759, read 95, write 0, timeout 991
Requests/sec:   1751.24
Transfer/sec:    294.15KB
```

```
### 启动 4 个线程, 每个线程维持 3000/4 个TCP连接 （-c可以理解为并发数）
wrk -c 3000 -t 4 -d 120s -H "Authorization: Bearer pwd" --latency http://localhost:8080/school/list

Running 2m test @ http://localhost:8080/school/list
  4 threads and 3000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   131.99ms  276.01ms   1.98s    87.10%
    Req/Sec   474.62    237.05     1.29k    67.47%
  Latency Distribution
     50%    4.81ms
     75%   41.81ms
     90%  466.00ms
     99%    1.34s 
  201385 requests in 2.00m, 33.03MB read
  Socket errors: connect 2753, read 157, write 31, timeout 801
Requests/sec:   1677.28
Transfer/sec:    281.69KB
```

[性能常识 WRK 个人使用经验总结](https://testerhome.com/topics/22601)

```
Req/Sec：表示的是每个线程每秒的完成的请求数，顺序分别是： 平均值，标准差，最大值，正负标准差；
```
