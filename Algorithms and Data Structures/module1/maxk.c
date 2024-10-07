#include <stdio.h>

int main(){
    int n, k;
    scanf("%d", &n);
    int numbers[n];
    for (int i = 0; i < n; i++){
        scanf("%d", &numbers[i]);
    } 
    scanf("%d", &k);
    int sum = 0;
    for (int i = 0; i < k; i++){
        sum += numbers[i];
    }
    int ans = sum;
    for (int i = k; i < n; i++){
        sum = sum - numbers[i - k] + numbers[i];
        if (sum > ans){
            ans = sum;
        }
    }
    printf("%d", ans);
}