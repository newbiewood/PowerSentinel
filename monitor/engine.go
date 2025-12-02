package monitor

import (
	"PowerSentinel/config"
	"PowerSentinel/notify"
	"PowerSentinel/system"
	"fmt"
	"math"
	"time"

	"github.com/distatus/battery"
)

// Start 启动监控循环
func Start(cfg *config.AppConfig) {
	ticker := time.NewTicker(time.Duration(cfg.CheckInterval) * time.Second)

	// 状态标记，防止重复提醒
	var (
		lastPluggedState   = true // 假设初始是接通的
		lowBatteryNotified = false
	)

	fmt.Println("⚡ PowerSentinel 正在后台运行...")

	for range ticker.C {
		batteries, err := battery.GetAll()
		if err != nil || len(batteries) == 0 {
			continue
		}

		// 通常取第一块电池
		batt := batteries[0]

		// 计算百分比
		currentPct := int(math.Round((batt.Current / batt.Full) * 100))

		// --- 修复部分：使用 String() 获取状态字符串 ---
		// 状态包括: "Unknown", "Empty", "Full", "Charging", "Discharging"
		stateStr := batt.State.String()

		// 判断是否充电/接通电源 (Charging 或 Full 或 Unknown 都视为接通，只有 Discharging 是明确的未接通)
		isPlugged := stateStr == "Charging" || stateStr == "Full" || stateStr == "Unknown"

		// --- 逻辑 1: 电源连接状态改变提醒 ---
		if lastPluggedState == true && isPlugged == false {
			// 刚刚拔掉电源
			notify.ShowAlert("电源断开", fmt.Sprintf("系统已切换至电池供电，当前电量: %d%%", currentPct))
			lowBatteryNotified = false
		}
		lastPluggedState = isPlugged

		// --- 逻辑 2: 低电量自动关机 (仅在未充电时) ---
		if !isPlugged && currentPct <= cfg.ShutdownThreshold {
			notify.ShowAlert("严重警告", "电量低于阈值，系统将在60秒后关机！请立即连接电源！")
			err := system.ShutdownWindows()
			if err != nil {
				fmt.Printf("关机失败: %v\n", err)
			} else {
				break
			}
		}

		// --- 逻辑 3: 低电量弹窗提醒 ---
		if !isPlugged && currentPct <= cfg.AlertThreshold {
			if !lowBatteryNotified {
				notify.ShowAlert("低电量警告", fmt.Sprintf("电量已低于 %d%% (当前: %d%%)，请连接电源。", cfg.AlertThreshold, currentPct))
				lowBatteryNotified = true
			}
		} else if currentPct > cfg.AlertThreshold {
			lowBatteryNotified = false
		}
	}
}
