//
// 发布订阅模式的代理组件,沟通发布者和订阅者,降低发布者负载
//
package proxy

import (
	consts "github.com/Basic-Components/req-rep-proxy/consts"
	loadconfig "github.com/Basic-Components/req-rep-proxy/loadconfig"

	zmq "github.com/pebbe/zmq4"
	log "github.com/sirupsen/logrus"
)

func switchMessages(poller *zmq.Poller, frontend *zmq.Socket, backend *zmq.Socket) {
	sockets, _ := poller.Poll(-1)
	for _, socket := range sockets {
		switch s := socket.Socket; s {
		case frontend:
			log.WithFields(log.Fields{
				consts.TYPE: consts.NAME,
				"Direction": "request",
				"Socket":    s.String()}).Debug("Send message to server!")
			for {
				msg, _ := s.Recv(0)
				if more, _ := s.GetRcvmore(); more {
					backend.Send(msg, zmq.SNDMORE)
				} else {
					backend.Send(msg, 0)
					break
				}
			}
		case backend:
			log.WithFields(log.Fields{
				consts.TYPE: consts.NAME,
				"Direction": "response",
				"Socket":    s.String()}).Debug("Return message to client!")
			for {
				msg, _ := s.Recv(0)
				if more, _ := s.GetRcvmore(); more {
					frontend.Send(msg, zmq.SNDMORE)
				} else {
					frontend.Send(msg, 0)
					break
				}
			}
		}
	}
}

// 代理本体
func Run(config loadconfig.Config) {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	frontend.Bind(config.FrontendURL)
	backend.Bind(config.BackendURL)

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	//  Switch messages between sockets
	for {
		switchMessages(poller, frontend, backend)
	}
}
