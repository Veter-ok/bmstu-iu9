#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Stack {
    int *data;
    int *max1;
    int *max2;
    int top1;
    int top2;
    int cap;
};

struct Stack InitDoubleStack(int cap){
    int *data = (int *)malloc(cap * sizeof(int));
    int *max1 = (int *)malloc(cap * sizeof(int));
    int *max2 = (int *)malloc(cap * sizeof(int));
    struct Stack stack = {data, max1, max2, 0, cap - 1, cap};
    return stack;
}

int StackEmpty1(struct Stack *stack){
    return stack->top1 == 0;
}

int StackEmpty2(struct Stack *stack){
    return stack->top2 == stack->cap - 1;
}

void Push1(struct Stack *stack, int x){
    stack->data[stack->top1] = x;
    if (stack->top1 == 0) {
        stack->max1[stack->top1] = x;
    } else {
        stack->max1[stack->top1] = (x > stack->max1[stack->top1 - 1]) ? x : stack->max1[stack->top1 - 1];
    }
    stack->top1++;
}

void Push2(struct Stack *stack, int x){
    stack->data[stack->top2] = x;
    if (stack->top2 == stack->cap - 1) {
        stack->max2[stack->top2] = x;
    } else {
        stack->max2[stack->top2] = (x > stack->max2[stack->top2 + 1]) ? x : stack->max2[stack->top2 + 1];
    }
    stack->top2--;
}

int Pop1(struct Stack *stack){
    return stack->data[--stack->top1];
}

int Pop2(struct Stack *stack){
    return stack->data[++stack->top2];
}

struct Stack InitQueue(int cap){
    return InitDoubleStack(cap);
}

int QueueEmpty(struct Stack *stack){
    return StackEmpty1(stack) && StackEmpty2(stack);
}

void Enqueue(struct Stack *stack, int x){
    Push1(stack, x);
}

int Dequeue(struct Stack *stack){
    if (StackEmpty2(stack)){
        while (!StackEmpty1(stack)) {
            Push2(stack, Pop1(stack));
        }
    }
    return Pop2(stack);
}

int Maximum(struct Stack *stack){
    if (StackEmpty1(stack)) {
        return stack->max2[stack->top2 + 1];
    } 
    if (StackEmpty2(stack)) {
        return stack->max1[stack->top1 - 1];
    }
    int max1 = stack->max1[stack->top1 - 1];
    int max2 = stack->max2[stack->top2 + 1];
    return (max1 > max2) ? max1 : max2;
}

int main(){
    struct Stack stack = InitQueue(1000000);
    char command[20] = {0};
    while (strcmp(command, "END")){
        scanf("%s", command);
        if (!strcmp(command, "ENQ")){
            int x;
            scanf("%d", &x);
            Enqueue(&stack, x);
        } 
        if (!strcmp(command, "DEQ")) printf("%d\n", Dequeue(&stack));
        if (!strcmp(command, "EMPTY")) {
            if (QueueEmpty(&stack)){
                printf("true\n");
            }else{
                printf("false\n");
            }
        }
        if (!strcmp(command, "MAX")) printf("%d\n", Maximum(&stack));
    }
    free(stack.data);
    free(stack.max1);
    free(stack.max2);
    return 0;
}