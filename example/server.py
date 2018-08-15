import sys
import argparse
import zmq


def _server(args):
    context = zmq.Context()
    socket = context.socket(zmq.REP)
    socket.connect(args.url)
    print(f"server connect @ {args.url} ")
    while True:
        message = socket.recv()
        print(f"Received request: {message},send back!")
        socket.send(message)


def _parser_args(params):
    """解析命令行参数."""
    parser = argparse.ArgumentParser()
    parser.add_argument('--url', type=str, default="tcp://localhost:5560", help="指定连接到哪个组件")
    parser.set_defaults(func=_server)
    args = parser.parse_args(params)
    args.func(args)


def main(argv=sys.argv[1:]):
    u"""服务启动入口.

    设置覆盖顺序`命令行参数`>`'-c'指定的配置文件`>`项目启动位置的配置文件`>默认配置.
    """
    _parser_args(argv)


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('- Ctrl+C pressed in Terminal')
        print("Server shutdown!")
