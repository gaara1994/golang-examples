package main

import (
	"go.uber.org/zap"
)

func main() {
	//Zap提供了两种类型的日志记录器 Sugared Logger 和 Logger
	//1.Sugared Logger 支持结构化和printf风格的日志记录。适合开发的时候使用
	sugar, _ := zap.NewDevelopment()
	defer sugar.Sync()

	sugar.Info("无法获取网址")
	//2022-09-15T11:00:55.579+0800    INFO    z_test/main.go:13       无法获取网址

	//2.Logger 比 SugaredLogger 更快内存开销小,是json格式的日志，适合线上使用
	logger, _ := zap.NewProduction()
	logger.Info("无法获取网址2")
	//{"level":"info","ts":1663210855.579693,"caller":"z_test/main.go:18","msg":"无法获取网址2"}

	//sugar 和 logger之间是可以在上游转换的，这很简单
	logger2sugar := logger.Sugar()
	logger2sugar.Info("无法获取网址3")

	//日志不只是打印，还需要写进文件，这就得需要 zap.New()方法来手动传递所有配置


}
