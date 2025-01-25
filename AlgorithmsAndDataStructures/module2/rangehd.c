#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void Build(int* T, char* v, int i, int a, int b) {
    if (a == b) {
        T[i] = 1 << (v[a] - 'a');
    } else {
        int m = (a + b) / 2;
        Build(T, v, 2 * i + 1, a, m);
        Build(T, v, 2 * i + 2, m + 1, b);
        T[i] = T[2 * i + 1] ^ T[2 * i + 2];
    }
}

int* Init(char* v, size_t n) {
    int* tree = malloc(4 * n * sizeof(int));
    Build(tree, v, 0, 0, n - 1);
    return tree;
}

void Update(int* T, char *arr, int j, char* s, int i, int a, int b) {
    int len = strlen(s);
    if (a == b) {
        if (j <= a && j + len - 1 >= a) {
            T[i] = 1 << (s[a - j] - 'a');
            arr[a] = s[a - j];
        }
    } else {
        int m = (a + b) / 2;
        if (j <= m) {
            Update(T, arr, j, s, 2 * i + 1, a, m);
        }
        if (j + len - 1 > m) {
            Update(T, arr, j, s, 2 * i + 2, m + 1, b);
        }
        T[i] = T[2 * i + 1] ^ T[2 * i + 2];
    }
}

int query(int* T, int l, int r, int i, int a, int b) {
    if (l > b || r < a) return 0;
    if (l <= a && r >= b) return T[i];
    int m = (a + b) / 2;
    return query(T, l, r, 2 * i + 1, a, m) ^ query(T, l, r, 2 * i + 2, m + 1, b);
}

int countBits(int mask) {
    int count = 0;
    while (mask) {
        count += (mask & 1);
        mask >>= 1;
    }
    return count;
}

int main() {
    char *string = (char *)malloc(1000001);
    fgets(string, 1000001, stdin);
    string[strcspn(string, "\n")] = '\0';
    int n = strlen(string);
    int* tree = Init(string, n);
    char command[10] = {0};
    while (strcmp(command, "END") != 0) {
        scanf("%s", command);
        if (strcmp(command, "HD") == 0) {
            int l, r;
            scanf("%d %d", &l, &r);
            int mask = query(tree, l, r, 0, 0, n - 1);
            printf("%s\n", countBits(mask) <= 1 ? "YES" : "NO");
        }
        if (strcmp(command, "UPD") == 0) {
            int i;
            char s[1000001];
            scanf("%d %s", &i, s);
            Update(tree, string, i, s, 0, 0, n - 1);
        }
    }
    free(string);
    free(tree);
    return 0;
}