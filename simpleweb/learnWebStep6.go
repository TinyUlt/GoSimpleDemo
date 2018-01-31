package main

import (
	"log"
	"net/http"
)

func main() {
	RouterBinding()                          // 路由绑定函数
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func RouterBinding() {

	//	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./ui"))))

	http.Handle("/", http.FileServer(http.Dir("./ui/")))
	//	http.HandleFunc("/images/", fileUpload.DownloadPictureAction)
}
