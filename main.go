package main

import (
	"flag"
	"math/rand"
	"time"
	"www.miniton-gateway.com/boot"
)

var (
	mode = flag.String("m", "dev", "run mode,dev/test/prod")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	// 初始化
	boot.Init(*mode)
	// 启动
	boot.Run()
	// 注册退出信号
	boot.AwaitSignal()
}
