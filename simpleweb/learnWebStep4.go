package main

import (
	"fmt"
	"html/template"
	"net/http" //搭建web服务很重要的一个包
	"regexp"
)

type MyMux struct {
	routers map[string]func(http.ResponseWriter, *http.Request)
}

//这个示例函数，将传进来的字符串用*****包起来
func Index(w http.ResponseWriter, r *http.Request) {
	//用于保存数据的map
	data := make(map[string]string)
	tempfunc := make(template.FuncMap)
	tempfunc["showname"] = ShowName
	//得给模板起个名字才行
	t := template.New("index2.html")
	t = t.Funcs(tempfunc)
	t, _ = t.ParseFiles("./index2.html")
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

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHttp")
	//遍历routers，寻找匹配的path
	for path, f := range p.routers {
		if ok, _ := regexp.MatchString("^"+path+"$", r.URL.Path); ok {
			fmt.Println(r.URL.Path)
			//			fmt.Println(f)
			f(w, r)
			return
		}
	}
	fmt.Fprintf(w, "Error: Don't match URL '%s'", r.URL.Path)
}
func Static(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deal Static: ", r.URL.Path)
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "."+r.URL.Path)
}
func main() {
	mux := &MyMux{}
	mux.routers = make(map[string]func(http.ResponseWriter, *http.Request))
	mux.routers["/"] = Index
	mux.routers["/static/.+"] = Static
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

/*
 */
