package main

import (
	"flag"
	"fmt"
	"net"
)

var (
	serverIp   string
	serverPort int
)

func init() {
	// 初始化命令参数
	// -ip string
	//		服务器IP地址 (default "
	// -port string
	//		服务器端口号 (default "8888")
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器IP地址")
	flag.IntVar(&serverPort, "port", 8888, "服务器端口号")
}

func main() {
	flag.Parse()
	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 单独开启一个goroutine去处理服务器的回执消息
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	// 主协程负责发送消息
	go func() {
		for {
			// 输出菜单
			fmt.Println("1.公聊模式")
			fmt.Println("2.私聊模式")
			fmt.Println("3.更新用户名")
			fmt.Println("0.退出")
			// 获取用户输入
			var msg string
			fmt.Scanln(&msg)
			// 判断是否有效
			if msg == "1" {
				conn.Write([]byte(msg))
				fmt.Println("请输入聊天内容，exit退出")
			} else if msg == "2" {
				conn.Write([]byte(msg))
				fmt.Println("请输入聊天内容，exit退出")
			} else if msg == "3" {
				conn.Write([]byte(msg))
				fmt.Println("请输入用户名：")
			} else {
				fmt.Println("已退出")
			}
		}
	}()

}
