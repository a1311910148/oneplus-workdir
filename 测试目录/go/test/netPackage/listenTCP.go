package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "192.168.123.209:3000")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}
	// 接受连接并处理数据
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败：", err)
			continue
		}

		fmt.Println("接受连接：", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取数据并打印
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取数据失败：", err)
		return
	}

	fmt.Println("接收到数据：", string(buf[:n]))
	conn.Write([]byte("HTTP/1.1 200 OK  \r\nContent-Length: 122\r\n<h1>hello</h1>"))
	fmt.Println([]byte("HTTP/1.1 200 OK  \r\nContent-Length: 122\r\n<h1>hello</h1>"))
}
