#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int **Delta1(char *S, int size){
    int lenS = strlen(S);
    int **delta = (int **)malloc(lenS * sizeof(int *));
    for (int a = 0; a < lenS; a++){
        delta[a] = (int *)malloc(size * sizeof(int));
        for (int b = 0; b < size; b++){
            delta[a][b] = lenS;
        }
    }
    delta[0][S[0]] = lenS - 1;
    for (int j = 1; j < lenS; j++){
        for (int i = 0; i < size; i++){
            delta[j][i] = delta[j - 1][i];
        }
        delta[j][S[j]] = lenS - j - 1;
    }
    return delta;
}

int BMSubst(char *S, int size, char* T){
    int lenS = strlen(S);
    int lenT = strlen(T);
    int **delta = Delta1(S, size);
    int k = lenS - 1;
    while(k < lenT){
        int i = lenS - 1;
        while(T[k] == S[i]){
            if (i == 0){
                for (int i = 0; i < lenS; i++) free(delta[i]);
                free(delta);
                return k;
            }
            i--;
            k--;
        }
        k += (delta[i][T[k]] > lenS - i) ? delta[i][T[k]] : lenS - i;
    }
    for (int i = 0; i < lenS; i++){
        free(delta[i]);
    }
    free(delta);
    return lenT;
}


int main(int argc, char *argv[]){
    char *S = argv[1];
    char *T = argv[2];
    int ans = BMSubst(S, 160, T);
    printf("%d", ans);
    return 0;
}