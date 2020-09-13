package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("CPU執行緒:", cpuNum)
}
