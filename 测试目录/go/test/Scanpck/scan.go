package scanpck

import (
	"fmt"
	"main/config"
)

func ScanfTest() {
	num1 := 0
	num2 := 1
	fmt.Println("请输入第一个数字")
	fmt.Scanf("%d", &num1)
	fmt.Println("请输入第二个数字")
	fmt.Scanf("%d", &num2)
	fmt.Println(num1, num2)
}

func ScanTest() {
	num1 := 0
	num2 := 1
	fmt.Println("请输入第一个数字")
	fmt.Scan(&num1)
	fmt.Println("请输入第二个数字")
	fmt.Scan(&num2)
	fmt.Println(num1, num2)
}

func UserInput() {
	num1 := 0
	name := ""
	age := 0
	// 获取用户选择
	fmt.Println("请输入您的选择：")
	fmt.Scanf("%d", &num1)
	// 获取用户输入的姓名
	fmt.Println("请输入您的姓名：")
	fmt.Scanf("%s", &name)
	// 获取用户输入的年龄
	fmt.Println("请输入您的年龄：")
	fmt.Scanf("%d", &age)
	// 汇总输出
	fmt.Printf("您的选择是：%d，您的姓名是：%s，您的年龄是：%d\n", num1, name, age)
	// 打印浮点数config.Version
	fmt.Printf("config.Version: %f\n", config.Version)
	fmt.Println(config.VpsIp1)
	fmt.Println(config.VpsIp2)
}

type user struct {
	Name     string
	Age      int
	Password string
	Account  string
}

func (u *user) SetName(name string) {
	u.Name = name
}

func (u *user) SetAge(age int) {
	u.Age = age
}

func (u *user) SetPassword(password string) {
	u.Password = password
}

var u user = user{
	Name:     "张三",
	Age:      18,
	Password: "123456",
	Account:  "zhangsan",
}

func GetUser() user {
	return u
}
