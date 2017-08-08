package main
import "fmt"
import "io/ioutil"
func main() {var str = "package main%cimport %cfmt%c%cimport %cio/ioutil%c%cfunc main() {var str = %c%s%c; var ctn = fmt.Sprintf(str, 10, 34, 34, 10, 34, 34, 10, 34, str, 34, 34, 34, 10, 10, 9, 10, 9, 10, 10); ioutil.WriteFile(%cGrace_kid.go%c, []byte(ctn), 0644);}%c/*%c%cEXP1%c%cEXP2%c*/%c"; var ctn = fmt.Sprintf(str, 10, 34, 34, 10, 34, 34, 10, 34, str, 34, 34, 34, 10, 10, 9, 10, 9, 10, 10); ioutil.WriteFile("Grace_kid.go", []byte(ctn), 0644);}
/*
	EXP1
	EXP2
*/
