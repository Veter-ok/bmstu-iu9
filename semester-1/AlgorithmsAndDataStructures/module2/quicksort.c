#include <stdio.h>
#include <stdlib.h>

void swap(int *a, int *b){
    int copyA = *a;
    *a = *b;
    *b = copyA;
}

void selectionSort(int *array, int low, int high){
    for (int j = high; j > low; j--){
        int k = j;
        for (int i = j - 1; i >= low; i--){
            if (array[k] < array[i]){
                k = i;
            }
        }
        swap(&array[j], &array[k]);
    }
}

int partition(int *array, int low, int high){
    int i = low;
    for (int j = low; j < high; j++) {
        if (array[j] < array[high]){
            swap(&array[i], &array[j]);
            i++;
        }
    }
    swap(&array[i], &array[high]);
    return i;
}

void quicksort(int *arr, int low, int high, int m){
    while (low < high){
        if (high - low + 1 <= m) {
            selectionSort(arr, low, high);
            break;
        }else{
            int q = partition(arr, low, high);
            if (q - low < high - q){
                quicksort(arr, low, q - 1, m);
                low = q + 1;
            }else{
                quicksort(arr, q + 1, high, m);
                high = q - 1;
            }
        }
    }
}

int main(){
    int n, m;
    scanf("%d %d", &n, &m);

    int *arr = (int *)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }

    quicksort(arr, 0, n - 1, m);

    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    free(arr);
    return 0;
}