#include <stdio.h>

unsigned long peak(unsigned long nel, int (*less)(unsigned long i, unsigned long j)) {
    if (nel == 0 || nel == 1 || less(1, 0)){
        return 0;
    }
    for(int i = 1; i < nel - 1; i++){
        if (!less(i - 1, i) && less(i + 1, i)){
            return i;
        }
    }
    return nel - 1;
}