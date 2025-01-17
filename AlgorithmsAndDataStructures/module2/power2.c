#include <stdio.h>

int combinationUtil(int arr[], int data[], int start, int end, int index, int r) {
    int ans = 0;
    if (index == r) {
        int sum = 0;
        for (int i = 0; i < r; i++) sum += data[i]; 
        if (sum && !(sum & (sum - 1))) return 1;
        return 0;
    }
    for (int i=start; (i <= end && end - i + 1 >= r - index); i++) {
        data[index] = arr[i];
        ans += combinationUtil(arr, data, i+1, end, index+1, r);
    }
    return ans;
}

int main() {
    int n;
    int ans = 0;
    scanf("%d", &n);
    int numbers[n];
    for (int i = 0; i < n; i++) scanf("%d", &numbers[i]);
    for (int i = 1; i <= n; i++){
        int data[i];
        ans += combinationUtil(numbers, data, 0, n-1, 0, i);
    }
    printf("%d", ans);
}