package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "host")
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	var service  = "www.baidu.com:80"

	//conn, err := net.Dial("ip4:icmp", service)
	//Fatal error : write ip4 192.168.0.119->110.242.68.4: wsasend: An attempt was made to access a socket in a way forbidden by its access permissions.
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}


func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for true {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}