#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void DistributionSort(char *string, int length, char *result){
    int count[26] = {0};
    int k = 0;
    for (int j = 0; j < length; j++){
        k = string[j] - 'a';
        count[k]++;
    }
    for (int i = 1; i < 26; i++){
        count[i] += count[i - 1];
    }
    int i = 0;
    for (int j = length - 1; j >= 0; j--){
        k = string[j] - 'a';
        i = count[k] - 1;
        count[k] = i;
        result[i] = string[j];
    }
    result[length] = '\0';
}

int main(){
    char *string = (char *)malloc(1000001);
    char *result = (char *)malloc(1000001);
    fgets(string, 1000001, stdin);
    string[strcspn(string, "\n")] = '\0';
    DistributionSort(string, strlen(string), result);
    printf("%s", result);
    free(string);
    free(result);
    return 0;
}