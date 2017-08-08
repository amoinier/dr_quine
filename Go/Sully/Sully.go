package main
import "fmt"
import "io/ioutil"
import "runtime"
import "strings"
import "os/exec"
import "strconv"
var i = 5;
func main() { if (i >= 0) { _, fileName, _, _ := runtime.Caller(1); fmt.Printf("%s\n", fileName); if (strings.Index(fileName, strconv.Itoa(i)) != -1) { i--; }
c := strconv.Itoa(i); str := "Sully_" + c + ".go"; str2 := "Sully_" + c; var quine = "package main%cimport %cfmt%c%cimport %cio/ioutil%c%cimport %cruntime%c%cimport %cstrings%c%cimport %cos/exec%c%cimport %cstrconv%c%cvar i = %d;%cfunc main() { if (i >= 0) { _, fileName, _, _ := runtime.Caller(1); if (strings.Index(fileName, strconv.Itoa(i)) != -1) { i--; }%cc := strconv.Itoa(i); str := %cSully_%c + c + %c.go%c; str2 := %cSully_%c + c; var quine = %c%s%c; var ctn = fmt.Sprintf(quine, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, i, 10, 10, 34, 34, 34, 34, 34, 34, 34, quine, 34, 34, 34, 34, 34, 34, 34, 10, 10); ioutil.WriteFile(str, []byte(ctn), 0644); exec.Command(%cgo build -o %c+str2+%c %c+str); if (i > 0) { exec.Command(%c./%c + str2); }}%creturn;}%c"; var ctn = fmt.Sprintf(quine, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, 34, 34, 10, i, 10, 10, 34, 34, 34, 34, 34, 34, 34, quine, 34, 34, 34, 34, 34, 34, 34, 10, 10); ioutil.WriteFile(str, []byte(ctn), 0644); exec.Command("go build -o "+str2+" "+str); if (i > 0) { exec.Command("./" + str2); }}
return;}
