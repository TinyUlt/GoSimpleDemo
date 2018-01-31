package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2014, 1, 7, 5, 59, 25, 0, time.Local)
	n := t.Unix()
	fmt.Println(n)
}
