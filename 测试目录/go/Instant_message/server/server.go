package server

import (
	"encoding/json"
	"fmt"
	"net"
)

type Server struct {
	Ip      string
	Port    string
	UserArr []User
}

// start server
func (s *Server) Start() {
	// 启动服务器
	Listener, err := net.Listen("tcp", s.Ip+":"+s.Port)
	// 捕捉错误
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 轮询获取连接
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listener.Accept err:", err)
			continue
		}
		// 启动一个协程和客户端保持通讯
		go s.process(Conn)
	}
}

// broadcast
func (s *Server) Broadcast(user *User, content string) {
	// 创建消息
	msg := user.Name + ":" + content

	for _, v := range s.UserArr {
		if v.Name != user.Name {
			v.C <- msg
		}
	}
}

// process
func (s *Server) process(Conn net.Conn) {
	// 读取客户端发送的信息
	defer Conn.Close()

	user := CreateUser(Conn)

	// 广播在线
	s.Broadcast(user, "在线")

	// 添加用户
	s.UserArr = append(s.UserArr, *user)

	for {
		buf := make([]byte, 1024)
		n, err := Conn.Read(buf) //Conn.Read方法
		if n == 0 {
			Conn.Close()
			return
		}
		if err != nil {
			fmt.Println("Conn.Read err:", err)
			return
		}

		// 解析json
		var m Message

		json.Unmarshal(buf[:n], &m)

		fmt.Println("user:", user.Name, "content:", m.Content) //打印消息

		user.Message(m)
	}
}

func CreateServer() *Server {
	return &Server{
		Ip:   "127.0.0.1",
		Port: "8888"}
}
