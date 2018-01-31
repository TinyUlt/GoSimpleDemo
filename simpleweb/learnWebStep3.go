package main

import (
	"fmt"
	"html/template"
	"net/http" //搭建web服务很重要的一个包
)

//这个示例函数，将传进来的字符串用*****包起来
func Index(w http.ResponseWriter, r *http.Request) {
	//用于保存数据的map
	data := make(map[string]string)
	tempfunc := make(template.FuncMap)
	tempfunc["showname"] = ShowName
	//得给模板起个名字才行
	t := template.New("index.html")
	t = t.Funcs(tempfunc)
	t, _ = t.ParseFiles("./index.html")
	data["Name"] = "BCL"
	t.Execute(w, data)
}

//这个示例函数，将传进来的字符串用*****包起来
func ShowName(args ...interface{}) string {
	//这里只考虑一个参数的情况
	var str string = ""
	if s, ok := args[0].(string); ok {
		str = "*****" + s + "*****"
	} else {
		str = "Nothing"
	}
	return str
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
