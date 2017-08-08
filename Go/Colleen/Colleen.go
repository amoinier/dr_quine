package main
/*
	FT_PRINTMYNAME
*/
import "fmt"
func ft_printmyname() string {return("Alex");}
func main() {
/*
	str to write in printf
*/
	var str = "package main%c/*%c%cFT_PRINTMYNAME%c*/%cimport %cfmt%c%cfunc ft_printmyname() string {return(%cAlex%c);}%cfunc main() {%c/*%c%cstr to write in printf%c*/%c%cvar str = %c%s%c;%c%cfmt.Printf(str, 10, 10, 9, 10, 10, 34, 34, 10, 34, 34, 10, 10, 10, 9, 10, 10, 9, 34, str, 34, 10, 9, 10, 9, 10, 10);%c%cft_printmyname();%c}%c";
	fmt.Printf(str, 10, 10, 9, 10, 10, 34, 34, 10, 34, 34, 10, 10, 10, 9, 10, 10, 9, 34, str, 34, 10, 9, 10, 9, 10, 10);
	ft_printmyname();
}
