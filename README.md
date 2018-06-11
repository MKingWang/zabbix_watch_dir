## 监控指定目录下文件的创建

问题1 目中只能做到 监控创建和编辑 删除等操作 无法监控是否关闭
> 已经在更换notify库后解决

问题2 在测试中 偶尔回出现 多次写入


**使用方法** 
```
make all
cp noitfyServer noityfCli $ZABBIX_SCRIPT
setsid ./no
./ftpNotify $DIR
```

自定义消息队列服务，监听端口建立在8888

编写客户端程序，存放在client目录中 用于在zabbix监控时直接从server端获取文件操作