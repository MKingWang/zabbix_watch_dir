## 监控指定目录下文件的创建

问题1 目中只能做到 监控创建和编辑 删除等操作 无法监控是否关闭
> 已经在更换notify库后解决

问题2 在测试中 偶尔回出现 多次写入


**使用方法** 
```
go build -o ftpNotify main.go queue.go watch.go
./ftpNotify dir
```

tools 里面是测试，用于不断从队列中取出数据

***关于redis地址等参数，暂时使用全局变量写死在代码中，根据需要可以移出到配置文件***
