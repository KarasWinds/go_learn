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
	_, err = conn.Do("HMSet", "user01", "name", "tom", "age", 29)
	if err != nil {
		fmt.Println("HMset err:", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet", "user01", "name", "age"))
	if err != nil {
		fmt.Println("hMget err:", err)
		return
	}

	fmt.Printf("get name :%v,age:%v\n", r[0], r[1])
}
