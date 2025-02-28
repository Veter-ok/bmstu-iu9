#include <stdio.h>
#include <stdlib.h>

struct Task {
    int index;
    int k;
    int v;
};

struct PriorityQueue {
    struct Task **heap;
    int cap;
    int count;
};

void swap(struct Task **a, struct Task **b) {
    struct Task *copyA = *a;
    *a = *b;
    *b = copyA;
}

struct PriorityQueue InitPriorityQueue(int n){
    struct Task **heap = (struct Task **)malloc(n * sizeof(struct Task *));
    struct PriorityQueue queue = {
        heap, n, 0
    };
    return queue;
}

void Insert(struct PriorityQueue *queue, struct Task *x){
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

struct Task* ExtractMax(struct PriorityQueue *queue){
    struct Task *ptr = queue->heap[0];
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
    int n, m;
    int ans = 0;
    scanf("%d\n%d", &n, &m);
    struct PriorityQueue queue = InitPriorityQueue(n);
    for (int i = 0; i < m; i++){
        int start, duration;
        scanf("%d %d", &start, &duration);
        struct Task *new_task = malloc(sizeof(struct Task));
        int time_finish = start + duration;
        if (queue.count < n){
            new_task->k = -time_finish;
            new_task->v = time_finish;
            Insert(&queue, new_task); 
        }else{
            struct Task *pop_task = ExtractMax(&queue);
            time_finish = start > pop_task->v ? time_finish : (pop_task->v + duration);
            new_task->k = -time_finish;
            new_task->v = time_finish;
            Insert(&queue, new_task);
            free(pop_task);
        }
        ans = time_finish > ans ? time_finish : ans;  
    }
    for (int i = 0; i < n; i++){
        struct Task *pop_task = ExtractMax(&queue);
        if (pop_task != NULL){
            ans = pop_task->v > ans ? pop_task->v : ans;
            free(pop_task);
        }
    }
    printf("%d", ans);
    free(queue.heap);
    return 0;
}