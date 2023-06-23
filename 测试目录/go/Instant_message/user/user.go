package user

import "net"

// Conn连接的用户对对象
type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}
