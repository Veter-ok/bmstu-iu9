#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Element {
    int index;
    int k;
    int v;
};

struct PriorityQueue {
    struct Element **heap;
    int cap;
    int count;
};

void swap(struct Element **a, struct Element **b) {
    struct Element *copyA = *a;
    *a = *b;
    *b = copyA;
}

struct PriorityQueue InitPriorityQueue(int n){
    struct Element **heap = (struct Element **)malloc(n * sizeof(struct Element *));
    struct PriorityQueue queue = {
        heap, n, 0
    };
    return queue;
}

void Insert(struct PriorityQueue *queue, struct Element *x){
    int i = queue->count++;
    x->index = i;
    queue->heap[i] = x;
    while (i > 0 && queue->heap[(i - 1) / 2]->k < x->k) {
        swap(&queue->heap[(i - 1) / 2], &queue->heap[i]);
        queue->heap[i]->index = i;
        i = (i - 1) / 2;
    }
    queue->heap[i]->index = i;
}

struct Element* ExtractMax(struct PriorityQueue *queue){
    struct Element *ptr = queue->heap[0];
    queue->count--;
    if (queue->count > 0){
        queue->heap[0] = queue->heap[queue->count];
        queue->heap[0]->index = 0;
        int i = 0;
        while (1){
            int l = 2*i + 1;
            int r = l + 1;
            int j = i;
            if ((l < queue->count) && queue->heap[i]->k < queue->heap[l]->k) {
                i = l;
            }
            if ((r < queue->count) && queue->heap[i]->k < queue->heap[r]->k) {
                i = r;
            }
            if (i == j){
                break;
            }
            swap(&queue->heap[i], &queue->heap[j]);
            queue->heap[i]->index = i;
            queue->heap[j]->index = j;
        }
    }
    return ptr;
} 

int main(){
    int k;
    scanf("%d", &k);
    int *lengths = (int *)malloc(k * sizeof(int));
    int **numbers = (int **)malloc(k * sizeof(int *));
    int size = 0;
    for (int i = 0; i < k; i++){
        scanf("%d", &lengths[i]);
        numbers[i] = (int *)malloc(lengths[i] * sizeof(int));
        size += lengths[i];
    }
    for (int i = 0; i < k; i++){
        for (int j = 0; j < lengths[i]; j++){
            scanf("%d", &numbers[i][j]);
        }
    }
    struct PriorityQueue queue = InitPriorityQueue(size);
    int indexes[k];
    for (int i = 0; i < k; i++) {
        struct Element *element = malloc(sizeof(struct Element));
        element->k = -numbers[i][0];
        element->v = i;
        element->index = 0;
        Insert(&queue, element);
        indexes[i] = 1;
    }
    for (int i = 0; i < size; i++){
        struct Element *element = ExtractMax(&queue);
        printf("%d ", -element->k);
        int arrayIndex = element->v;
        if (indexes[arrayIndex] < lengths[arrayIndex]) {
            element->k = -numbers[arrayIndex][indexes[arrayIndex]];
            element->index = indexes[arrayIndex]++;
            Insert(&queue, element);
        }else{
            free(element);
        }
    }
    printf("\n");
    for (int i = 0; i < k; i++) {
        free(numbers[i]);
    }
    free(numbers);
    free(lengths);
    free(queue.heap);
    return 0;
}