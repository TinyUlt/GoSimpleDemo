package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

//用户信息
type User struct {
	userName       string
	userAddr       *net.UDPAddr
	userListenConn *net.UDPConn
	chatToConn     *net.UDPConn
}

//服务器监听端口
const LISTENPORT = 1616

//缓冲区
const BUFFSIZE = 1024

var buff = make([]byte, BUFFSIZE)

//在线用户
var onlineUser = make([]User, 0)

//在线状态判断缓冲区
var onlineCheckAddr = make([]*net.UDPAddr, 0)

//错误处理
func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

//消息处理
func HandleMessage(udpListener *net.UDPConn) {

	n, addr, err := udpListener.ReadFromUDP(buff)

	HandleError(err)

	if n > 0 {
		msg := AnalyzeMessage(buff, n)

		switch msg[0] {
		//连接信息
		case "connect  ":
			//获取昵称+端口
			userName := msg[1]
			userListenPort := msg[2]
			//获取用户ip
			ip := AnalyzeMessage([]byte(addr.String()), len(addr.String()))
			//显示登录信息
			fmt.Println(" 昵称:", userName, " 地址:", ip[0], " 用户监听端口:", userListenPort, " 登录成功！")
			//创建对用户的连接，用于消息转发
			userAddr, err := net.ResolveUDPAddr("udp4", ip[0]+":"+userListenPort)
			HandleError(err)

			userConn, err := net.DialUDP("udp4", nil, userAddr)
			HandleError(err)

			//因为连接要持续使用，不能在这里关闭连接
			//defer userConn.Close()
			//添加到在线用户
			onlineUser = append(onlineUser, User{userName, addr, userConn, nil})

		case "online   ":
			//收到心跳包
			onlineCheckAddr = append(onlineCheckAddr, addr)

		case "outline  ":
			//退出消息，未实现
		case "chat     ":
			//会话请求
			//寻找请求对象
			index := -1
			for i := 0; i < len(onlineUser); i++ {
				if onlineUser[i].userName == msg[1] {
					index = i
				}
			}
			//将所请求对象的连接添加到请求者中
			if index != -1 {
				nowUser, _ := FindUser(addr)
				onlineUser[nowUser].chatToConn = onlineUser[index].userListenConn
			}
		case "get      ":
			//向请求者返回在线用户信息
			index, _ := FindUser(addr)
			onlineUser[index].userListenConn.Write([]byte("当前共有" + strconv.Itoa(len(onlineUser)) + "位用户在线"))
			for i, v := range onlineUser {
				onlineUser[index].userListenConn.Write([]byte("" + strconv.Itoa(i+1) + ":" + v.userName))
			}
		default:
			//消息转发
			//获取当前用户
			index, _ := FindUser(addr)
			//获取时间
			nowTime := time.Now()
			nowHour := strconv.Itoa(nowTime.Hour())
			nowMinute := strconv.Itoa(nowTime.Minute())
			nowSecond := strconv.Itoa(nowTime.Second())
			//请求会话对象是否存在
			if onlineUser[index].chatToConn == nil {
				onlineUser[index].userListenConn.Write([]byte("对方不在线"))
			} else {
				onlineUser[index].chatToConn.Write([]byte(onlineUser[index].userName + " " + nowHour + ":" + nowMinute + ":" + nowSecond + "\n" + msg[0]))
			}

		}
	}
}

//消息解析，[]byte -> []string
func AnalyzeMessage(buff []byte, len int) []string {
	analMsg := make([]string, 0)
	strNow := ""
	for i := 0; i < len; i++ {
		if string(buff[i:i+1]) == ":" {
			analMsg = append(analMsg, strNow)
			strNow = ""
		} else {
			strNow += string(buff[i : i+1])
		}
	}
	analMsg = append(analMsg, strNow)
	return analMsg
}

//寻找用户，返回（位置，是否存在）
func FindUser(addr *net.UDPAddr) (int, bool) {
	alreadyhave := false
	index := -1
	for i := 0; i < len(onlineUser); i++ {

		if onlineUser[i].userAddr.String() == addr.String() {
			alreadyhave = true
			index = i
			break
		}
	}
	return index, alreadyhave
}

//处理用户在线信息（暂时仅作删除用户使用）
func HandleOnlineMessage(addr *net.UDPAddr, state bool) {
	index, alreadyhave := FindUser(addr)
	if state == false {
		if alreadyhave {
			onlineUser = append(onlineUser[:index], onlineUser[index+1:len(onlineUser)]...)
		}
	}
}

