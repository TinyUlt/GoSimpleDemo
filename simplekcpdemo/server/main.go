package main

import (
	"fmt"
	"net"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())

	var RA *net.UDPAddr
	output1 := func(buf []byte, size int) {

		listener.WriteToUDP(buf[:size], RA)

	}
	kcp1 := kcp.NewKCP(0x11223344, output1)
	kcp1.WndSize(128, 128)
	kcp1.NoDelay(1, 10, 2, 1)
	buffer := make([]byte, 20)
	data := make([]byte, 1024)
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)

			kcp1.Update()

			for {
				hr := int32(kcp1.Recv(buffer))
				// 没有收到包就退出
				if hr < 0 {
					break
				}
				fmt.Printf("%s", buffer[:5])
				kcp1.Send([]byte("world"))
			}
		}

	}()

	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		RA = remoteAddr
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}

		kcp1.Input(data[:n], true, false)
	}
}

/*
func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}

		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)

		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}
*/
