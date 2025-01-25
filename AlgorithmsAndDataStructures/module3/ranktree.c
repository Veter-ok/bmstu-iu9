#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Elem {
    int k;
    char *v;
    int *span;
    struct Elem **next;
};

struct Elem* InitSkipList(int m){
    struct Elem *element = malloc(sizeof(struct Elem));
    element->next = calloc(m, sizeof(struct Elem*));
    element->span = calloc(m, sizeof(int));
    element->v = NULL;
    element->k = -1000000000;
    return element;
}

struct Elem* Succ(struct Elem* l){
    return l->next[0];
}

int Rank(struct Elem* l, int m, int k){
    if (k == -1000000000) return -1;
    int rank = 0;
    for (int i = m - 1; i >= 0; i--){
        while(l->next[i] != NULL && l->next[i]->k < k){
            rank += l->span[i];
            l = l->next[i];
        }
    }
    return rank;
}

void Skip(struct Elem* l, int m, int k, struct Elem **p){
    for (int i = m - 1; i >= 0; i--){
        while(l->next[i] != NULL && l->next[i]->k < k){
            l = l->next[i];
        }
        p[i] = l;
    }
}

char* Lookup(struct Elem* l, int m, int k){
    struct Elem **p = malloc(m * sizeof(struct Elem*));
    Skip(l, m, k, p);
    struct Elem *x = Succ(p[0]);
    free(p);
    return x->v;
}

void Insert(struct Elem* l, int m, int k, char* v){
    struct Elem **p = malloc(m * sizeof(struct Elem*));
    Skip(l, m, k, p);
    struct Elem *x = InitSkipList(m);
    x->k = k;
    x->v = v;
    int rankX = Rank(l, m, p[0]->k) + 1;
    int r = rand() * 2;
    int i = 0;
    for (i = 0; i < m && r % 2 == 0; i++){
        x->next[i] = p[i]->next[i];
        p[i]->next[i] = x;
        int rank = Rank(l, m, p[i]->k);
        x->span[i] = p[i]->span[i] - (rankX - rank) + 1;
        p[i]->span[i] = rankX - rank;
        r /= 2;
    }
    for(; i < m; i++){
        x->next[i] = NULL;
        p[i]->span[i]++;
    }
    free(p);
}

void Delete(struct Elem* l, int m, int k){
    struct Elem **p = malloc(m * sizeof(struct Elem*));
    Skip(l, m, k, p);
    struct Elem *x = Succ(p[0]);
    int rankX = Rank(l, m, k);
    int i = 0;
    for (i = 0; i < m && p[i]->next[i] == x; i++){
        p[i]->next[i] = x->next[i];
        p[i]->span[i] += x->span[i] - 1;
    }
    for (; i < m; i++){
        p[i]->span[i]--;
    }
    free(x->next);
    free(x->span);
    free(x->v);
    free(x);
    free(p);
}

int main(){
    int n;
    int m = 20;
    scanf("%d", &n);
    struct Elem *skipList = InitSkipList(m);
    char command[20] = {0};
    while (strcmp(command, "END")){
        scanf("%s", command);
        if (!strcmp(command, "INSERT")){
            int k;
            char *v = malloc(11 * sizeof(char));
            scanf("%d%s", &k, v);
            Insert(skipList, m, k, v);
        }
        if (!strcmp(command, "LOOKUP")){
            int k;
            scanf("%d", &k);
            char *v = Lookup(skipList, m, k);
            printf("%s\n", v);
        }
        if (!strcmp(command, "DELETE")){
            int k;
            scanf("%d", &k);
            Delete(skipList, m, k);
        }
        if (!strcmp(command, "RANK")){
            int k;
            scanf("%d", &k);
            int rank = Rank(skipList, m, k);
            printf("%d\n", rank);
        }
    }
    while (skipList != NULL) {
        struct Elem *copySkipList = skipList->next[0];
        if (skipList->v != NULL) free(skipList->v);
        free(skipList->next);
        free(skipList->span);
        free(skipList);
        skipList = copySkipList;
    }
    return 0;
}