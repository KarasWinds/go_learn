package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("time type:%T ,value:%v\n", now, now)

	fmt.Printf("年:%v\n", now.Year())
	fmt.Printf("月:%v\n", now.Month())
	fmt.Printf("月:%v\n", int(now.Month()))
	fmt.Printf("日:%v\n", now.Day())
	fmt.Printf("周:%v\n", now.Weekday())
	fmt.Printf("周:%v\n", int(now.Weekday()))
	fmt.Printf("時:%v\n", now.Hour())
	fmt.Printf("分:%v\n", now.Minute())
	fmt.Printf("秒:%v\n", now.Second())

	fmt.Printf("現在時間: %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("現在時間: %d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("unix時間戳:%v,unixnano時間戳:%v\n", now.Unix(), now.UnixNano())
}
