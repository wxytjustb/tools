package main

import (
  "github.com/chasex/redis-go-cluster"
  "fmt"
  "time"
  )

func main()  {
    cluster, err := redis.NewCluster(&redis.Options{
        StartNodes: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
        ConnTimeout: 50 * time.Millisecond,
        ReadTimeout: 50 * time.Millisecond,
        WriteTimeout: 50 * time.Millisecond,
        KeepAlive: 16,
        AliveTime: 60 * time.Second,
    })
  fmt.Println(cluster)
#  fmt.Println(err)
}
