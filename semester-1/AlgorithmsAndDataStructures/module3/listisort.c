#include <stdio.h>
#include <stdlib.h>

struct Elem {
    struct Elem *prev, *next;
    int v;
};

struct Elem* insertionSort(struct Elem *start, int n){
    struct Elem *elem = start; 
    int elem_value = elem->v;
    struct Elem *loc_el = start; 
    for (int i = 1; i < n; i++){
        elem = elem->next; 
        elem_value = elem->v;
        loc_el = elem->prev; 
        int loc = i - 1;
        while (loc >= 0 && loc_el->v > elem_value){
            loc_el->next->v = loc_el->v;
            loc_el = loc_el->prev;
            loc--;
        }
        loc_el->next->v = elem_value;
    }
    return start;
}

int main(){
    int n;
    scanf("%d", &n);
    struct Elem *first = malloc(sizeof(struct Elem));
    first->prev = first;
    first->next = first;
    struct Elem *currently = first;
    for(int i = 0; i < n; i++){
        struct Elem *element = malloc(sizeof(struct Elem));
        int value;
        scanf("%d", &value);
        element->prev = currently;
        element->next = currently->next;
        currently->next->prev = currently;
        currently->next = element;
        currently->v = value;
        if(i == n - 1){
            free(element);
            break;
        }
        currently = element;
    }
    currently->next = first;
    first->prev = currently;
    currently = currently->next;
    struct Elem *start = insertionSort(currently, n);
    for (int i = 0; i < n; i++) {
        printf("%d ", start->v);
        struct Elem *save_elem = start; 
        start = start->next; 
        free(save_elem);
    }
    return 0;
}