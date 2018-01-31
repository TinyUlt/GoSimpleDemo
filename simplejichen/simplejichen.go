package main

import "fmt"

type inter interface {
	function()
}

type innerS struct {
	in1 int
	in2 int
}

func (this *innerS) Init() {

	this.in1 = 1
}
func (this *innerS) function() {

	fmt.Println("innerS function")
}
func (this *innerS) function2() {

	fmt.Println("innerS function2", this.in1)
}
func (this *innerS) function3() {
	fmt.Println("innerS function3")
	this.function()
}

type outerS struct {
	innerS
	in1 int
}

func (this *outerS) function() {
	fmt.Println("outerS function", this.in1)
	this.innerS.function()
}

func main() {

	outer := new(outerS)
	outer.function3()
	/*	outer.Init()
		outer.function()
		outer.function2()
	*/
	var interf inter = outer
	interf.function()

	inner := new(innerS)
	interf = inner
	interf.function()
}
