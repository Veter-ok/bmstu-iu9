#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *fibstr(int n)
{
    if (n <= 2){
        char *result = (char *)malloc(2);
        result[0] = n % 2 == 0 ? 'b' : 'a';
        result[1] = '\0';
        return result;
    }
    char *str1 = malloc(2 * sizeof(char));
    char *str2 = malloc(2 * sizeof(char));
    char *result = NULL;
    strcpy(str1, "a");
    strcpy(str2, "b");
    for (int i = 2; i < n; i++){
        result = malloc((strlen(str1) + strlen(str2) + 1) * sizeof(char));
        strcpy(result, str1);
        strcat(result, str2);
        free(str1);
        str1 = str2;
        str2 = result;
    }
    free(str1);
    return result;
}

int main(){
    int n;
    scanf("%d", &n);
    char *result = fibstr(n);
    printf("%s", result);
    free(result);
    return 0;
}