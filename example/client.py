import sys
import argparse
import zmq


def _client(args):
    #  Prepare our context and sockets
    context = zmq.Context()
    socket = context.socket(zmq.REQ)
    socket.connect(args.url)

    #  Do 10 requests, waiting each time for a response
    for request in range(1, 11):
        socket.send(b"Hello")
        message = socket.recv()
        print("Received reply %s [%s]" % (request, message))


def _parser_args(params):
    """解析命令行参数."""
    parser = argparse.ArgumentParser()
    parser.add_argument('--url', type=str, default="tcp://localhost:5559", help="指定连接到哪个组件")
    parser.set_defaults(func=_client)
    args = parser.parse_args(params)
    args.func(args)


def main(argv=sys.argv[1:]):
    u"""服务启动入口.

    设置覆盖顺序`命令行参数`>`'-c'指定的配置文件`>`项目启动位置的配置文件`>默认配置.
    """
    _parser_args(argv)


if __name__ == '__main__':
    main()
