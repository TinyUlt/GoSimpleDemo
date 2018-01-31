//package main

/*
func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/tinyult/.golang/simpledemo/simpleweb/static/")))
	http.ListenAndServe(":8080", nil)
}
*/
/*
import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	//\"encoding/json\"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers   []Server
	ServersID string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9002", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("begin")
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hi, I love you %s", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
	}
	fmt.Println("end")
}
*/
/*
package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
func main() {
	http.HandleFunc("/hello", hello)
	http.Handle("/handle", http.HandlerFunc(say))
	http.ListenAndServe(":8001", nil)
	select {} //阻塞进程
}
*/
package main

import (
	"fmt"
	"net/http" //搭建web服务很重要的一个包
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, this is my first page!")
}

func main() {
	http.HandleFunc("/", Index)
	// 监听本机的8080端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
