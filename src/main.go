package main

import (
	routerBase "./router"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// 基础提示信息输出

func BaseServer() {
	// 输出启动方式信息
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("获取ip出错，错误信息：%v", err)
		os.Exit(1)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				fmt.Println("ip:", ip)
			}
		}
	}
	info := `
    ***************************************************************

    欢迎使用LRP服务, 使用GO启动PerformancePlatform/main.go

    您的IP地址: %s
    您的端口号: %v

    Quit the server with Control-C

    ***************************************************************
	`
	fmt.Printf(info, ip, 9090)
}

//程序的主入口

func main() {

	BaseServer()

	router := routerBase.NewRouter()
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}