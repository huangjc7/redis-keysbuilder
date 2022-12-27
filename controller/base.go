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
}

func RandStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewRedis(db *redis.ClusterClient) *redisController {
	return &redisController{db}
}

func (c *redisController) RedisPing() error {
	_, err := c.Ping(ctx).Result()
	return err
}

func (c *redisController) WriteKey() {
	err := c.Set(ctx, RandStr(5), RandStr(7), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}
