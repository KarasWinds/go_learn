package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis.Dial err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("conn sussess...")
	_, err = conn.Do("HSet", "user01", "name", "tom")
	if err != nil {
		fmt.Println("hset err:", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err:", err)
		return
	}

	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err:", err)
		return
	}
	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget err:", err)
		return
	}
	fmt.Printf("get name :%v,age:%v\n", r1, r2)
}
