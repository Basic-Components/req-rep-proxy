# req-rep-proxy

请求应答模式的负载均衡中间件

前后端都为绑定ip,需要先用server连接本组件,之后客户端再发送请求

## 使用方法

使用命令行`.req-rep-proxy`启动组件,下面是可选的参数:

| 标志          | 默认值       | 说明                  |
| ------------- | ------------ | --------------------- |
| -help         | false        | 帮助命令              |
| -debug        | false        | 是否使用debug模式启动 |
| -server_name  | unknown      | 后端服务名            |
| -frontend_url | tcp://*:5559 | 前端连接的地址        |
| -backend_url  | tcp://*:5560 | 后端绑定的地址        |
| -log_format   | json         | 设定log的形式         |
| -log_output   | ""           | 设定log输出的流位置   |
| -config_path  | ""           | 设定读取配置文件地址  |

启动的时候按需求填入参数.

配置文件为json格式,以下为默认配置的配置文件形式:

```json
{
	"server_name":"unknown",
	"frontend_url":"tcp://*:5559",
	"backend_url":"tcp://*:5560",
	"debug":false,
	"log_format":"json",
	"log_output":"req-rep-proxy.log"
}
```

配置的优先级为: `命令行参数>配置文件>默认`

## 通过docker使用

镜像为:`hsz1273327/req-rep-proxy`,一个可以参考的使用方式是执行:`docker run -p 5559:5559 -p 5560:5560  hsz1273327/req-rep-proxy ./req-rep-proxy -debug`

但通常这个组件是一个服务群的对外端点,使用`docer-compose.yml`进行编排,细节不表,这边给出一个参考配置文件:

```yml
version: '3'
services:

  # ############################################代理
  proxy:
    image: hsz1273327/req-rep-proxy:latest
    networks:
      - out 
      - server-group
    command: ./req-rep-proxy

 # ############################################实际的服务
  sev1:
    image: xxx:latest
    networks:
      - server-group
    command: sev

  sev2:
    image: xxx:latest
    networks:
      - server-group
    command: sev

# ############################################配置网络
networks:
  out:
    external: true
  server-group:
    external: true
```
