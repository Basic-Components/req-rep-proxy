//
// 发布订阅模式的代理组件,沟通发布者和订阅者,降低发布者负载
//
package proxy

import (
	zmq "github.com/pebbe/zmq4"
)

func switchMessages(poller *zmq.Poller, frontend *zmq.Socket, backend *zmq.Socket) {
	sockets, _ := poller.Poll(-1)
	for _, socket := range sockets {
		switch s := socket.Socket; s {
		case frontend:
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
func Proxy(bind_frontend string, bind_backend string) {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	frontend.Bind(bind_frontend)
	backend.Bind(bind_backend)

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	//  Switch messages between sockets
	for {
		switchMessages(poller, frontend, backend)
	}
}
