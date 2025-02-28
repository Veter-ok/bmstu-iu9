#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Queue{
    int *data;
    int cap, count, head, tail;
};

struct Queue InitQueue(int n){
    int *data = (int *)malloc(n * sizeof(int));
    struct Queue queue = {data, n, 0, 0, 0};
    return queue;
}

void ExpandingQueue(struct Queue *queue){
    int new_cap = queue->cap * 2;
    int *data = (int *)malloc(new_cap * sizeof(int));
    for (int i = 0; i < queue->count; i++){
        data[i] = queue->data[(queue->head + i) % queue->cap];
    }
    free(queue->data);
    queue->data = data;
    queue->tail = queue->count;
    queue->head = 0;
    queue->cap = new_cap;
}

int QueueEmpty(struct Queue *queue){
    printf(queue->count ? "false\n" : "true\n");
    return queue->count == 0;
}

void Enqueue(struct Queue *queue){
    int x;
    scanf("%d", &x);
    if (queue->count == queue->cap){
        ExpandingQueue(queue);
    }
    queue->data[queue->tail++] = x;
    if (queue->tail == queue->cap) queue->tail = 0;
    queue->count++;
}

void Dequeue(struct Queue *queue){
    int x = queue->data[queue->head++];
    if (queue->head == queue->cap) queue->head = 0;
    queue->count--;
    printf("%d\n", x);
}


int main(){
    struct Queue queue = InitQueue(4);
    char command[20] = {0};
    while (strcmp(command, "END") != 0){
        scanf("%s", command);
        if (strcmp(command, "ENQ") == 0) Enqueue(&queue);
        if (strcmp(command, "DEQ") == 0) Dequeue(&queue);
        if (strcmp(command, "EMPTY") == 0) QueueEmpty(&queue);
    }
    free(queue.data);
    return 0;
}