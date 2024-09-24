#include <stdio.h>
#include <math.h>

int main() {
    long long int x;
    long long unsigned maxNumber = 0;
    scanf("%llu", &x);
    long long unsigned moduleX = x >= 0 ? x : -x;
    long long unsigned length = (int)sqrt(moduleX)+1; 
    unsigned primeNumbers[length];
    for (int i = 0; i < length; i++){
        primeNumbers[i] = 1;
    }

    for (long long unsigned p = 2; pow(p, 2) <= length; p++){
        for (int i = 2*p; i <= length; i += p){
            primeNumbers[i - 2] = 0;
        }
    }

    for (int i = 0; i < length; i++){
        if (primeNumbers[i] && moduleX % (i + 2) == 0){
            maxNumber = (i + 2);
            while (moduleX % (i + 2) == 0){
                moduleX /= (i + 2);
            }
        }
    }

    if (maxNumber < moduleX){
        printf("%llu", moduleX);
    }else{
        printf("%llu", maxNumber);
    }
}