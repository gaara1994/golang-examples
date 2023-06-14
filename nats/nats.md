

# 1.安装

1.服务端 nats-server: [Releases · nats-io/nats-server (github.com)](https://github.com/nats-io/nats-server/releases)

```shell
[root@localhost ~]# wget https://github.com/nats-io/nats-server/releases/download/v2.9.17/nats-server-v2.9.17-amd64.rpm
[root@localhost ~]# yum install -y nats-server-v2.9.17-amd64.rpm 
```

启动服务端

```shell
[root@localhost ~]# nats-server
[10246] 2023/06/13 09:45:43.820464 [INF] Starting nats-server
[10246] 2023/06/13 09:45:43.820503 [INF]   Version:  2.9.17
[10246] 2023/06/13 09:45:43.820505 [INF]   Git:      [4f2c9a5]
[10246] 2023/06/13 09:45:43.820507 [INF]   Name:     NAQDNB6XVYGIYICP5B2DGMP2F63R7ZWNAVWL2RVP3FGWI3XYUFHJNMZW
[10246] 2023/06/13 09:45:43.820508 [INF]   ID:       NAQDNB6XVYGIYICP5B2DGMP2F63R7ZWNAVWL2RVP3FGWI3XYUFHJNMZW
[10246] 2023/06/13 09:45:43.820912 [INF] Listening for client connections on 0.0.0.0:4222
[10246] 2023/06/13 09:45:43.821491 [INF] Server is ready

```

2.客户端 nats-cli：[Releases · nats-io/natscli (github.com)](https://github.com/nats-io/natscli/releases)

```shell
[root@localhost ~]# wget https://github.com/nats-io/natscli/releases/download/v0.0.35/nats-0.0.35-amd64.rpm
[root@localhost ~]# yum install -y nats-0.0.35-amd64.rpm
```



# 2.NATS的三种模式

## 1.发布/订阅

侦听某个主题的订阅者会收到有关该主题发布的消息。如果订阅者未主动侦听主题，则不会收到消息。

当订阅者注册自己以接收来自发布者的消息时，消息传递的 1：N 扇出模式可确保发布者发送的任何消息都能到达已注册的所有订阅者。

![sub](.\images\sub.svg)



1.先有订阅者

```shell
#订阅科技主题
[root@localhost ~]# nats sub keji
14:23:47 Subscribing on keji
```

```shell
#订阅科技主题
[root@localhost ~]# nats sub keji
14:24:08 Subscribing on keji
```

```shell
#订阅经济主题
[root@localhost ~]# nats sub jingji
14:24:39 Subscribing on jingji
```



2.发布者

```shell
[root@localhost ~]# nats pub keji 英伟达发布了全新显卡H100
14:25:41 Published 34 bytes to "keji"
```



3.可以看到只有订阅了科技主题的订阅者才会收到这条消息，而订阅了经济主题的订阅者没有收到消息。

```shell
[root@localhost ~]# nats sub keji
14:23:47 Subscribing on keji
[#1] Received on "keji"
英伟达发布了全新显卡H100
```

```shell
[root@localhost ~]# nats sub keji
14:24:08 Subscribing on keji
[#1] Received on "keji"
英伟达发布了全新显卡H100
```

```shell
[root@localhost ~]# nats sub jingji
14:24:39 Subscribing on jingji
```







## 2.请求-回复

请求-回复是现代分布式系统中的常见模式。发送请求，应用程序要么等待具有特定超时的响应，要么异步接收响应。

![requ](.\images\requ.svg)



运行回复客户端的监听者。

```shell
[root@localhost ~]# nats reply help.message '我来处理'
15:00:55 Listening on "help.message" in group "NATS-RPLY-22"
```

```shell
[root@localhost ~]# nats request help.message '谁来解决这个问题？'
15:00:45 Sending request on "help.message"
15:00:45 Received with rtt 649.899µs
我来处理
```



## 3.消息队列

