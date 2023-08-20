package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Init 为函数中使用的变量设置初始值.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(getRandMessage(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string)

	for _, v := range names {
		// 生成欢迎信息
		msg, err := Hello(v)

		if err != nil {
			return nil, err
		}

		msgs[v] = msg

	}

	return msgs, nil
}

func getRandMessage() string {
	msgs := []string{
		"你好，%v欢迎你",
		"emmm,您来了,%v",
		"呼呼,%v,您来啦",
	}
	return msgs[rand.Intn(3)]
}
