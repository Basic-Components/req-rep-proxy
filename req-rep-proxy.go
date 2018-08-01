package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

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

func main() {
	var (
		frontend  string
		backend   string
		h         bool
		debug     bool
		logFormat string
		logOutput string
	)

	flag.BoolVar(&h, "h", false, "帮助命令")
	flag.BoolVar(&debug, "d", false, "是否使用debug模式启动")
	flag.StringVar(&frontend, "f", "tcp://*:5559", "前端连接的地址")
	flag.StringVar(&backend, "b", "tcp://*:5560", "后端绑定的地址")
	flag.StringVar(&logFormat, "log_format", "json", "设定log的形式")
	flag.StringVar(&logOutput, "log_output", "", "设定log输出的流位置")

	flag.Parse()
	if h {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	switch logFormat {
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

	if logOutput == "" {
		log.SetOutput(os.Stdout)
	} else {
		logFile, _ := os.OpenFile("logOutput", os.O_CREATE|os.O_WRONLY, 0666)
		defer logFile.Close()
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}
	log.WithFields(log.Fields{
		"device": "req-rep-proxy",
	}).Info("Device start !")
	log.WithFields(log.Fields{
		"device": "req-rep-proxy",
	}).Info("server please connect to %s", backend)
	log.WithFields(log.Fields{
		"device": "req-rep-proxy",
	}).Info("client please connect to %s", frontend)
	closeHandler()
	proxy.Proxy(frontend, backend)

}
