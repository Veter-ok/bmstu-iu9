#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Elem {
    int k;
    char *v;
    int count;
    struct Elem* parent; 
    struct Elem* left;
    struct Elem* right; 
};

struct Elem* Minimum(struct Elem *t){
    struct Elem *x = t;
    while(x->left != NULL) x = x->left;
    return x;
}

struct Elem* Succ(struct Elem *x){
    if (x->right != NULL){
        return Minimum(x->right);
    }
    struct Elem* y = x->parent;
    while (y != NULL && x == y->right){
        x = y;
        y = y->parent;
    }
    return y;
}

char* SearchByRank(struct Elem *t, int count){
    struct Elem *x = t;
    while (x->count != count){
        if (count < x->count){
            x = x->left;
        }else{
            count -= x->count + 1;
            x = x->right;
        }
    }
    return x->v;
}

void Insert(struct Elem **t, int k, char *v){
    struct Elem *y = malloc(sizeof(*y));
    y->k = k;
    y->v = v;
    y->parent = NULL;
    y->left = NULL;
    y->right = NULL;
    y->count = 0;
    if (*t == NULL){
        *t = y;
    }else{
        struct Elem *x = *t;
        while (1){
            if (k < x->k){
                x->count++;
                if (x->left == NULL){
                    x->left = y;
                    y->parent = x;
                    break;
                }
                x = x->left;
            }else{
                if (x->right == NULL){
                    x->right = y;
                    y->parent = x;
                    break;
                }
                x = x->right;
            }
        }
    }
}

struct Elem* Descend(struct Elem *t, int k){
    struct Elem *x = t;
    while (x != NULL && x->k != k){
        if (k < x->k){
            x = x->left;
        }else{
            x = x->right;
        }
    }
    return x;
}

char* Lookup(struct Elem *t, int k){
    return Descend(t, k)->v;
}

void updateCount(struct Elem *t) {
    while (t != NULL && t->parent != NULL) {
        if (t->parent->left == t) {
            t->parent->count--;
        }
        t = t->parent;
    }
}

void ReplaceNode(struct Elem **t, struct Elem *x, struct Elem *y){
    if (x == *t){
        *t = y;
        if (y != NULL) y->parent = NULL;
    }else{
        struct Elem *p = x->parent;
        if (y != NULL) y->parent = p;
        if (p->left == x){
            p->left = y;
        }else{
            p->right = y;
        }
    }
}

void Delete(struct Elem **t, int k){
    struct Elem *x = Descend(*t, k);
    if (x->left == NULL && x->right == NULL){
        updateCount(x);
        ReplaceNode(t, x, NULL);
    }else if (x->left == NULL){
        updateCount(x);
        ReplaceNode(t, x, x->right);
    }else if (x->right == NULL){
        updateCount(x);
        ReplaceNode(t, x, x->left);
    }else{
        struct Elem *y = Succ(x);
        updateCount(y);
        ReplaceNode(t, y, y->right);
        y->count = x->count;
        x->left->parent = y;
        y->left = x->left;
        if (x->right != NULL) {
            x->right->parent = y;
            y->right = x->right;
        }
        ReplaceNode(t, x, y);
    }
    free(x->v);
    free(x);
}

void FreeTree(struct Elem *t) {
    if (t == NULL) {
        return;
    }
    FreeTree(t->left);
    FreeTree(t->right);
    free(t->v);
    free(t);
}

int main(){
    struct Elem *binaryTree = NULL;
    char command[20] = {0};
    while (strcmp(command, "END")){
        scanf("%s", command);
        if (!strcmp(command, "INSERT")){
            int k;
            char *v = malloc(11 * sizeof(char));
            scanf("%d%s", &k, v);
            Insert(&binaryTree, k, v);
        }
        if (!strcmp(command, "LOOKUP")){
            int k;
            scanf("%d", &k);
            char *v = Lookup(binaryTree, k);
            printf("%s\n", v);
        }
        if (!strcmp(command, "DELETE")){
            int k;
            scanf("%d", &k);
            Delete(&binaryTree, k);
        }
        if (!strcmp(command, "SEARCH")){
            int k;
            scanf("%d", &k);
            char *v = SearchByRank(binaryTree, k);
            printf("%s\n", v);
        }
    }
    FreeTree(binaryTree);
    return 0;
}