package fmtpkg

import (
	"fmt"
	"strconv"
	"strings"
)

func FmtVerbs() {
	a := 122
	fmt.Printf("a=%v\na的地址=%b\na的地址=%v\n", a, a, &a)

	str := "hello"
	fmt.Printf("a=%v\n\na的地址=%v\n", str, &str)
	ptr := &str
	*ptr += "world"
	fmt.Println(str)

	ip := "192.168.123.185"
	arr := strings.Split(ip, ".")
	fmt.Println(arr)
	// ipArr := make([]int, 4)
	num, _ := strconv.Atoi(arr[0])
	num1, _ := strconv.Atoi(arr[1])
	num2, _ := strconv.Atoi(arr[2])
	num3, _ := strconv.Atoi(arr[3])
	fmt.Println(num & 255)
	fmt.Println(num1 & 255)
	fmt.Println(num2 & 255)
	fmt.Println(num3 & 0)

}

func For() {
	index := 1
	for {
		index += 1
		fmt.Println(index)
	}
}

// 函数

func Function() {
	sum := add(1, 2)
	fmt.Println(sum)

	// 匿名函数
	sub := func() int {
		return 1
	}

	fmt.Println(sub())

	// 匿名函数自调用
	func() {
		fmt.Println("匿名函数自调用")
	}()
}

// 函数声明
func add(a, b int) int {
	return a + b + 2
}

// defer

func Defer() {
	defer func() {
		fmt.Println("go")
	}()

	a := 1
	b := 1
	var num int
	arr := []int{a, b}
	s := []int{4, 4}
	func() {
		num = a + b
	}()

	fmt.Println(num, arr, s)
	s1 := append(s, 1)
	fmt.Println(s1, s)
	s1[0] = 222
	fmt.Println(s1, s)

	for i, v := range s1 {
		fmt.Println(i, v)
	}
	// f1 := 2.333
	//  f2 = f1的占用字节数
	// f2 := math.Float64bits(f1)
}
