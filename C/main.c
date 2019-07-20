#include <stdio.h>

extern char* Version();

int main() {
   printf("Version() --> '%s'\n", Version());
   return 0;
}
