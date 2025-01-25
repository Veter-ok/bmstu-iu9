#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int peak(size_t nel, long *arr, int i) {
    if (i == 0) {
        return arr[i] >= arr[i + 1] ? 1 : 0;
    }
    if (i == nel - 1){
        return arr[i] >= arr[i - 1] ? 1 : 0;
    }
    return arr[i] >= arr[i + 1] && arr[i] >= arr[i - 1] ? 1 : 0;
}

void Build(int* T, long* v, int i, int a, int b, size_t n) {
    if (a == b) {
        T[i] = peak(n, v, a);
    } else {
        int m = (a + b) / 2;
        Build(T, v, 2 * i + 1, a, m, n);
        Build(T, v, 2 * i + 2, m + 1, b, n);
        T[i] = T[2 * i + 1] + T[2 * i + 2];
    }
}

int* Init(long* v, size_t n) {
    int* tree = malloc(4 * n * sizeof(int));
    Build(tree, v, 0, 0, n - 1, n);
    return tree;
}

void Update(int* T, long *arr, int n, int j, int i, int a, int b) {
    if (a == b) {
        T[i] = peak(n, arr, a);
    } else {
        int m = (a + b) / 2;
        if (j <= m) {
            Update(T, arr, n, j, 2 * i + 1, a, m);
        } else {
            Update(T, arr, n, j, 2 * i + 2, m + 1, b);
        }
        T[i] = T[2 * i + 1] + T[2 * i + 2];
    }
}

int query(int* T, int l, int r, int i, int a, int b) {
    if (l > b || r < a) return 0;
    if (l <= a && r >= b) return T[i];
    int m = (a + b) / 2;
    int v1 = query(T, l, r, 2 * i + 1, a, m);
    int v2 = query(T, l, r, 2 * i + 2, m + 1, b);
    return v1 + v2;
}

int main() {
    size_t n;
    scanf("%zu", &n);
    long* arr = malloc(n * sizeof(long));
    for (size_t i = 0; i < n; i++) {
        scanf("%ld", &arr[i]);
    }
    int* tree = Init(arr, n);
    char command[10] = {0};
    while (strcmp(command, "END") != 0) {
        scanf("%s", command);
        if (strcmp(command, "PEAK") == 0) {
            int l, r;
            scanf("%d%d", &l, &r);
            printf("%d\n", query(tree, l, r, 0, 0, n - 1));
        }
        if (strcmp(command, "UPD") == 0) {
            int i, v;
            scanf("%d%d", &i, &v);
            arr[i] = v;
            Update(tree, arr, n, i, 0, 0, n - 1);
            if (i > 0)     Update(tree, arr, n, i - 1, 0, 0, n - 1);
            if (i < n - 1) Update(tree, arr, n, i + 1, 0, 0, n - 1);
        }
    }
    free(arr);
    free(tree);
    return 0;
}