package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func initRedis(redisAddr string) (redis.Conn, error) {
	return redis.Dial("tcp", redisAddr)
}

/*
将节点内容写入redis队列
*/
func pushQueue(key interface{}) error {
	redisCli, err := initRedis("192.168.25.57:6379")
	if err != nil {
		fmt.Errorf("init redis err:%v", err)
		return err
	}
	defer redisCli.Close()

	_, err = redisCli.Do("LPUSH", "genequeue", key)
	if err != nil {
		fmt.Errorf("push queue err:%v", err)
		return err
	}

	return nil
}
