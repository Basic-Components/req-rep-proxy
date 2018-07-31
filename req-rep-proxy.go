package main

import (
	"os"
	"os/signal"
	"syscall"

	parser "github.com/Basic-Components/req-rep-proxy/argparser"
	proxy "github.com/Basic-Components/req-rep-proxy/proxy"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func closeHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("\r- Ctrl+C pressed in Terminal")
		log.Info("Device shutdown!")
		os.Exit(0)
	}()

}

func main() {
	closeHandler()
	frontend, backend, debug := parser.Parser()
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("Device start!")
	proxy.Proxy(frontend, backend)

}
