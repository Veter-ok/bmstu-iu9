#include <stdio.h>

void shellsort(unsigned long nel,
    int (*compare)(unsigned long i, unsigned long j),
    void (*swap)(unsigned long i, unsigned long j))
{
    int fib_1 = 1;
    int fib_2 = 1;
    while (fib_1 + fib_2 < nel) {
        int copyFib_1 = fib_1;
        fib_1 = fib_1 + fib_2;
        fib_2 = copyFib_1;
    }
    while(fib_1 > 0) {
        unsigned long d = fib_1;
        for (unsigned long i = d; i < nel; i++) {
            unsigned long loc = i;
            while (loc >= d && compare(loc, loc - d) < 0) {
                swap(loc, loc - d);
                loc -= d;
            }
        }
        int copyFib_2 = fib_2;
        fib_2 = fib_1 - fib_2;
        fib_1 = copyFib_2;
    }
}