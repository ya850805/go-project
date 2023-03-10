package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("CPU個數=", cpuNum)

	//可以自己設置使用多少個CPU
	runtime.GOMAXPROCS(cpuNum - 1)
}
