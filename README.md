# bench_mark

1. bench mark 基准测试单元测试验证消耗
2. pprof 火焰图线上应用runtime采集CPU\协程\堆栈等
3. hey 性能压力测试工具获取吞吐、QPS ｜ 最低标准是500并发请求成功率80%+
4. prometheus指标采集获取各个URL接口响应时长占比分布
5. 性能测试监控TP50、TP99、TP999

### 关于应用性能

```
1. 知道自己应用哪个模块哪个接口响应最慢
2. 知道应用随着请求数量增加哪块最薄弱，最容易出现短板
3. 知道自己站点消费请求数量最大值，做出相应的服务降级、动态扩容、限流、熔断等操作
```

### QPS

```
接口性能的最大QPS，是在不断追加并发数量，当程序内部Redis或MySQL等其他依赖出现宕机或者内存爆满的条件下，能够承载的最大每秒钟处理数量

就是这个接口的最大QPS

可以根据这个QPS做一些保护服务的错误比如限流
```

### QPS和并发

```
并发数和QPS不是同一个东西，但是并发数会影响QPS

比如并发10的情况下接口的响应时长 1ms ,那么可以计算 10*1000 = 10000 此时的QPS

但是并发变成了100的时候接口的响应时长变成了 500ms ,此时 100*2 = 200 此时的QPS

并发是10000的时候，接口直接崩溃，何来 QPS


而正常情况下，并发数增高接口响应时间变长，QPS在低并发到高并发的情况下，会先从小到大，到达峰值之后会降低，最后应用直接宕机

这个临界点，也就是应用的最大QPS，并发1000的情况下
```


