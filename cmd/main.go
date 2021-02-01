package main

import (
	"flag"
	"fmt"
	"go-template/pkg/config"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"os"
)

func main() {
	// 配置日志
	log.ConfigLog()

	// flags
	active := flag.String("active", "", "active profile")
	flag.Parse()

	// info
	pwd, _ := os.Getwd()
	log.Info("PWD:", pwd)

	// 加载配置文件
	activeProfile := ".env"
	if *active != "" {
		activeProfile = fmt.Sprintf(".env-%s", *active)
	}
	log.Info("Active Profile:", activeProfile)
	config.LoadEnvFile(activeProfile)

	// 加载路由并阻塞
	route.LoadRoutes()
}
