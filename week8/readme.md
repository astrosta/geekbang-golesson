os: linux
cpuInfo: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz*8

# 10字节
## set性能
```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 500000 -t set -d 10 -r 5000000
====== SET ======                                                     
  500000 requests completed in 4.79 seconds
  8 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 104340.57 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.049     0.016     0.047     0.111     0.159     2.191
```
### get性能
```shell
./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 500000 -t get -d 10 -r 5000000
====== GET ======                                                     
500000 requests completed in 4.41 seconds
8 parallel clients
10 bytes payload
keep alive: 1
host configuration "save": 3600 1 300 100 60 10000
host configuration "appendonly": no
multi-thread: no
Summary:
throughput summary: 113352.98 requests per second
latency summary (msec):
avg       min       p50       p95       p99       max
0.043     0.016     0.039     0.063     0.119     0.847
```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:893376
used_memory_human:872.44K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 475927
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:43161840
used_memory_human:41.16M
# 平均每个kv占用内存89字节
```

# 20字节
## set性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 500000 -t set -d 20 -r 2500000
====== SET ======                                                     
  500000 requests completed in 4.59 seconds
  8 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 108956.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.046     0.016     0.047     0.079     0.135     2.135
```
### get性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 500000 -t get -d 20 -r 2500000
====== GET ======                                                     
  500000 requests completed in 4.66 seconds
  8 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 107342.21 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.046     0.016     0.039     0.087     0.159     1.143
```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:893552
used_memory_human:872.61K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 453092
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:44959952
used_memory_human:42.88M
# 平均每个kv占用内存97字节
```

# 50字节
## set性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t set -d 50 -r 5000000
====== SET ======                                                     
  100000 requests completed in 0.94 seconds
  8 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 106382.98 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.047     0.016     0.047     0.079     0.167     1.495
```
### get性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t get -d 50 -r 5000000
====== GET ======                                                     
  100000 requests completed in 0.90 seconds
  8 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 111234.70 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.046     0.016     0.047     0.071     0.135     0.559

```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:893704
used_memory_human:872.76K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 98988
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:13820840
used_memory_human:13.18M
# 平均每个kv占用内存130字节
```


# 200字节
## set性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t set -d 200 -r 5000000
====== SET ======                                                     
  100000 requests completed in 0.91 seconds
  8 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 110253.59 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.045     0.016     0.047     0.071     0.143     1.287

```
### get性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t get -d 200 -r 5000000
====== GET ======                                                     
  100000 requests completed in 0.90 seconds
  8 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 110987.79 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.046     0.016     0.047     0.063     0.119     0.647

```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:893856
used_memory_human:872.91K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 98957
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:30442048
used_memory_human:29.03M
# 平均每个kv占用内存300字节
```

# 2000字节
## set性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t set -d 2000 -r 5000000
====== SET ======                                                     
  100000 requests completed in 0.96 seconds
  8 parallel clients
  2000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 104166.67 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.055     0.024     0.055     0.087     0.159     1.663


```
### get性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t get -d 2000 -r 5000000
====== GET ======                                                     
  100000 requests completed in 0.93 seconds
  8 parallel clients
  2000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 107526.88 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.047     0.016     0.047     0.071     0.159     0.535


```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:894160
used_memory_human:873.20K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 98982
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:210992720
used_memory_human:201.22M
# 平均每个kv占用内存2123字节
```

# 5000字节
## set性能
```shell
[root@localhost src]# ./redis-benchmark -h 127.0.0.1 -p 6379 -c 8 -n 100000 -t set -d 5000 -r 5000000
====== SET ======                                                   
  100000 requests completed in 2.31 seconds
  8 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 43346.34 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.162     0.024     0.151     0.295     0.455    30.543

```
### get性能
```shell
====== GET ======                                                   
  100000 requests completed in 1.25 seconds
  8 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 80128.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.069     0.016     0.047     0.167     0.295    13.191

```
### 内存信息
```shell
# 写入数据之前内存信息
127.0.0.1:6379> info memory
used_memory:894008
used_memory_human:873.05K
# 写入key的数量
127.0.0.1:6379> dbsize
(integer) 98939
# 写入数据之后内存信息
127.0.0.1:6379> info memory
used_memory:514842360
used_memory_human:490.99M
# 平均每个kv占用内存5195字节
```