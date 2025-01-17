#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int overlap(const char *a, const char *b) {
    int len_a = strlen(a), len_b = strlen(b);
    for (int i = 0; i < len_a; i++) {
        if (strncmp(a + i, b, len_a - i) == 0) {
            return len_a - i;
        }
    }
    return 0;
}

int shortest_superstring(int n, char **strings) {
    int overlap_matrix[n][n];
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (i != j) {
                overlap_matrix[i][j] = overlap(strings[i], strings[j]);
            } else {
                overlap_matrix[i][j] = 0;
            }
        }
    }
    int dp[1 << n][n];
    int parent[1 << n][n];
    memset(dp, 0x3f, sizeof(dp));
    memset(parent, -1, sizeof(parent));

    for (int i = 0; i < n; i++) {
        dp[1 << i][i] = strlen(strings[i]);
    }

    for (int mask = 1; mask < (1 << n); mask++) {
        for (int i = 0; i < n; i++) {
            if (!(mask & (1 << i))) continue;
            for (int j = 0; j < n; j++) {
                if (mask & (1 << j)) continue;
                int next_mask = mask | (1 << j);
                int new_length = dp[mask][i] + strlen(strings[j]) - overlap_matrix[i][j];
                if (new_length < dp[next_mask][j]) {
                    dp[next_mask][j] = new_length;
                    parent[next_mask][j] = i;
                }
            }
        }
    }
    int min_length = 0x3f3f3f3f, last = -1, final_mask = (1 << n) - 1;
    for (int i = 0; i < n; i++) {
        if (dp[final_mask][i] < min_length) {
            min_length = dp[final_mask][i];
            last = i;
        }
    }
    char result[100 * n];
    char merged[100 * n];
    int mask = final_mask;
    int order[n], idx = n - 1;
    while (last != -1) {
        order[idx--] = last;
        int temp = parent[mask][last];
        mask ^= (1 << last);
        last = temp;
    }
    strcpy(result, strings[order[0]]);
    for (int i = 1; i < n; i++) {
        strcpy(merged, result);
        strcat(merged, strings[order[i]] + overlap_matrix[order[i - 1]][order[i]]);
        strcpy(result, merged);
    }
    printf("%d\n", min_length);
    return min_length;
}

int main() {
    int n;
    scanf("%d", &n);
    char **strings = malloc(n * sizeof(char *));
    for (int i = 0; i < n; i++) {
        strings[i] = (char *)malloc(1001 * sizeof(char)); 
        scanf("%s", strings[i]);
    }
    shortest_superstring(n, strings);
    for (int i = 0; i < n; i++) {
        free(strings[i]);
    }
    free(strings);
    return 0;
}