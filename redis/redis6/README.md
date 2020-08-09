redis集群部署
----------------






进入容器，创建集群

```
redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1
```

```
--cluster-replicas 1 ：表示一个主节点对应一个从节点
```

执行完生成：
```
[OK] All 16384 slots covered
```

查看集群状态

```
root@docker-desktop:/data# redis-cli --cluster check 127.0.0.1:7000
127.0.0.1:7000 (2aec3f50...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (505e9309...) -> 0 keys | 5462 slots | 1 slaves.
127.0.0.1:7005 (5c4f6af7...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 2aec3f50d0eee66c103b0ab79bc42040cbf97e20 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: 505e93093a49be38f013b6e8be5aca82ab1fda5c 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: f0a0726baa6059f3ac062b9cc9e073beb5835b3d 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 2aec3f50d0eee66c103b0ab79bc42040cbf97e20
S: cc0b4b2428e02069b48d49a4391c097d6576a00e 127.0.0.1:7004
   slots: (0 slots) slave
   replicates 505e93093a49be38f013b6e8be5aca82ab1fda5c
S: b96782e84db5e8edd5caf183d278385ff0b412c5 127.0.0.1:7002
   slots: (0 slots) slave
   replicates 5c4f6af737da739be3711cafdad1dc7be3657edb
M: 5c4f6af737da739be3711cafdad1dc7be3657edb 127.0.0.1:7005
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.

```

关闭一个7002节点，可发现，7005自动变为master

```
docker stop redis6_redis-7002_1

root@docker-desktop:/data# redis-cli --cluster check 127.0.0.1:7000
127.0.0.1:7000 (2aec3f50...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (505e9309...) -> 0 keys | 5462 slots | 1 slaves.
127.0.0.1:7005 (5c4f6af7...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 2aec3f50d0eee66c103b0ab79bc42040cbf97e20 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: 505e93093a49be38f013b6e8be5aca82ab1fda5c 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: f0a0726baa6059f3ac062b9cc9e073beb5835b3d 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 2aec3f50d0eee66c103b0ab79bc42040cbf97e20
S: cc0b4b2428e02069b48d49a4391c097d6576a00e 127.0.0.1:7004
   slots: (0 slots) slave
   replicates 505e93093a49be38f013b6e8be5aca82ab1fda5c
S: b96782e84db5e8edd5caf183d278385ff0b412c5 127.0.0.1:7002
   slots: (0 slots) slave
   replicates 5c4f6af737da739be3711cafdad1dc7be3657edb
**M: 5c4f6af737da739be3711cafdad1dc7be3657edb 127.0.0.1:7005
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)**
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.

```


新增一个节点


参考：
https://redis.io/topics/cluster-tutorial
