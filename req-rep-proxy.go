package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	loadconfig "github.com/Basic-Components/req-rep-proxy/loadconfig"
	proxy "github.com/Basic-Components/req-rep-proxy/proxy"
	log "github.com/sirupsen/logrus"
)

func closeHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		fmt.Println("Device shutdown!")
		os.Exit(0)
	}()

}
func makeConfig() loadconfig.Config {
	var (
		frontendURL string
		backendURL  string
		help        bool
		debug       bool
		logFormat   string
		logOutput   string
		configPath  string
	)
	// 解析命令行参数
	flag.BoolVar(&help, "help", false, "帮助命令")
	flag.BoolVar(&debug, "debug", false, "是否使用debug模式启动")
	flag.StringVar(&frontendURL, "frontend_url", "", "前端连接的地址")
	flag.StringVar(&backendURL, "backend_url", "", "后端绑定的地址")
	flag.StringVar(&logFormat, "log_format", "", "设定log的形式")
	flag.StringVar(&logOutput, "log_output", "", "设定log输出的流位置")
	flag.StringVar(&configPath, "config_path", "", "设定log输出的流位置")
	flag.Parse()

	//构造config
	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var config = loadconfig.LoadConfig(configPath)
	config.Debug = debug
	if frontendURL != "" {
		config.FrontendURL = frontendURL
	}
	if backendURL != "" {
		config.BackendURL = backendURL
	}
	if logFormat != "" {
		config.LogFormat = logFormat
	}
	if logOutput != "" {
		config.LogOutput = logOutput
	}
	return config
}

func main() {
	var config = makeConfig()
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	switch config.LogFormat {
	case "json":
		{
			log.SetFormatter(&log.JSONFormatter{})
		}
	case "text":
		{
			log.SetFormatter(&log.TextFormatter{})
		}
	default:
		{
			log.SetFormatter(&log.JSONFormatter{})
		}
	}
	if config.LogOutput == "" {
		log.SetOutput(os.Stdout)
	} else {
		logFile, _ := os.OpenFile("logOutput", os.O_CREATE|os.O_WRONLY, 0666)
		defer logFile.Close()
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}
	log.WithFields(log.Fields{
		"Device": "req-rep-proxy",
	}).Info("Server start !")
	log.WithFields(log.Fields{
		"Device": "req-rep-proxy",
	}).Info("Client should connect to url %s", config.FrontendURL)
	log.WithFields(log.Fields{
		"Device": "req-rep-proxy",
	}).Info("Server should connect to url %s", config.BackendURL)
	closeHandler()
	proxy.Proxy(config.FrontendURL, config.BackendURL)
}
