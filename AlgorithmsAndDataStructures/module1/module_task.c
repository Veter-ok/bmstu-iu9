#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]){
    char *x = argv[1];
    int len = strlen(x) - 1;
    if (x[0] == '1' && len == 0){
        printf("10");
        return 0;
    }

    if (x[len] == '0'){
        x[len] = '1';
    }else if (x[len] == '1' && x[len - 1] == '0'){
        x[len] = '0';
        x[len - 1] = '1';
    }
    for (int i = len; i >= 3; i--){
        if (x[i] == '1' && x[i-1] == '1' && x[i-2] == '0'){
            x[i] = '0';
            x[i-1] = '0';
            x[i-2] = '1';
        }
    }
    if (x[1] == '1' && x[0] == '1'){
        x[0] = '0';
        x[1] = '0';
        printf("1");
    }
    printf("%s", x);
}