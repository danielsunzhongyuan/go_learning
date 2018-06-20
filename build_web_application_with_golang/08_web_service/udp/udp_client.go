package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	// 解析 IP:Port
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	fmt.Println("udpAddr:", udpAddr)
	checkErrorForClient(err)

	// 建立连接
	conn, err := net.DialUDP("udp", nil, udpAddr)
	fmt.Println("conn:", conn)
	checkErrorForClient(err)

	// 发送请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErrorForClient(err)

	// 读取结果（这里读取的方式先打一个问号？）
	result, err := ioutil.ReadAll(conn)

	fmt.Println("result:", string(result))

	// 这次在检查错误的时候，报错并退出了，但上面的结果打印是没问题的。错误信息：
	// read udp 127.0.0.1:61588->127.0.0.1:7777: read: connection reset by peerexit status 1
	// 一种说法是上面的 ReadAll 的这种方式是不合适的，先放在这里，回头再看。
	checkErrorForClient(err)
	os.Exit(0)
}

func checkErrorForClient(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
