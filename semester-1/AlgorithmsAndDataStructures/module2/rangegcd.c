#include <stdio.h>
#include <stdlib.h>

long gcd(long a, long b) {
    while (b) {
        long t = b;
        b = a % b;
        a = t;
    }
    return a < 0 ? -a : a;
}

void Build(long* T, long* v, int i, int a, int b) {
    if (a == b) {
        T[i] = v[a];
    } else {
        int m = (a + b) / 2;
        Build(T, v, 2 * i + 1, a, m);
        Build(T, v, 2 * i + 2, m + 1, b);
        T[i] = gcd(T[2 * i + 1], T[2 * i + 2]);
    }
}

long query(long* T, int l, int r, int i, int a, int b) {
    if (l > b || r < a) return 0;
    if (l <= a && r >= b) return T[i];
    int m = (a + b) / 2;
    int v1 = query(T, l, r, 2 * i + 1, a, m);
    int v2 = query(T, l, r, 2 * i + 2, m + 1, b);
    return gcd(v1, v2);
}

long* Init(long* v, size_t n) {
    long* tree = malloc(4 * n * sizeof(long));
    Build(tree, v, 0, 0, n - 1);
    return tree;
}

int main() {
    size_t n, m;
    scanf("%zu", &n);
    long* arr = malloc(n * sizeof(long));
    for (size_t i = 0; i < n; i++) {
        scanf("%ld", &arr[i]);
    }
    long* tree = Init(arr, n);
    scanf("%zu", &m);
    for (int i = 0; i < m; i++){
        int l, r;
        scanf("%d%d", &l, &r);
        printf("%ld\n", query(tree, l, r, 0, 0, n - 1));
    }
    free(arr);
    free(tree);
    return 0;
}