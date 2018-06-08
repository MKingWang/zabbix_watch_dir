## 监控指定目录下文件的创建

问题1 目中只能做到 监控创建和编辑 删除等操作 无法监控是否关闭
> 已经在更换notify库后解决

问题2 在测试中 偶尔回出现 多次写入


**使用方法** 
```
go build -o ftpNotify main.go queue.go watch.go
./ftpNotify dir
```

自定义消息队列服务，监听端口建立在8888