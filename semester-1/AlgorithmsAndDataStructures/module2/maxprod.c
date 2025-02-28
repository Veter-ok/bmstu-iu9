#include <stdio.h>

int main(){
    int n;
    int l = 0;
    int r = 0;
    scanf("%d", &n);
    double numbers[n];
    for (int i = 0; i < n; i++){
        int a, b;
        scanf("%d/%d", &a, &b);
        numbers[i] = (double)a / b;
    }
    double cur_mul = 1;
    double ans = 0;
    int start = 0;
    for (int i = 0; i < n; i++){
        cur_mul *= numbers[i];
        if (cur_mul > ans){
            ans = cur_mul;
            l = start;
            r = i;
        }
        if (cur_mul < 1) {
            cur_mul = 1.0;
            start = i + 1;
        }
    }
    printf("%d %d", l, r);
    return 0;
}