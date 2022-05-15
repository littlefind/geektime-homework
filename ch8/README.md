# 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 10k 20k 50k 100k字节 value 大小，redis get set 性能。
```
shell> redis-benchmark -c 100 -n 10000 -q -d 10 -t get,set
SET: 81300.81 requests per second, p50=0.607 msec       
GET: 84745.77 requests per second, p50=0.551 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 20 -t get,set
SET: 88495.58 requests per second, p50=0.559 msec       
GET: 85470.09 requests per second, p50=0.575 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 50 -t get,set
SET: 86956.52 requests per second, p50=0.543 msec       
GET: 89285.71 requests per second, p50=0.551 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 200 -t get,set
SET: 86956.52 requests per second, p50=0.535 msec       
GET: 84745.77 requests per second, p50=0.543 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 1024 -t get,set
SET: 92592.59 requests per second, p50=0.535 msec       
GET: 86956.52 requests per second, p50=0.535 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 5120 -t get,set
SET: 83333.34 requests per second, p50=0.575 msec       
GET: 83333.34 requests per second, p50=0.855 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 10240 -t get,set
SET: 75187.97 requests per second, p50=0.727 msec       
GET: 74074.07 requests per second, p50=0.647 msec                   

shell> redis-benchmark -c 100 -n 10000 -q -d 20480 -t get,set
SET: 61728.39 requests per second, p50=0.831 msec       
GET: 55555.55 requests per second, p50=0.879 msec                  

shell> redis-benchmark -c 100 -n 10000 -q -d 51200 -t get,set
SET: 28490.03 requests per second, p50=1.671 msec                   
GET: 10050.25 requests per second, p50=7.343 msec            

shell> redis-benchmark -c 100 -n 10000 -q -d 102400 -t get,set
SET: 24752.47 requests per second, p50=2.031 msec                   
GET: 8976.66 requests per second, p50=10.519 msec
```
可以发现在5k以后get/set性能出现了大幅下滑

# 作业2：写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
## 0
```
used_memory:1071360
used_memory_human:1.02M
```

## 1W
```
used_memory:1842432
used_memory_human:1.76M
```
(1842432-1071360) / 10000 = 77(byte)

## 5W
```
used_memory:4795840
used_memory_human:4.57M
```
(4795840-1071360) / 50000 = 74(byte)

## 10W
```
used_memory:8520320
used_memory_human:8.13M
```
(8520320-1071360) / 100000 = 74(byte)

## 20W
```
used_memory:15969088
used_memory_human:15.23M
```
(15969088-1071360) / 200000 = 74(byte)

## 30W
```
used_memory:24466432
used_memory_human:23.33M
```
(24466432-1071360) / 300000 = 77(byte)

## 40W
```
used_memory:30866624
used_memory_human:29.44M
```
(30866624-1071360) / 400000 = 74(byte)

## 50W
```
used_memory:37266816
used_memory_human:35.54M
```
(37266816-1071360) / 500000 = 72(byte)

## 100W
```
used_memory:73469504
used_memory_human:70.07M
```
(73469504-1071360) / 1000000 = 72(byte)
