package main

import (
	"fmt"
	"math/rand"
	"net/url"
)

type T struct {
	A int
	B int
}

func main() {
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Opt())

	fmt.Println(url.JoinPath("/a", "/bc"))
	fmt.Println(url.PathEscape("/a市中心"))
	fmt.Println(url.PathUnescape("%2Fa%E5%B8%82%E4%B8%AD%E5%BF%83"))
	// m, _ := url.ParseQuery(`x=1&y=2&name=songzhixin`)
	fmt.Println(rand.Int())
	// fmt.Println(math.Max(1, 2))
	// var t T
	// yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
	// fmt.Println(t)

	// yamlStr := "a: 1\nb: 2"

	// var dc = yaml.NewDecoder(strings.NewReader(yamlStr))
	// var r T
	// dc.Decode(&r)
	// fmt.Println(r)

	// yamlStr := T{A: 1, B: 2}

	// encoder := yaml.NewEncoder(os.Stdout)

	// encoder.Encode(yamlStr)

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
	// http rouyaNewEncoder(w io.Writer)
	// /api/adduser Get
	// 回调函数获取查询字符串
	// http.HandleFunc("/api/adduser", func(w http.ResponseWriter, r *http.Request) {
	// 	defer fmt.Println("发生错误")
	// 	// 获取查询字符串
	// 	name := r.URL.Query().Get("name")
	// 	age := r.URL.Query().Get("age")
	// 	email := r.URL.Query().Get("email")
	// 	// name, age, email := GetUserQuery(r)
	// 	ageInt, _ := strconv.Atoi(age)
	// 	mongodb.AddUser(mongodb.User{Name: name, Age: ageInt, Email: email})
	// 	w.Write([]byte("ok"))
	// })
	// http.ListenAndServe(":3001", nil)
}
