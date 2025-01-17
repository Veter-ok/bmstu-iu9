#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Здравствуйте, Александр Владимирович!
// Поздравляю Вас с Новым годом! 
// Желаю счастья, здоровья и успехов в Новом году!

void Build(int* T, int* v, int i, int a, int b) {
    if (a == b) {
        T[i] = v[a];
    } else {
        int m = (a + b) / 2;
        Build(T, v, 2 * i + 1, a, m);
        Build(T, v, 2 * i + 2, m + 1, b); 
        T[i] = T[2 * i + 1] > T[2 * i + 2] ? T[2 * i + 1] : T[2 * i + 2];
    }
}

int* Init(int* v, size_t n) {
    int* tree = malloc(4 * n * sizeof(int)); 
    Build(tree, v, 0, 0, n - 1);  
    return tree;
}

void Update(int* T, int j, int x, int i, int a, int b) {
    if (a == b) {
        T[i] = x;
    } else {
        int m = (a + b) / 2;
        if (j <= m) {
            Update(T, j, x, 2 * i + 1, a, m);
        } else {
            Update(T, j, x, 2 * i + 2, m + 1, b);
        }
        T[i] = T[2 * i + 1] > T[2 * i + 2] ? T[2 * i + 1] : T[2 * i + 2];
    }
}

int query(int* T, int l, int r, int i, int a, int b) {
    if (l > b || r < a) return -1000000000;
    if (l <= a && r >= b) return T[i];
    int m = (a + b) / 2;
    int v1 = query(T, l, r, 2 * i + 1, a, m); 
    int v2 = query(T, l, r, 2 * i + 2, m + 1, b);
    return v1 > v2 ? v1 : v2;
}

int main() {
    size_t n;
    scanf("%zu", &n);
    int* arr = malloc(n * sizeof(int));
    for (size_t i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    int* tree = Init(arr, n);
    char command[4] = {0};
    while ((strcmp(command, "END"))) {
        scanf("%s", command);
        if (!strcmp(command, "MAX")) {
            int l, r;
            scanf("%d%d", &l, &r);
            printf("%d\n", query(tree, l, r, 0, 0, n - 1));
        }
        if (!strcmp(command, "UPD")) {
            int i, v;
            scanf("%d%d", &i, &v);
            Update(tree, i, v, 0, 0, n - 1);
        }
    }
    free(arr);
    free(tree);
    return 0;
}