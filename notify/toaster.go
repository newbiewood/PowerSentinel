package notify

import (
	"github.com/gen2brain/beeep"
)

func ShowAlert(title, message string) {
	// 在Windows上会显示为系统通知或弹窗
	err := beeep.Alert(title, message, "")
	if err != nil {
		// 如果弹窗失败，尝试简单的Notify
		beeep.Notify(title, message, "")
	}
}
