package myhttp

import (
	"os"
)

func GetOsConfig() map[string]string {
	// os包 获取系统信息
	hostname, _ := os.Hostname()
	// 获取 总内存
	// map[type]type {size hostname  wd pid ppid uid}
	infos := make(map[string]string)
	infos["hostname"] = hostname
	return infos
}
