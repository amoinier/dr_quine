#include <stdio.h>
#define MAIN() int main(void) {char *str = "#include <stdio.h>%c#define MAIN() int main(void) {char *str = %c%s%c; FILE *fp = fopen(%cGrace_kid.c%c, %cw%c); if (fp) {fprintf(fp, str, 10, 34, str, 34, 34, 34, 34, 34, 10, 10, 10, 10, 10, 10, 10, 10);}}%c#define EXP1 THIS LINE START A DEFINE MAIN%c#define EXP2 THATS ALL, ENVOY!%cMAIN()%c/*%cEXP1%cEXP2%c*/%c"; FILE *fp = fopen("Grace_kid.c", "w"); if (fp) {fprintf(fp, str, 10, 34, str, 34, 34, 34, 34, 34, 10, 10, 10, 10, 10, 10, 10, 10);}}
#define EXP1 THIS LINE START A DEFINE MAIN
#define EXP2 THATS ALL, ENVOY!
MAIN()
/*
EXP1
EXP2
*/
