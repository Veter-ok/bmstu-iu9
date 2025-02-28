#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int* Prefix(char *S){
    int length = strlen(S);
    int *p = (int*)malloc(length * sizeof(int));;
    for (int i = 0; i < length; i++){
        p[i] = 0;
    }
    int t = 0;
    for (int i = 1; i < length; i++){
        while (t > 0 && S[t] != S[i]) {
            t = p[t - 1];
        }
        if (S[t] == S[i]){
            t++;
        }
        p[i] = t;
    }
    return p;
}

int main(int argc, char *argv[]) {
    char *S = argv[1];
    int len = strlen(S);
    int *pi = Prefix(S); 
    for (int i = 1; i < len; i++) {
        int period_len = (i + 1) - pi[i];
        int repeat = (i + 1) / period_len; 
        if ((i + 1) % period_len == 0 && repeat != 1) {
            printf("%d %d\n", i + 1, repeat); 
        }
    }
    free(pi);
    return 0;
}