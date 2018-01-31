package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func CallerName(start int, skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip + start); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}
func BuyuLog(start int, skip int, format string, a ...interface{}) string {

	str := "\n"
	s := fmt.Sprintf(format, a...)
	str = fmt.Sprintf("%s%-4v %-103v \n", str, "#", s)

	for i := 0; i < skip; i++ {
		name, file, line, ok := CallerName(start, i)
		if !ok {
			break
		}
		str = fmt.Sprintf("%s%-5v%-5v %-140v %-30v %-6d |\n", str, "@", i, name, filepath.Base(file), line)
	}
	return str
}
func test() {
	fmt.Print(BuyuLog(2, 4, "mylog1112233dfgsheudughedugherufhiufhesdfgsdffsrfsferfethrrfgrtgrfrt%d", 111))

}
func test2() {
	test()
}
func main() {
	//	file, _ := exec.LookPath(os.Args[0])
	//	filepath.
	//	fmt.Println(filepath.Dir(file))
	//	fmt.Println(filepath.Base(file))
	test2()
	test2()
	// Output:
	// skip = 0
	//   file = caller.go, line = 19
	//   name = main.main
	// skip = 1
	//   file = $(GOROOT)/src/pkg/runtime/proc.c, line = 220
	//   name = runtime.main
	// skip = 2
	//   file = $(GOROOT)/src/pkg/runtime/proc.c, line = 1394
	//   name = runtime.goexit
}
