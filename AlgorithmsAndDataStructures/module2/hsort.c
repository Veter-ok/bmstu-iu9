#include <stdio.h>
#include <stdlib.h>

int count_a(char *str) {
    int count = 0;
    while (*str) {
        if (*str == 'a') {
            count++;
        }
        str++;
    }
    return count;
}

void swap(void *a, void *b, size_t width) {
    for (size_t i = 0; i < width; i++){
        char c = ((char *)a)[i];
        ((char *)a)[i] = ((char *)b)[i];
        ((char *)b)[i] = c;
    }
}

int compare(const void *a, const void *b) {
    int count_a1 = count_a(*(char **)a);
    int count_a2 = count_a(*(char **)b);
    return count_a1 - count_a2;
}

void Heapify(void *base, size_t n, size_t i, size_t width, int (*compare)(const void *a, const void *b)){
    while (1){
        size_t l = 2*i + 1;
        size_t r = l + 1;
        size_t j = i;
        if ((l < n) && compare((char *)base + l * width, (char *)base + i * width) > 0) {
            i = l;
        }
        if ((r < n) && compare((char *)base + r * width, (char *)base + i * width) > 0) {
            i = r;
        }
        if (i == j){
            break;
        }
        swap((char *)base + i * width, (char *)base + j * width, width);
    }
}

void BuildHeap(void *base, size_t nel, size_t width, int (*compare)(const void *a, const void *b)){
    for (int i = ((int)(nel / 2)) - 1; i >= 0; i--){
        Heapify(base, nel, i, width, compare);
    }
}

void hsort(void *base, size_t nel, size_t width, int (*compare)(const void *a, const void *b)) {
    BuildHeap(base, nel, width, compare);
    for(size_t i = nel - 1; i > 0; i--){
        swap((char *)base, (char *)base + i * width, width);
        Heapify(base, i, 0, width, compare);
    }
}

int main(){
    int length;
    scanf("%d", &length);
    char **strings = (char **)malloc(length * sizeof(char *));
    for (int i = 0; i < length; i++) {
        strings[i] = (char *)malloc(1001 * sizeof(char)); 
        scanf("%s", strings[i]);
    }
    hsort(strings, length, sizeof(char *), compare);
    for (int i = 0; i < length; i++) {
        printf("%s\n", strings[i]);
         free(strings[i]);
    }

    free(strings);
    return 0;
}