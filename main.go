package main

import (
	"os"
)

func main() {
	queue := make(chan string, 100)
	filename := os.Args[1]
	go watchDir(filename, queue)
	queueServer(queue)
}
