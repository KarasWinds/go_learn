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
	_, err = conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("set err:", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err:", err)
		return
	}
	fmt.Println("get name :", r)
}
