#include <stdio.h>
#include <stdlib.h>

void insertionSort(int* array, int startIndex, int endIndex){
    int i = startIndex + 1;
    while (i < endIndex) {
        int element = array[i]; 
        int loc = i - 1;
        while (loc >= startIndex && abs(array[loc]) > abs(element)){
            array[loc + 1] = array[loc];
            loc -= 1;
        }
        array[loc + 1] = element;
        i += 1;
    }
}

void merge(int* array, int k, int l, int m){
    int t[m - k + 1];
    int i = k;
    int j = l + 1;
    int h = 0;
    while (h < m - k + 1){
        if (j <= m && (i == l + 1 || abs(array[j]) < abs(array[i]))){
            t[h] = array[j];
            j += 1;
        }else{
            t[h] = array[i];
            i += 1;
        }
        h += 1;
    }
    for (int index = 0; index <= h - 1; index++) {
        array[k + index] = t[index];
    }
}

void mergeSortRec(int* array, int low, int high){
    if (low < high){
        int med = (low + high) / 2;
        if (high - low + 1 < 5){
            insertionSort(array, low, high + 1);
        }else{
            mergeSortRec(array, low, med);
            mergeSortRec(array, med + 1, high);
            merge(array, low, med, high);
        }
    }
}

int main(){
    int length;
    scanf("%d", &length);
    int array[length];
    for (int i = 0; i < length; i++){
        scanf("%d", &array[i]);
    }
    mergeSortRec(array, 0, length - 1);
    for (int i = 0; i < length; i++){
        printf("%d ", array[i]);
    }
    return 0;
}