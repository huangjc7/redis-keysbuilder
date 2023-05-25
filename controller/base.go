package controller

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"math/rand"
)

var ctx = context.Background()

type redisController struct {
	*redis.ClusterClient
	kCh chan string
}

func NewRedis(db *redis.ClusterClient, kCh chan string) *redisController {
	return &redisController{db, kCh}
}

func (c *redisController) randStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *redisController) ProducerRedisKeys(keynumbers int) {
	for i := 0; i < keynumbers; i++ {
		c.kCh <- c.randStr(8)
	}
	close(c.kCh)
}

func (c *redisController) Run(n int) {
	go c.ProducerRedisKeys(n)
	c.WriteKey()
}

func (c *redisController) RedisPing() error {
	_, err := c.Ping(ctx).Result()
	return err
}

func (c *redisController) WriteKey() {
	i := 0
	for v := range c.kCh {
		err := c.Set(ctx, v, v, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
		i++ //计数器
		log.Printf("已经写入多少 %d keys", i)
	}

}
