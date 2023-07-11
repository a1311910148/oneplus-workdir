package main

import (
	"fmt"
	"net/http"
	"servertest/mongodb"
	"strconv"
)

func main() {
	// fmt.Println(mathMod.Add(1, 2))
	// fmt.Println(mathMod.Comp(1, 2))
	// myhttp.Start("3000")
	// fmtpkg.FmtVerbs()
	// fmtpkg.For()
	// fmtpkg.Defer()
	// scanpck.UserInput()
	// spider.SpiderPage()
	// ui.Render()
	// ins := scanpck.Factory("admin", map[string]interface{}{"name": "admin", "age": 18, "password": "123456", "account": "admin", "level": 1})
	// fmt.Println(ins)
	// mongodb.Test()
	// 启动一个http服务器 用于接收请求
	// http router:
	// /api/adduser Get
	// 回调函数获取查询字符串
	http.HandleFunc("/api/adduser", func(w http.ResponseWriter, r *http.Request) {
		defer fmt.Println("发生错误")
		// 获取查询字符串
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		email := r.URL.Query().Get("email")
		ageInt, _ := strconv.Atoi(age)
		mongodb.AddUser(mongodb.User{Name: name, Age: ageInt, Email: email})
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":3000", nil)
}
