package incomemessage

import (
	"fmt"
	"message/user"
)

type Message struct {
	Type    string
	Name    string
	Content string
	From    string
	To      string
}

// 广播消息
func (s *Message) BroadCast(userArr []user.User, m Message) {
	// 判断广播类型
	switch m.Type {
	case "login":
		for _, user := range userArr {
			if user.Name == m.Name {
				user.Conn.Write([]byte("用户名重复"))
				return
			}
		}
	case "message":
		for _, user := range userArr {
			user.Conn.Write([]byte(m.Content))
		}
	default:
		fmt.Println("广播类型错误")
	}
}
