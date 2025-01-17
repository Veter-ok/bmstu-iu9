#include <stdio.h>
#include <stdlib.h>

struct Elem {
    struct Elem *prev, *next;
    int number;
};

struct Elem* InitDoubleLinkedList(){
    struct Elem *linkedList = malloc(sizeof(struct Elem));
    linkedList->next = linkedList;
    linkedList->prev = linkedList;
    return linkedList;
}

void InsertAfter(struct Elem *x, struct Elem *y){
    struct Elem* z = x->next;
    x->next = y;
    y->prev = x;
    y->next = z;
    z->prev = y;
}

void Delete(struct Elem *x){
    struct Elem* y = x->prev;
    struct Elem* z = x->next;
    y->next = z;
    z->prev = y;
    x->prev = NULL;
    x->next = NULL;
}

void selectionSort(struct Elem *head, struct Elem *new_array, int n){
    while (head->next != head && head->prev != head){
        struct Elem *min_elem = head->next;
        head = head->next;
        for (int i = 0; i < n; i++){
            if (abs(head->number) < abs(min_elem->number)){
                min_elem = head;
            }
            head = head->next;
        }
        Delete(min_elem);
        InsertAfter(new_array, min_elem);
        new_array = new_array->next;
        n--;
    }
}

int main(){
    int n;
    scanf("%d", &n);
    struct Elem *new_head = InitDoubleLinkedList();
    struct Elem *head = InitDoubleLinkedList();
    struct Elem *currently = head;
    for(int i = 0; i < n; i++){
        struct Elem *element = malloc(sizeof(struct Elem));
        int value;
        scanf("%d", &value);
        element->number = value;
        InsertAfter(currently, element);
        currently = currently->next;
    }
    currently->next = head;
    head->prev = currently;
    currently = currently->next;

    selectionSort(head, new_head, n);
    new_head = new_head->next;
    for (int i = 0; i < n; i++){
        struct Elem *nextElem = new_head->next;
        InsertAfter(head, new_head);
        head = head->next;
        new_head = nextElem;
    }
    head = head->next->next;
    for (int i = 0; i < n; i++){
        printf("%d ", head->number);
        struct Elem *save_elem = head;
        head = head->next;
        free(save_elem);
    }
    free(new_head);
    free(head);
    return 0;
}