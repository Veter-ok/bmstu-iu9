#include <stdio.h>

int main(){
    long long n, x0;
    scanf("%lld %lld", &n, &x0);
    long long P[n+1]; 
    for (int i = 0; i <= n; i++) {
        scanf("%lld", &P[i]);
    }

    long long valPol = P[0];
    for (int i = 1; i <= n; i++){
        valPol = valPol * x0 + P[i];
    }

    long long valDerPol = P[0] * n;
    for (int i = 1; i < n; i++){
        valDerPol = valDerPol * x0 + (P[i] * (n - i));
    }

    printf("%lld %lld", valPol, valDerPol);
    return 0;
}