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


int main(int argc, char *argv[]){
    char *S = argv[1];
    char *T = argv[2];
    int lenS = strlen(S);
    int lenT = strlen(T);
    int *pi = Prefix(S);
    int max_index = 0;
    for (int i = 0; i < lenT; i++) {
        if (max_index == lenS) {
            max_index = pi[max_index - 1];
        }
        if (S[max_index] == T[i]) {
            max_index++;
        }else{
            if (max_index == 0){
                printf("no");
                free(pi);
                return 0;
            }
            while (max_index != 0 && S[max_index] != T[i]) {
                max_index = pi[max_index - 1];
            }
            i--;
        }
    }
    printf("yes");
    free(pi);
    return 0;
}