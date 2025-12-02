# README

我的第一个程序
开机自启需第一次管理员启动 配置喜喜在config.json里

golang开发自构建请安装依赖
go get github.com/distatus/battery
go get github.com/gen2brain/beeep
go get golang.org/x/sys/windows/registry

构建
go build -ldflags -H=windowsgui -o PowerSentinel.exe