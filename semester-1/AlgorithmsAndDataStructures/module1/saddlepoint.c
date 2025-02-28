#include <stdio.h>

int main(){
    unsigned int n, m;
    scanf("%u %u", &n, &m);
    long long int matrix[n][m];
    long long int maxInRows[n];
    long long int minInColumns[m];

    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++){
            scanf("%lld", &matrix[i][j]);
            minInColumns[j] = 100000000;
        }
        maxInRows[i] = -100000000;
    }

    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++){
            if (matrix[i][j] > maxInRows[i]){
                maxInRows[i] = matrix[i][j];
            }
            if (matrix[i][j] < minInColumns[j]){
                minInColumns[j] = matrix[i][j];
            }
        }
    }

    for (int i = 0; i < n; i++){
        for (int j = 0; j < m; j++){
            if (maxInRows[i] == minInColumns[j]){
                printf("%d %d", i, j);
                return 0;
            }
        }
    }
    printf("none");
}