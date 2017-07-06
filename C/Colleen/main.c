#include <stdio.h>
/*
FT_PRINTMYNAME
*/
char *ft_printmyname(void) {return("Alex");}
int main(void) {
/*
str to write in printf
*/
	char *str = "#include <stdio.h>%c/*%cFT_PRINTMYNAME%c*/%cchar *ft_printmyname(void) {return(%cAlex%c);}%cint main(void) {%c/*%cstr to write in printf%c*/%c%cchar *str = %c%s%c;%c%cprintf(str, 10, 10, 10, 10, 34, 34, 10, 10, 10, 10, 10, 9, 34, str, 34, 10, 9, 10, 9, 10, 10);%c%cft_printmyname();%c}%c";
	printf(str, 10, 10, 10, 10, 34, 34, 10, 10, 10, 10, 10, 9, 34, str, 34, 10, 9, 10, 9, 10, 10);
	ft_printmyname();
}
