package redis

import (
	"context"
	"fmt"
	"time"
)

var ctx = context.Background()

func SetHashField(biz, key, field, value string) error {
	redisClient := GetRedisClient(biz).Client
	return redisClient.HSet(ctx, key, field, value).Err()
}

func GetHashField(biz, key, field string) (string, error) {
	redisClient := GetRedisClient(biz).Client
	return redisClient.HGet(ctx, key, field).Result()
}

func GetHashValue(biz, key, field string) (map[string]string, error) {
	redisClient := GetRedisClient(biz).Client
	return redisClient.HGetAll(ctx, key).Result()
}

func SetKeyValue(biz, key, value string) error {
	redisClient := GetRedisClient(biz).Client
	return redisClient.Set(ctx, key, value, 0).Err()
}

func GetKeyValue(biz, key string) (string, error) {
	redisClient := GetRedisClient(biz).Client
	return redisClient.Get(ctx, key).Result()
}

func SetKeyWithExpiration(biz, key, value string, expiration time.Duration) error {
	redisClient := GetRedisClient(biz).Client
	return redisClient.Set(ctx, key, value, expiration).Err()
}

func Test() {
	// 示例操作：设置 Hash 结构中的字段
	err := SetHashField("user","myHash", "field1", "value1")
	if err != nil {
		fmt.Println(err)
	}

	// 示例操作：获取 Hash 结构中的字段值
	val, err := GetHashField("user","myHash", "field1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("field1:", val)
	}

	// 示例操作：设置 Key-Value 结构中的键值对
	err = SetKeyValue("user","myKey", "myValue")
	if err != nil {
		fmt.Println(err)
	}

	// 示例操作：获取 Key-Value 结构中的值
	val, err = GetKeyValue("user","myKey")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("myKey:", val)
	}
}
