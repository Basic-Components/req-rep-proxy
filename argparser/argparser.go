package argparser

import (
	"flag"
	"os"
)

// 解析命令行
func Parser() (string, string, bool) {
	var frontend string
	var backend string
	var h bool
	var debug bool
	flag.BoolVar(&h, "h", false, "帮助命令")
	flag.BoolVar(&debug, "d", false, "是否使用debug模式启动")
	flag.StringVar(&frontend, "f", "tcp://*:5559", "前端连接的地址")
	flag.StringVar(&backend, "b", "tcp://*:5560", "后端绑定的地址")
	flag.Parse()
	if h {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return frontend, backend, debug
}
