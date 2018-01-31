package main

import "github.com/kataras/iris"

func main() {
	iris.Config.IsDevelopment = true // this will reload the templates on each request
	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.MustRender("hi.html", struct{ Name string }{Name: "iris"})
}
