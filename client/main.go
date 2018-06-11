package main

import (
	"fmt"
	"net"

	"github.com/garyburd/redigo/redis"
)

func initRedis(redisAddr string) (redis.Conn, error) {
	return redis.Dial("tcp", redisAddr)
}

/*
从reids获取消息
*/
func gethQueuFromRedis() error {
	redisCli, err := initRedis("192.168.21.143:6379")
	if err != nil {
		fmt.Errorf("init redis err:%v", err)
		return err
	}
	defer redisCli.Close()
	for {

		info, err := redis.Strings(redisCli.Do("BRPOP", "genequeue", 20))
		if err != nil {
			fmt.Errorf("push queue err:%v", err)
			return err
		}

		fmt.Println(info[1])

	}
	return nil
}

/*
从队列服务器获取消息
*/
func gethQueuFromServer() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	buf := make([]byte, 4096)
	if err != nil {
		fmt.Errorf("Client init error:%v", err)
	}
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Errorf("%v", err)

	}
	fmt.Println(string(buf))
}

func main() {
	gethQueuFromServer()
}
