# req-rep-proxy
req-rep load balancing proxy请求应答模式的负载均衡中间件

前后端都为绑定ip,需要先用server连接本组件,之后客户端再发送请求

## 使用方法

使用命令行`.req-rep-proxy`启动组件,下面是可选的参数:

| 标志         | 默认值       | 说明                  |
| ------------ | ------------ | --------------------- |
| -h           | false        | 帮助命令              |
| -d           | false        | 是否使用debug模式启动 |
| -f           | tcp://*:5559 | 前端连接的地址        |
| -b           | tcp://*:5560 | 后端绑定的地址        |
| --log_format | json         | 设定log的形式         |
| --log_output | ""           | 设定log输出的流位置   |

启动的时候按需求填入参数,