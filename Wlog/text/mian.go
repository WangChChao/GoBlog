package main

import (
	"Wlog/Debug"
	"time"
)

func main() {
	var logger = Wlog.NewLogger("Debug", Wlog.KB*100) // 只显示Error及更高级的日志
	ExePath, _ := Wlog.GetCurrentPath()               // 获取可执行文件的路径
	logger.Debug("这是很普通的debug一条日志", "可执行文件路径:", ExePath)
	logger.Debug()
	logger.Trace()
	logger.Info()
	logger.Waring()
	logger.Error("这是一条错误日志")
	logger.Fatal()
	logger.SetWriter("WangChChao.log", true)
	defer logger.CloseFile()
	for {
		logger.Debug("这是很普通的debug一条日志", "可执行文件路径:", ExePath)
		logger.Error("从理论上讲，这条日志应该写在文件里")
		time.Sleep(50 * time.Millisecond)
	}
}
