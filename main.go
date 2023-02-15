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
	p             string
	help          bool
	clusterClient *redis.ClusterClient
	wg            sync.WaitGroup
)

func init() {
	flag.StringVar(&p, "p", "", "redis认证密码")
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
		Password: p,
	})

	redis := controller.NewRedis(clusterClient)

	err := redis.RedisPing()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			redis.WriteKey()
			fmt.Printf("已经写入%v key\n", i)
			defer wg.Done()
		}()
		wg.Wait()
	}
}
