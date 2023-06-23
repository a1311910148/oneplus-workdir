package main

import (
	"fmt"
	"message/server"
)

func main() {
	fmt.Println("start")
	server.CreateServer().Start()
}
