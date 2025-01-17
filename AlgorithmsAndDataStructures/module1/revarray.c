#include <stdio.h>

void revarray(void *base, size_t nel, size_t width) {
    char *start = base;
    char *end = start + ((nel - 1) * width);

    if (start < end){
        for (size_t i = 0; i < width; i++) {
            char preserve_start = start[i];
            start[i] = end[i];
            end[i] = preserve_start;
        }
        revarray(start + width, nel - 2, width);
    }
}