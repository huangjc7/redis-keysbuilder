package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"redis/controller"
	"sync"
)

var (
	h             string
	n             int
	help          bool
	clusterClient *redis.ClusterClient
	wg            sync.WaitGroup
)

func init() {
	flag.StringVar(&h, "h", "127.0.0.1:6379", "redis连接地址 ip+port")
	flag.IntVar(&n, "n", 10000, "写入key数量 默认10000")
	flag.BoolVar(&help, "help", false, "查看帮助")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}

	//初始化redis实例

	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			h,
		},
	})

	redis := controller.NewRedis(clusterClient)

	err := redis.RedisPing()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		go func() {
			wg.Add(1)
			redis.WriteKey()
			fmt.Printf("已经写入%v key\n", i)
			wg.Done()
		}()
		wg.Wait()
	}
}
