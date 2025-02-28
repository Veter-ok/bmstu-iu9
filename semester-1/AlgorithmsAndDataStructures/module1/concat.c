#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *concat(char **s, int n)
{
    int res_length = 0;
    for (int i = 0; i < n; i++){
        res_length += strlen(s[i]);
    }
    char *result = (char *)malloc(res_length + 1);
    int index = 0;
    for (int i = 0; i < n; i++) {
        int j = 0;
        for (int j = 0; s[i][j] != '\0'; j++){
            result[index] = s[i][j];
            index += 1;
        }
    }
    result[res_length] = '\0';
    return result;
}

int main(){
    int length;
    scanf("%d", &length);
    char **strings = (char **)malloc(length * sizeof(char *));
    for (int i = 0; i < length; i++) {
        strings[i] = (char *)malloc(1001 * sizeof(char)); 
        scanf("%s", strings[i]);
    }
    char *result = concat(strings, length);
    printf("%s", result);
    for (int i = 0; i < length; i++) {
        free(strings[i]); 
    }
    free(strings);
    free(result); 
    return 0;
}