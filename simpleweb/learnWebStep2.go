package main

import (
	"fmt"
	"html/template"
	"net/http" //搭建web服务很重要的一个包
)

func Index(w http.ResponseWriter, r *http.Request) {
	//解析指定模板文件index.html
	data := make(map[string]string)
	t, _ := template.ParseFiles("index.html")
	data["Name"] = "BCL"
	t.Execute(w, data)
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
 */
