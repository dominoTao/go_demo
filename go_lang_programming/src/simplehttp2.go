package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "host")
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	service := "www.baidu.com:80"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError1(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError1(err)
	result, err := ioutil.ReadAll(conn)
	checkError1(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}