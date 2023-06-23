package server

import (
	"encoding/json"
	"fmt"
	incomemessage "message/incomeMessage"
	"message/user"
	"net"
)

type Server struct {
	Ip      string
	Port    string
	UserArr []user.User
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

// process
func (s *Server) process(Conn net.Conn) {
	// 读取客户端发送的信息
	defer Conn.Close()
	adrr := Conn.RemoteAddr().String()

	// 实例化User
	u := &user.User{
		Conn: Conn,
		C:    make(chan string),
		Name: "",
		Addr: adrr,
	}

	// 添加用户
	s.UserArr = append(s.UserArr, *u)

	for {
		buf := make([]byte, 1024)
		n, err := Conn.Read(buf) //Conn.Read方法
		if err != nil {
			fmt.Println("Conn.Read err:", err)
			return
		}

		// 解析json
		var m incomemessage.Message

		json.Unmarshal(buf[:n], &m)

		fmt.Println("adrr:", adrr, "content:", m.Content) //打印消息

		// 判断是否为登录
		if m.Type == "login" {
			// 设置用户姓名
			u.Name = m.Name
			// 向其他用户广播
			for _, user := range s.UserArr {
				if user.Name != m.Name {
					user.Conn.Write([]byte(m.Name + "online\n"))
				}
			}
			return
		}

		// 点对点通信
		if m.To != "" {
			for _, user := range s.UserArr {
				if user.Name == m.To {
					user.Conn.Write([]byte(m.Content))
				}
			}
			return
		}

		// 广播消息
		for _, user := range s.UserArr {
			user.Conn.Write([]byte(m.Content))
		}

		// {"type":"login","name":"wanglei"}
		// {"type":"message","name":"wanglei","content":"is here people ","from":"wanglei","to":""}
		// 判断是否收信方
		// if m.To != "" {
		// 	// 遍历ConnArr
		// 	for _, user := range s.UserArr {
		// 		if user.Name == m.To {
		// 			user.Con .Write([]byte(m.Content))
		// 		}
		// 	}
		// 	return
		// }

	}
}

func CreateServer() *Server {
	return &Server{
		Ip:   "127.0.0.1",
		Port: "8888"}
}
