#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Elem {
    struct Elem *next;
    int v;
    unsigned int k;
};

unsigned int getHash(unsigned int i, size_t m){
    return i % m;
}

struct Elem** initHashTable(size_t m){
    struct Elem **hashTable = calloc(m, sizeof(struct Elem *));
    return hashTable;
}

struct Elem* ListSearch(struct Elem* l, unsigned int k){
    struct Elem *x = l;
    while (x != NULL && x->k != k){
        x = x->next;
    }
    return x;
}

void DeleteElement(struct Elem** t, unsigned int k, size_t m){
    unsigned int i = getHash(k, m);
    struct Elem *current = t[i];
    struct Elem *prev = NULL;
    while (current != NULL && current->k != k){
        prev = current;
        current = current->next;
    }
    if (current == NULL) return;

    if (prev == NULL){
        t[i] = current->next;
    } else {
        prev->next = current->next;
    }
    free(current);
}

void Assign(struct Elem** t, unsigned int k, int v, size_t m){
    if (v == 0) {
        DeleteElement(t, k, m);
        return;
    }
    unsigned int i = getHash(k, m);
    struct Elem *foundedElem = ListSearch(t[i], k);
    if (foundedElem == NULL){
        struct Elem *new_elem = malloc(sizeof(struct Elem));
        new_elem->k = k;
        new_elem->v = v;
        new_elem->next = t[i];
        t[i] = new_elem;
    }else{
        foundedElem->v = v;
    }
}

int At(struct Elem** t, unsigned int k, size_t m){
    struct Elem *p = ListSearch(t[getHash(k, m)], k);
    if (p == NULL) return 0;
    return p->v;
}

int main(){
    size_t m;
    scanf("%zu", &m);
    struct Elem **hashTable = initHashTable(m);
    char command[20] = {0};
    while (strcmp(command, "END") != 0){
        scanf("%s", command);
        if (strcmp(command, "ASSIGN") == 0){
            unsigned int i;
            int v;
            scanf("%u%d", &i, &v);
            Assign(hashTable, i, v, m);
        }
        if (strcmp(command, "AT") == 0){
            unsigned int i;
            scanf("%u", &i);
            printf("%d\n", At(hashTable, i, m));
        }
    }
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