package system

import (
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

// ShutdownWindows 执行关机命令
// 使用 shutdown /s /t 60 给予用户60秒缓冲时间取消
func ShutdownWindows() error {
	cmd := exec.Command("shutdown", "/s", "/t", "60", "/c", "PowerSentinel: 电池电量极低，系统即将关闭")
	return cmd.Run()
}

// SetAutoStart 设置或取消开机自启
func SetAutoStart(enable bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	appName := "PowerSentinel"

	if enable {
		exePath, err := os.Executable()
		if err != nil {
			return err
		}
		return k.SetStringValue(appName, exePath)
	} else {
		return k.DeleteValue(appName)
	}
}
