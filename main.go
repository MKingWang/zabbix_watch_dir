package main

import (
	"os"
)

func main() {
	filename := os.Args[1]
	go watchDir(filename)

	select {}

}
