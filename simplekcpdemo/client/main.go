package main

import (
	"fmt"
	"net"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

func main() {

	//c := make(chan bool)
	sip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	output1 := func(buf []byte, size int) {
		conn.Write(buf[:size])
	}
	kcp1 := kcp.NewKCP(0x11223344, output1)

	kcp1.WndSize(128, 128)
	kcp1.NoDelay(1, 10, 2, 1)
	//kcp1.Input(buffer[:hr], true, false)
	buffer := make([]byte, 20)
	data := make([]byte, 1024)
	//	for {
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			kcp1.Update()
			for {
				hr := int32(kcp1.Recv(buffer))
				//				fmt.Println("ttttttttt")
				// 没有收到包就退出
				if hr < 0 {
					break
				}
				fmt.Printf("%s", buffer[:5])
				//				kcp1.Send([]byte("world"))

			}

		}
	}()
	for {
		kcp1.Send([]byte("hello"))

		n, _ := conn.Read(data)

		kcp1.Input(data[:n], true, false)
	}

}

/*
func main() {
	sip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello"))
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
}
*/
