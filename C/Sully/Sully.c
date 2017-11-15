int i = 5;
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
int main(void) {
if (i >= 0) {
if (!strstr(__FILE__, "Sully.c")) {
i--;
}
char str[70];
sprintf(str, "Sully_%1$d.c", i);
FILE *fp = fopen(str, "w");
if (fp) {
char *quine = "int i = %4$d;%1$c#include <stdio.h>%1$c#include <string.h>%1$c#include <stdlib.h>%1$cint main(void) {%1$cif (i >= 0) {%1$cif (!strstr(__FILE__, %2$cSully.c%2$c)) {%1$ci--;%1$c}%1$cchar str[70];%1$csprintf(str, %2$cSully_%5$c1$d.c%2$c, i);%1$cFILE *fp = fopen(str, %2$cw%2$c);%1$cif (fp) {%1$cchar *quine = %2$c%3$s%2$c;%1$cfprintf(fp, quine, 10, 34, quine, i, 37);%1$cfclose(fp);%1$csprintf(str, %2$cgcc Sully_%5$c1$d.c -o Sully_%5$c1$d%2$c, i);%1$csystem(str);%1$c}%1$cif (i > 0) {%1$csprintf(str, %2$c./Sully_%5$c1$d%2$c, i);%1$csystem(str);%1$c}%1$c}%1$creturn 0;%1$c}%1$c";
fprintf(fp, quine, 10, 34, quine, i, 37);
fclose(fp);
sprintf(str, "gcc Sully_%1$d.c -o Sully_%1$d", i);
system(str);
}
if (i > 0) {
sprintf(str, "./Sully_%1$d", i);
system(str);
}
}
return 0;
}
