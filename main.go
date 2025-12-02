package main

import (
	"PowerSentinel/config"
	"PowerSentinel/monitor"
	"PowerSentinel/system"
	"log"
)

func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("无法加载配置: %v", err)
	}

	// 2. 设置开机自启
	// 注意：这可能会被杀毒软件拦截，建议第一次运行时以管理员权限运行
	if err := system.SetAutoStart(cfg.AutoStart); err != nil {
		log.Printf("警告: 无法设置开机自启: %v", err)
	}

	// 3. 启动监控引擎
	// 这是一个阻塞调用，会一直运行
	monitor.Start(cfg)
}
