#include <stdio.h>


int main() {
    long long unsigned a, b, m;
    long long unsigned result = 0;
    long long unsigned cof[64];
    scanf("%llu %llu %llu", &a, &b, &m);

    for (int i = 63; i >= 0; i--){
        cof[i] = b % 2;
        b /= 2;
    }

    for (int i = 0; i < 63; i++){
        result = (result + a * cof[i]) % m * 2 % m;
    }
    result = (result + a * cof[63]) % m;
    printf("%llu", result);
}