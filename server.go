package main

import (
	"fmt"
	"log"
	"net"
)

/*
服务端监听程序
*/
func queueServer(queue chan string) {
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go clientHandle(conn, queue)
	}
}

/*
客户端处理程序
*/
func clientHandle(conn net.Conn, queue <-chan string) {
	defer conn.Close()
	var strs string

	//从chan中读取所有数据输出给客户端
endloop:
	for {
		select {
		case rep := <-queue:
			strs = fmt.Sprintf("%s%s\n", strs, rep)
		default:
			break endloop
		}
	}
	conn.Write([]byte(strs))
}
