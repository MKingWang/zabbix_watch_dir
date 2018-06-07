package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func initRedis(redisAddr string) (redis.Conn, error) {
	return redis.Dial("tcp", redisAddr)
}

//定义redis地址和队列key
var redisAddr = "192.168.21.143:6379"
var queue = "genequeue"

/*
将节点内容写入redis队列
*/
func pushQueue(key interface{}) error {
	redisCli, err := initRedis(redisAddr)
	if err != nil {
		fmt.Errorf("init redis err:%v", err)
		return err
	}
	defer redisCli.Close()

	_, err = redisCli.Do("LPUSH", queue, key)
	if err != nil {
		fmt.Errorf("push queue err:%v", err)
		return err
	}

	return nil
}
