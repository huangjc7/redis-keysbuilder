package main

import (
	"flag"
	"github.com/go-redis/redis/v9"
	"log"
	"redis/controller"
	"sync"
	"time"
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

func CreateRedisKeys() {
	//for i := 0; i < n; i++ {
	//	wg.Add(1)
	//	go func() {
	//		redis.WriteKey()
	//		fmt.Printf("已经写入%v key\n", i)
	//		defer wg.Done()
	//	}()
	//	wg.Wait()
	//}
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

	kCh := make(chan string, 1000)
	//vCh := make(chan map[string]string, 1000)
	redis := controller.NewRedis(clusterClient, kCh)

	err := redis.RedisPing()
	if err != nil {
		log.Fatal(err)
	}
	go redis.ProducerRedisKeys(n)
	go redis.WriteKey()
	time.Sleep(time.Second * 25)

	// TODO
	// 没有控制写入key的数量
	// 没有做详细对比数据（使用channel快多少）
}
