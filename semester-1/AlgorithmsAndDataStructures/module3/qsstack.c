#include <stdio.h>
#include <stdlib.h>

struct Task {
    int low, high;
};

void swap(int *a, int *b){
    int copyA = *a;
    *a = *b;
    *b = copyA;
}

int partition(int *array, struct Task task){
    int i = task.low;
    for (int j = task.low; j < task.high; j++) {
        if (array[j] < array[task.high]){
            swap(&array[i], &array[j]);
            i++;
        }
    }
    swap(&array[i], &array[task.high]);
    return i;
}

void quicksort(int *arr, struct Task task){
    struct Task *stack = (struct Task *)malloc((task.high + 1) * sizeof(struct Task));
    stack[0] = task;
    int top = 0;
    while (top >= 0) {
        struct Task current = stack[top--];
        if (current.low < current.high) {
            int q = partition(arr, current);
            stack[++top] = (struct Task){current.low, q - 1};
            stack[++top] = (struct Task){q + 1, current.high};
        }
    }
    free(stack);
}

int main(){
    int n;
    scanf("%d", &n);
    int *arr = (int *)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    quicksort(arr, (struct Task){0, n - 1});
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    free(arr);
    return 0;
}