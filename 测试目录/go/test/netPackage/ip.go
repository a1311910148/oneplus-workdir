package main

import (
	"fmt"
	"net"
)

func main() {
	ipStr := "192.168.123.1"
	ip := net.ParseIP(ipStr)
	fmt.Println(ip.IsGlobalUnicast())
	address := "192.168.123.185:3000"
	ipStr, _, err := net.SplitHostPort(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipStr)
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(interfaces)
	for _, iface := range interfaces {
		fmt.Println("接口名称：", iface.Name)
		// fmt.Println("硬件地址：", iface.HardwareAddr)

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			fmt.Println("IP地址：", addr.String())
		}
	}

}
