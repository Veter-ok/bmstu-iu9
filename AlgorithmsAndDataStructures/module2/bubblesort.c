void bubblesort(unsigned long nel,
                int (*compare)(unsigned long i, unsigned long j),
                void (*swap)(unsigned long i, unsigned long j)) {
    unsigned long left = 0, right = nel - 1;
    int swapped = 0;
    while (left < right) {
        swapped = 0;
        for (unsigned long i = left; i < right; ++i) {
            if (compare(i, i + 1) > 0) {
                swap(i, i + 1);
                swapped = 1;
            }
        }
        right--;
        for (unsigned long i = right; i > left; --i) {
            if (compare(i - 1, i) > 0) {
                swap(i - 1, i);
                swapped = 1;
            }
        }
        left++;
        if (!swapped) {
            break;
        }
    }
}
