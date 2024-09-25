#include <stdio.h>

void revarray(void *base, size_t nel, size_t width) {
    char* start = (char *)base;
    // char* end = start + (nel - 1) * width;
    // for (int i = 0; i < 3; i++){
    //     // base = base + i;
    //     printf("%zu", base + i);
    // }
}

int main(){
    int array[3] = {1, 2, 3};
    int *ptr = array;
    revarray(ptr, 3, 6);
}