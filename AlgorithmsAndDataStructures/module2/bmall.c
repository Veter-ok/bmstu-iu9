#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *Delta1(char *S, int size){
    int *sigma = (int *)malloc(size * sizeof(int));
    int lenS = strlen(S);
    for (int a = 0; a < size; a++){
        sigma[a] = lenS;
    } 
    for (int j = 0; j < lenS; j++){
        sigma[S[j]] = lenS - j - 1;
    }
    return sigma;
}

int *Suffix(char *S){
    int lenS = strlen(S);
    int *delta = (int *)malloc(lenS * sizeof(int));
    delta[lenS - 1] = lenS - 1;
    int t = lenS - 1;
    for (int i = lenS - 2; i >= 0; i--){
        while (t < lenS - 1 && S[t] != S[i]) {
            t = delta[t + 1];
        }
        if (S[t] == S[i]){
            t--;
        }
        delta[i] = t;
    }
    return delta;
}

int *Delta2(char *S){
    int lenS = strlen(S);
    int *sigma = (int *)malloc(lenS * sizeof(int));
    int *delta = Suffix(S);
    int t = delta[0];
    for (int i = 0; i < lenS; i++){
        while (t < i){
            t = delta[t + 1];
        }
        sigma[i] = -i + t + lenS;
    }
    for (int i = 0; i < lenS - 1; i++){
        t = i;
        while (t < lenS - 1) {
            t = delta[t + 1];
            if (S[i] != S[t]){
                sigma[t] = -(i + 1) + lenS;
            }
        }
    }
    free(delta);
    return sigma;
}

void BMSubst(char *S, int size, char* T){
    int lenS = strlen(S);
    int lenT = strlen(T);
    int *sigma1 = Delta1(S, size);
    int *sigma2 = Delta2(S);
    int k = lenS - 1;
    while(k < lenT){
        int i = lenS - 1;
        int find = 0;
        while(T[k] == S[i]){
            if (i == 0){
                find = 1;
                printf("%d ", k);
                k += lenS;
                break;
            }
            i--;
            k--;
        }
        if (find){
            continue;
        }
        k += (sigma1[T[k]] >= sigma2[i]) ? sigma1[T[k]] : sigma2[i];
    }
    free(sigma1);
    free(sigma2);
}


int main(int argc, char *argv[]){
    char *S = argv[1];
    char *T = argv[2];
    BMSubst(S, 256, T);
    return 0;
}