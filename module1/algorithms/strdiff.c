#include <stdio.h>

int strdiff(char *a, char *b) {
    int bitIndex = 0;

    while (a != "\0" || b != "\0"){
        unsigned char charA = *a;
        unsigned char charB = *b;
        printf("%c", charA);
        if (*a != "\0") a++;
    }

    return 1;
}

int main(){
    char *a = "aa";
    char *b = "ai";

    int x = strdiff(a, b);
    
    return 0;
}