#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Elem {
    struct Elem *next;
    char *word;
};

// украл у себя из csort.c
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

void swap(struct Elem *a, struct Elem *b) {
    char *temp = a->word;
    a->word = b->word;
    b->word = temp;
}

struct Elem *bsort(struct Elem *list) {
    struct Elem *left = list;
    int swapped = 0;
    while (left != NULL) {
        swapped = 0;
        for (struct Elem *i = left; i->next != NULL; i = i->next) {
            if (strlen(i->word) > strlen(i->next->word)) {
                swap(i, i->next);
                swapped = 1;
            }
        }
        if (!swapped) {
            break;
        }
    }
    return list;
}

int main(){ 
    char *string = (char *)malloc(1001);
    fgets(string, 1000, stdin);
    string[strcspn(string, "\n")] = '\0';
    char *words[1001];
    int n = split_string(string, words);
    struct Elem *first = malloc(sizeof(struct Elem));
    first->next = first;
    struct Elem *currently = first;
    for(int i = 0; i < n; i++){
        struct Elem *element = malloc(sizeof(struct Elem));
        element->next = currently->next;
        currently->next = element;
        currently->word = words[i];
        if(i == n - 1){
            free(element);
            break;
        }
        currently = element;
    }
    currently->next = NULL;
    currently = first;
    struct Elem *start = bsort(currently);
    for (int i = 0; i < n; i++) {
        printf("%s ", start->word);
        struct Elem *save_elem = start; 
        start = start->next; 
        free(save_elem);
    }
    for (int i = 0; i < n; i++){
        free(words[i]);
    }
    free(string);
    return 0;
}