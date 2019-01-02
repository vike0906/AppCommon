package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"log"
)

const (
	redisURL            = "redis://127.0.0.1:6379"
	redisMaxIdle        = 3   //最大空闲连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
	redisPassword       = "123456"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
				log.Print("redis connection error...")
				panic(err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", redisPassword); authErr != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
				log.Print("redis auth password error...")
				panic(authErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
				log.Print("ping redis error...")
				panic(err)
			}
			return nil
		},
	}
}

func Set(k, v string) {
	c := RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetStringValue(k string) string {
	c := RedisPool.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

func SetKeyExpire(k string, ex int) {
	c := RedisPool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func SetJSONEX(key, value interface{}, ex int) {
	c := RedisPool.Get()
	defer c.Close()
	jsonByte, _ := json.Marshal(value)
	_, err := c.Do("SET", key, jsonByte, "EX", ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetJsonByte(key string) ([]byte, error) {
	c := RedisPool.Get()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonGet, nil
}

func CheckKey(k string) bool {
	c := RedisPool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := RedisPool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
