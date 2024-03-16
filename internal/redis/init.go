package redis

import (
	"chatbot/internal/config"
	"chatbot/internal/config/global"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

type RedisConfigMap map[string]RedisConfigInstance

type RedisConfigInstance struct {
	Addr         string        `yaml:"Addr"`
	Password     string        `yaml:"Password"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	PoolSize     int           `yaml:"PoolSize"`
	MinIdleConns int           `yaml:"MinIdleConns"`
	PoolTimeout  time.Duration `yaml:"PoolTimeout"`
	IdleTimeout  time.Duration `yaml:"IdleTimeout"`
}

var instances = make(map[string]*RedisClient)

func init() {
	config.LoadConfig()
	var redisConfigMap RedisConfigMap
	cfg := global.Cfg
	cfg.UnmarshalKey("redis", &redisConfigMap)

	for k, v := range redisConfigMap {

		instances[k] = &RedisClient{
			Client: redis.NewClient(&redis.Options{
				Addr:         v.Addr,
				Password:     "",
				DialTimeout:  v.DialTimeout * time.Second,
				ReadTimeout:  v.ReadTimeout * time.Second,
				WriteTimeout: v.WriteTimeout * time.Second,
				PoolSize:     v.PoolSize,
				MinIdleConns: v.MinIdleConns,
				PoolTimeout:  v.PoolTimeout * time.Second,
				IdleTimeout:  v.IdleTimeout * time.Minute,
			}),
		}
	}
}

func GetRedisClient(key string) *RedisClient {
	return instances[key]
}
