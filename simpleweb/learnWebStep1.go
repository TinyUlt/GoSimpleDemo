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

/*
我们来简单的了解一下，上面的代码都做了哪些事。
至于golang的语法，这里就不展开细讲了，不熟悉的朋友建议先把语法好好看一看。

首先来看 http.HandleFunc(“/”, Index) 这句，这句的意思就是注册路由，当客户端请求根页面的时候，交由Index函数去处理。
函数声明如下：

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

第一个参数 pattern 就是要匹配的路由，比如 “/” 用来匹配首页，
第二个参数是一个函数func(ResponseWriter, *Request)
所以在写自定义处理函数的时候得按照这种格式：
func Index(w http.ResponseWriter, r *http.Request)
这个函数将会由标准库去调用。
参数列表中的w是用于向客户端发送数据的而r是从客户端接收到的数据，服务器与客户端的所有交互都要依赖这两个参数，想深入了解的朋友可以去查看手册，这里就不赘述了。

接下来就是监听端口8080端口，并启动服务，golang为我们实现了一个完整的服务器，就像apache和nginx了，当然了功能肯定没有那么强大，很多东西需要我们自己来实现，所以编译好的go web程序并不依赖apache或者nginx，大大方便了程序的部署。

虽然只有短短的几行代码，但却实现了高并发，每个请求都是通过独立的goroutine去服务的，互不干扰。
*/
