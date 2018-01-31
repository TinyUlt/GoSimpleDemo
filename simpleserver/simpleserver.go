package main

import (
	"fmt"
	"net"
)

const (
	sn = "172.17.0.2:9228"
)

func main() {

	fmt.Println("starting")

	chanMain := make(chan bool)

	ServerConn, err := net.Listen("tcp", sn)
	if err != nil {
		fmt.Println("net.Listen error:", err)
	}

	for {
		conn, err := ServerConn.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("connection")
		go Connection(conn)
	}
	<-chanMain
}

func Connection(conn net.Conn) {

}
