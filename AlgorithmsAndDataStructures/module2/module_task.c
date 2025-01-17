#include <stdio.h>
#include <string.h>

#define R 4294967296
#define Q 29

unsigned long power(unsigned long a, unsigned long b){
    unsigned long res = 1;
    for (unsigned int i = 0; i < b; i++){
        res *= a;
    }
    return res;
}

unsigned long hash(char *str, int len) {
    unsigned long result = 0;
    for (int i = 0; i < len; i++) {
        result = result * Q + (unsigned char)str[i];
    }
    return result % R;
}

unsigned long updateHash(unsigned long hashVal, char str1, char str2, unsigned long pow){
    unsigned long res = (hashVal * Q - (unsigned char)str1 * pow + (unsigned char)str2) % R;
    return res;
}

void RKSubst(char *S, char *T){
    int k = 0;
    int lenS = strlen(S);
    int lenT = strlen(T);
    if (lenS > lenT){
        printf("\n");
        return;
    }
    unsigned long constPower = power(Q, lenS);
    unsigned long hs = hash(S, lenS);
    unsigned long ht = hash(T, lenS);
    for (; k < lenT - lenS + 1; k++){
        if (hs == ht && strncmp(S, T + k, lenS) == 0) {
            printf("%d ", k);
        }
        ht = updateHash(ht, T[k], T[k + lenS], constPower);
    } 
    printf("\n");
}

int main(int argc, char *argv[]){
    char *S = argv[1];
    char *T = argv[2];
    RKSubst(S, T);
    return 0;
}