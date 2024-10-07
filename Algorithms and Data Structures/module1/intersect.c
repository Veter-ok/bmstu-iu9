#include <stdio.h>

int main(){
    unsigned int A = 0, B = 0;
    unsigned int sizeA, sizeB, elementA, elementB;

    scanf("%u", &sizeA);
    for (int i = 0; i < sizeA; i++){
        scanf("%u", &elementA);
        A |= (1 << elementA);
    }

    scanf("%u", &sizeB);
    for (int i = 0; i < sizeB; i++){
        scanf("%u", &elementB);
        B |= (1 << elementB);
    }
    
    unsigned int intersection = A & B;
    for (int i = 0; i < 32; i++) {
        if (intersection & (1 << i)) {
            printf("%u ", i);
        }
    }
}