package main

import (
	"flag"
	"github.com/go-redis/redis/v9"
	"log"
	"redis/controller"
	"sync"
)

var (
	h             string
	k             int
	a             string
	help          bool
	clusterClient *redis.ClusterClient
	wg            sync.WaitGroup
)

func init() {
	flag.StringVar(&a, "a", "", "redis认证密码")
	flag.StringVar(&h, "h", "127.0.0.1:6379", "redis连接地址 ip+port")
	flag.IntVar(&k, "k", 10000, "写入key数量 默认10000")
	flag.BoolVar(&help, "help", false, "查看帮助")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}

	//初始化redis实例
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{h}, Password: a})

	kCh := make(chan string, 100)
	redis := controller.NewRedis(clusterClient, kCh)

	err := redis.RedisPing()
	if err != nil {
		log.Fatal(err)
	}
	redis.Run(k)
}
