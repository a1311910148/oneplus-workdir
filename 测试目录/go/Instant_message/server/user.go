package server

import (
	"net"
)

// Conn连接的用户对对象
type User struct {
	Name   string
	Addr   string
	C      chan string
	Conn   net.Conn
	Server *Server
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		user.Conn.Write([]byte(msg + "\n"))
	}
}

func (user *User) Message(msg Message) {
	if msg.Type == "rename" {
		// 判断用户名是否存在
		for _, v := range user.Server.UserArr {
			if v.Name == msg.Content {
				user.Conn.Write([]byte("用户名已存在\n"))
				return
			}
		}
		// 修改用户名 点对点
		user.Name = msg.Content
		user.Conn.Write([]byte("用户名已修改为" + msg.Content + "\n"))
	} else if msg.Type == "privateChat" {
		// 私聊用户
		for _, v := range user.Server.UserArr {
			if v.Name == msg.To {
				v.C <- msg.Content
				return
			}
		}
	} else {
		// 广播消息
		for _, v := range user.Server.UserArr {
			if v.Name != user.Name {
				v.C <- msg.Content
			}
		}
	}

}

func CreateUser(Conn net.Conn) *User {
	// 获取地址
	adrr := Conn.RemoteAddr().String()

	user := &User{
		Name: adrr,
		C:    make(chan string),
		Conn: nil,
	}
	// 开启一个协程
	go user.ListenMessage()

	return user

}
