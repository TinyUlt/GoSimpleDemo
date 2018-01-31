package main

import (
	"gt_msg"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

const (
	sn = "114.55.234.70:9228"
)

func main() {

	fname := "/home/tinyult/.golang/simpleprotobuf/filepath.json"
	log.Println("starting")
	/*
		p := tinypath.Person{
			Id:    123,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*tinypath.Person_PhoneNumber{
				{Number: "555-4321", Type: 1},
			},
		}

		book := &tinypath.AddressBook{}

		book.People = append(book.People, &p)
		// ...
	*/
	book := &gt_msg.CurrentFrame{
		Frame: 100,
	} //  tinypath.AddressBook{}
	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book = &gt_msg.CurrentFrame{} //  tinypath.AddressBook{}
	var msg proto.Message = book
	if err := proto.Unmarshal(in, msg); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	} else {

		log.Println(book.Frame)
		log.Println("end")
	}
}

/*
func main() {

	fmt.Println("starting")

	chanMain := make(chan bool)

	utils.InitProtoTool(gt_msg.FishMsg)

	fmt.Println("initPrototool ok")

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

	NetBuf := utils.NewNetBuffer([]byte{69, 123, 132, 104, 67, 95, 33, 74, 120, 131, 61, 101, 55, 101, 69, 44})

	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 13)) //50秒没有心跳就超时
		if err := NetBuf.Read(conn); err != nil {
			fmt.Println(err)
			break
		}
		for {
			if version, msg := NetBuf.GetAMsg(); msg != nil {
				fmt.Println(version)

				fmt.Println(utils.ProtoToString(msg))

				resp := Demux(msg)

				if resp != nil {
					buf := utils.ToData(resp, NetBuf.Key, 1)
					fmt.Println(len(buf))
					_, err := conn.Write(buf)
					if err != nil {
						fmt.Println(err)
					}

					break
				}

			} else {

				break
			}

			time.Sleep(time.Millisecond * 1)
		}
	}

	fmt.Println("connection end")

}

func Demux(req proto.Message) proto.Message {
	switch req.(type) {

	case *gt_msg.HHRequest:
		fmt.Println("enter")
		resp := new(gt_msg.HHResponse)

		resp.ServerTimeNow = proto.Int(int(time.Now().Unix()))
		return resp
	default:
		return nil
	}

}
*/
