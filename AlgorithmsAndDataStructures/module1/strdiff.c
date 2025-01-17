#include <stdio.h>

int strdiff(char *a, char *b) {
    int bitIndex = 0;
    while(*a != 0 && *b != 0){
        for (int i = 0; i < 8; i++){
            if ((*a & (1 << i)) != (*b & (1 << i))){
                return bitIndex + i;
            }
        }
        bitIndex += 8;
        a++;
        b++;
    }
    if (*a != *b){
        for (int i = 0; i < 8; i++){
            if ((*a & (1 << i)) != (*b & (1 << i))){
                return bitIndex + i;
            }
        }
    }
    return -1;
}

int main(){
    char *a = "The quick brown fox jumps over the lazy do";
    char *b = "T";
    int response = strdiff(a, b);
    printf("%d", response);
    return 0;
}