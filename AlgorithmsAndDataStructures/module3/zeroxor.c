#include <stdio.h>
#include <stdlib.h>

struct Elem {
    struct Elem *next;
    unsigned int count;
    long long xorValue;
};

unsigned int getHash(long long i, size_t m){
    return ((unsigned long long)i) % m;
}

struct Elem** initHashTable(size_t m){
    struct Elem **hashTable = calloc(m, sizeof(struct Elem *));
    return hashTable;
}

struct Elem* ListSearch(struct Elem* l, long long xorValue){
    struct Elem *x = l;
    while (x != NULL && x->xorValue != xorValue){
        x = x->next;
    }
    return x;
}

void Insert(struct Elem** t, long long xorValue, size_t m){
    unsigned int i = getHash(xorValue, m);
    struct Elem *foundedElem = ListSearch(t[i], xorValue);
    if (foundedElem == NULL){
        struct Elem *new_elem = malloc(sizeof(struct Elem));
        new_elem->xorValue = xorValue;
        new_elem->count = 1;
        new_elem->next = t[i];
        t[i] = new_elem;
    }else{
        foundedElem->count++;
    }
}

unsigned int getCount(struct Elem** t, long long xorValue, size_t m){
    struct Elem *p = ListSearch(t[getHash(xorValue, m)], xorValue);
    return p ? p->count : 0;
}

int main() {
    int n;
    scanf("%d", &n);
    size_t m = 100000;
    unsigned long long ans = 0;
    long long curXORvalue = 0;
    struct Elem **hashTable = initHashTable(m);
    for (int i = 0; i < n; i++) {
        long long x;
        scanf("%lld", &x);
        curXORvalue ^= x;
        if (curXORvalue == 0) ans += 1;
        ans += getCount(hashTable, curXORvalue, m);
        Insert(hashTable, curXORvalue, m);
    }
    printf("%llu", ans);
    for (size_t i = 0; i < m; i++){
        struct Elem *x = hashTable[i];
        while (x != NULL){
            struct Elem *next = x->next;
            free(x);
            x = next;
        }
    }
    free(hashTable);
    return 0;
}