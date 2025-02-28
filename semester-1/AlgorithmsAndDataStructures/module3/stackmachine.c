#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Stack {
    int *data;
    int top;
    int cap;
};

struct Stack InitStack(int cap){
    int *data = (int *)malloc(cap * sizeof(int));
    struct Stack stack = {data, 0, cap};
    return stack;
}

int StackEmpty(struct Stack *stack){
    return stack->top == 0;
}

void ExpandingStack(struct Stack *stack){
    int new_cap = stack->cap + 10;
    int *data = (int *)malloc(new_cap * sizeof(int));
    for (int i = 0; i < new_cap; i++){
        data[i] = i < stack->cap ? stack->data[i] : 0;
    }
    free(stack->data);
    stack->data = data;
    stack->cap = new_cap;
}

void Push(struct Stack *stack, int x){
    if (stack->top == stack->cap){
        ExpandingStack(stack);
    }
    stack->data[stack->top++] = x;
}

int Pop(struct Stack *stack){
    if (!StackEmpty(stack)){
        int x = stack->data[--stack->top];
        stack->data[stack->top] = 0;
        return x;
    }
    return -1;
}

void Const(struct Stack *stack){
    long long x;
    scanf("%lld", &x);
    Push(stack, x);
}

void Add(struct Stack *stack){
    Push(stack, Pop(stack) + Pop(stack));
}

void Sub(struct Stack *stack){
    Push(stack, Pop(stack) - Pop(stack));
}

void Mul(struct Stack *stack){
    Push(stack, Pop(stack) * Pop(stack));
}

void Div(struct Stack *stack){
    int x = Pop(stack) / Pop(stack);
    Push(stack, x);
}

void Max(struct Stack *stack){
    int a = Pop(stack);
    int b = Pop(stack);
    Push(stack, a >= b ? a : b);
}

void Min(struct Stack *stack){
    int a = Pop(stack);
    int b = Pop(stack);
    Push(stack, a <= b ? a : b);
}

void Neg(struct Stack *stack){
    Push(stack, -Pop(stack));
}

void Dup(struct Stack *stack){
    Push(stack, stack->data[stack->top - 1]);
}

void Swap(struct Stack *stack){
    int copyA = stack->data[stack->top - 1];
    stack->data[stack->top - 1] = stack->data[stack->top - 2];
    stack->data[stack->top - 2] = copyA;
}

int main(){
    struct Stack stack = InitStack(10);
    char command[20] = {0};
    while (strcmp(command, "END")){
        scanf("%s", command);
        if (!strcmp(command, "CONST")) Const(&stack);
        if (!strcmp(command, "ADD")) Add(&stack);
        if (!strcmp(command, "SUB")) Sub(&stack);
        if (!strcmp(command, "MUL")) Mul(&stack);
        if (!strcmp(command, "DIV")) Div(&stack);
        if (!strcmp(command, "MAX")) Max(&stack);
        if (!strcmp(command, "MIN")) Min(&stack);
        if (!strcmp(command, "NEG")) Neg(&stack);
        if (!strcmp(command, "DUP")) Dup(&stack);
        if (!strcmp(command, "SWAP")) Swap(&stack);
    }
    printf("%d", stack.data[stack.top - 1]);
    free(stack.data);
    stack.data = NULL;
    stack.cap = 0;
    stack.top = 0;
    return 0;
}