#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main() {
    long long int xx;
    long long unsigned maxNumber = 0;
    scanf("%llu", &xx);
    long long unsigned x = xx >= 0 ? xx : -xx;
    long long unsigned copyX = x;
    long long unsigned length = (int)sqrt(x)+1; 
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
        if (primeNumbers[i] && copyX % (i + 2) == 0){
            maxNumber = (i + 2);
            while (copyX % (i + 2) == 0){
                copyX /= (i + 2);
            }
        }
    }

    if (maxNumber < copyX){
        printf("%llu", copyX);
    }else{
        printf("%llu", maxNumber);
    }
}