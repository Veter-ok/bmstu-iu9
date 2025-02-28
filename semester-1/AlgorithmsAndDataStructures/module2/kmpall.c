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

void KMPSubst(char *S, char *T){
    int lenS = strlen(S);
    int lenT = strlen(T);
    int *p = Prefix(S);
    int q = 0;
    for (int k = 0; k < lenT; k++){
        while (q > 0 && S[q] != T[k]){
            q = p[q - 1];
        }  
        if (S[q] == T[k]){
            q++;
        }
        if (q == lenS){
            printf("%d ", k - lenS + 1);
            q = p[q - 1];
        }
    }
    free(p);
}


int main(int argc, char *argv[]){
    char *S = argv[1];
    char *T = argv[2];
    KMPSubst(S, T);
    return 0;
}