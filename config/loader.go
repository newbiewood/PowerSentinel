package config

import (
	"encoding/json"
	"os"
)

// AppConfig 定义配置文件结构，包含实际配置项和用于说明的字段
type AppConfig struct {
	// 实际配置项
	AlertThreshold            int    `json:"alert_threshold"`
	AlertThresholdDescription string `json:"_alert_threshold_说明"` // 新增说明字段

	ShutdownThreshold            int    `json:"shutdown_threshold"`
	ShutdownThresholdDescription string `json:"_shutdown_threshold_说明"` // 新增说明字段

	CheckInterval            int    `json:"check_interval"`
	CheckIntervalDescription string `json:"_check_interval_说明"` // 新增说明字段

	AutoStart            bool   `json:"auto_start"`
	AutoStartDescription string `json:"_auto_start_说明"` // 新增说明字段
}

// LoadConfig 读取配置文件，如果不存在则创建默认配置
func LoadConfig(filename string) (*AppConfig, error) {
	// 默认配置
	defaultConfig := &AppConfig{
		AlertThreshold:            20,
		AlertThresholdDescription: "电量低于此数值时弹窗提醒（如果未连接电源）。",

		ShutdownThreshold:            5,
		ShutdownThresholdDescription: "电量低于此数值时，系统将发出警告并在60秒后自动关机（如果未连接电源）。",

		CheckInterval:            30,
		CheckIntervalDescription: "程序检测电池状态的时间间隔（单位：秒）。",

		AutoStart:            true,
		AutoStartDescription: "程序是否自动添加到Windows开机启动项。",
	}

	file, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		// 创建默认配置文件
		// 使用 MarshalIndent 保证输出美观
		bytes, _ := json.MarshalIndent(defaultConfig, "", "    ")
		_ = os.WriteFile(filename, bytes, 0644)
		return defaultConfig, nil
	} else if err != nil {
		return nil, err
	}

	// 注意：Unmarshal 会自动忽略配置文件中多余的说明字段，只填充我们需要的四个配置值
	// 但为了确保用户能看到最新的说明，我们从默认配置开始 Unmarshal
	err = json.Unmarshal(file, defaultConfig)
	return defaultConfig, err
}
