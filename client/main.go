package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/garyburd/redigo/redis"
)

func initRedis(redisAddr string) (redis.Conn, error) {
	return redis.Dial("tcp", redisAddr)
}

/*
将节点内容写入redis队列
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

func gethQueuFromServer() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Errorf("Client init error:%v", err)
	}

	reader := bufio.NewReader(conn)
	request, _, err := reader.ReadLine()
	if err != nil {
		fmt.Errorf("read err:%v", err)
	}

	fmt.Println(string(request))

}

func main() {
	gethQueuFromServer()
}
