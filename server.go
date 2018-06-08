package main

import (
	"log"
	"net"
)

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

func clientHandle(conn net.Conn, queue <-chan string) {
	defer conn.Close()

	rep := <-queue
	conn.Write([]byte(rep))
}
