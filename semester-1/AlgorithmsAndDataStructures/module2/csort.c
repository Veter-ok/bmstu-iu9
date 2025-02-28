#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int split_string(char *src, char **words){
    char *start = src;
    int count = 0;
    while (*start != '\0'){
        while (*start == ' ') {
            start++;
        }
        if (*start == '\0') {
            break;
        }
        char *end = start;
        while (*end != '\0' && *end != ' '){
            end++;
        }
        int length = end - start;
        words[count] = (char *)malloc(length + 1);
        strncpy(words[count], start, length);
        words[count][length] = '\0';
        count++;
        start = end;
    }
    return count;
}

void csort(char *src, char *dest) {
    char *words[1001];
    int count_word = split_string(src, words);
    int count[count_word];
    for (int i = 0; i < count_word; i++){
        count[i] = 0;
    }
    for (int j = 0; j < count_word - 1; j++) {
        for (int i = j + 1; i < count_word; i++) {
            if (strlen(words[i]) < strlen(words[j])) {
                count[j]++;
            } else {
                count[i]++;
            }
        }
    }
    dest[0] = '\0';
    for (int i = 0; i < count_word; i++) {
        for (int j = 0; j < count_word; j++){
            if (count[j] == i){
                strcat(dest, words[j]);
                if (i < count_word - 1) {
                    strcat(dest, " ");
                }
            }
        }
    }
    for (int i = 0; i < count_word; i++) {
        free(words[i]);
    }
    
}

int main(){
    char *string = (char *)malloc(1001);
    char *result = (char *)malloc(1001);
    fgets(string, 1000, stdin);
    string[strcspn(string, "\n")] = '\0';
    csort(string, result);
    printf("%s", result);
    free(string);
    free(result);
    return 0;
}