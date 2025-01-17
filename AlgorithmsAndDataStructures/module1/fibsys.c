#include <stdio.h>

int main(){
    unsigned maxFibLen = 92;
    long long unsigned x;
    long long unsigned fibonnaci[maxFibLen];
    long long unsigned result[maxFibLen];
    long long unsigned firstBit = 0;
    scanf("%lld", &x);

    fibonnaci[0] = 1; fibonnaci[1] = 2;
    result[0] = 0;    result[1] = 0; 
    for (int i = 2; i < maxFibLen; i++){
        result[i] = 0;
        fibonnaci[i] = fibonnaci[i-1] + fibonnaci[i-2];
    }

    for (int i = maxFibLen - 1; i >= 0; i--){
        if (fibonnaci[i] <= x){
            result[i] = 1;
            firstBit = (i > firstBit) ? i : firstBit;
            x -= fibonnaci[i];
        }
    }

    for (int i = firstBit; i >= 0; i--){
        printf("%llu", result[i]);
    }
}