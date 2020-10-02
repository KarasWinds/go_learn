package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("conn.Do(Set) err:", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do(Get) err:")
		return
	}
	fmt.Println("name:", r)
}
