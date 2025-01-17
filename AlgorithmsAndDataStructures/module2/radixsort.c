#include <stdio.h>

union Int32 {
    int x;
    unsigned char bytes[4];
};

void DistributionSort(union Int32 *arr, union Int32 *result, int byte_n, int length){
    int count[256] = {0};
    int k = 0;
    for (int j = 0; j < length; j++){
        k = arr[j].bytes[byte_n];
        count[k]++;
    }
    for (int i = 1; i < 256; i++){
        count[i] += count[i - 1];
    }
    int i = 0;
    for (int j = length - 1; j >= 0; j--){
        k = arr[j].bytes[byte_n];
        i = count[k] - 1;
        count[k] = i;
        result[i] = arr[j];
    }
}

void RadixSort(union Int32 *arr, union Int32 *result, int n){
    for (int i = 0; i < 4; i++) {
        DistributionSort(arr, result, i, n);
        for (int j = 0; j < n; j++) {
            arr[j] = result[j];
        }
    }
    int shift = 0;
    while (shift < n && result[shift].x >= 0) {
        shift++;
    }
    union Int32 neg_num[n];
    for (int i = 0; i < n - shift; i++) {
        neg_num[i] = result[shift + i];
    }
    for (int i = 0; i < shift; i++) {
        neg_num[n - shift + i] = result[i];
    }
    for (int i = 0; i < n; i++) {
        result[i] = neg_num[i];
    }
}


int main(){
    int n;
    scanf("%d", &n);
    union Int32 value[n];
    union Int32 result[n];
    for(int i = 0; i < n; i++){
        int x;
        scanf("%d", &x);
        value[i].x = x;
    }
    RadixSort(value, result, n);
    for (int i = 0; i < n; i++){
        printf("%d ", result[i].x);
    }
    return 0;
}