//在线判断，心跳包处理，每5s查看一次所有已在线用户状态
func OnlineCheck() {
	for {
		onlineCheckAddr = make([]*net.UDPAddr, 0)
		sleepTimer := time.NewTimer(time.Second * 5)
		<-sleepTimer.C
		for i := 0; i < len(onlineUser); i++ {
			haved := false
		FORIN:
			for j := 0; j < len(onlineCheckAddr); j++ {
				if onlineUser[i].userAddr.String() == onlineCheckAddr[j].String() {
					haved = true
					break FORIN
				}
			}
			if !haved {
				fmt.Println(onlineUser[i].userAddr.String() + "退出！")
				HandleOnlineMessage(onlineUser[i].userAddr, false)
				i--
			}

		}
	}
}

func main() {

	//监听地址
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:"+strconv.Itoa(LISTENPORT))
	HandleError(err)
	//监听连接
	udpListener, err := net.ListenUDP("udp4", udpAddr)
	HandleError(err)

	defer udpListener.Close()

	fmt.Println("开始监听：")

	//在线状态判断
	go OnlineCheck()

	for {
		//消息处理
		HandleMessage(udpListener)
	}

}
func test(mode int) {
	// 创建模拟网络：丢包率10%，Rtt 60ms~125ms

	// 创建两个端点的 kcp对象，第一个参数 conv是会话编号，同一个会话需要相同
	// 最后一个是 user参数，用来传递标识
	output1 := func(buf []byte, size int) {

	}
	kcp1 := kcp.NewKCP(0x11223344, output1)

	current := uint32(iclock())
	slap := current + 20
	index := 0
	next := 0
	var sumrtt uint32
	count := 0
	maxrtt := 0

	// 配置窗口大小：平均延迟200ms，每20ms发送一个包，
	// 而考虑到丢包重发，设置最大收发窗口为128
	kcp1.WndSize(128, 128)

	// 判断测试用例的模式
	// 启动快速模式
	// 第二个参数 nodelay-启用以后若干常规加速将启动
	// 第三个参数 interval为内部处理时钟，默认设置为 10ms
	// 第四个参数 resend为快速重传指标，设置为2
	// 第五个参数 为是否禁用常规流控，这里禁止
	kcp1.NoDelay(1, 10, 2, 1)

	buffer := make([]byte, 2000)
	var hr int32

	ts1 := iclock()

	for {
		time.Sleep(1 * time.Millisecond)
		current = uint32(iclock())
		kcp1.Update()

		// 每隔 20ms，kcp1发送数据
		for ; current >= slap; slap += 20 {
			buf := new(bytes.Buffer)
			binary.Write(buf, binary.LittleEndian, uint32(index))
			index++
			binary.Write(buf, binary.LittleEndian, uint32(current))
			// 发送上层协议包
			kcp1.Send(buf.Bytes())
			//println("now", iclock())
		}

		// 处理虚拟网络：检测是否有udp包从p1->p2

		// 处理虚拟网络：检测是否有udp包从p2->p1
		for {
			hr = vnet.recv(0, buffer, 2000)
			if hr < 0 {
				break
			}
			// 如果 p1收到udp，则作为下层协议输入到kcp1
			kcp1.Input(buffer[:hr], true, false)
			//println("@@@@", hr, r)
		}

		// kcp1收到kcp2的回射数据
		for {
			hr = int32(kcp1.Recv(buffer[:10]))
			buf := bytes.NewReader(buffer)
			// 没有收到包就退出
			if hr < 0 {
				break
			}
			var sn uint32
			var ts, rtt uint32
			binary.Read(buf, binary.LittleEndian, &sn)
			binary.Read(buf, binary.LittleEndian, &ts)
			rtt = uint32(current) - ts

			if sn != uint32(next) {
				// 如果收到的包不连续
				//for i:=0;i<8 ;i++ {
				//println("---", i, buffer[i])
				//}
				println("ERROR sn ", count, "<->", next, sn)
				return
			}

			next++
			sumrtt += rtt
			count++
			if rtt > uint32(maxrtt) {
				maxrtt = int(rtt)
			}

			println("[RECV] mode=", mode, " sn=", sn, " rtt=", rtt)
		}

		if next > 100 {
			break
		}
	}
}

func iclock() int32 {
	return int32((time.Now().UnixNano() / 1000000) & 0xffffffff)
}

/*
func main(){

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
	return


}
*/
