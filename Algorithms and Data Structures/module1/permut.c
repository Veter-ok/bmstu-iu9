#include <stdio.h>

int isEqual(long long int *a, long long int *b) {
    for (int i = 0; i < 8; i++){
        if (a[i] != b[i]){
            return 0;
        }
    }
    return 1;
}

void swap(long long int *x, long long int *y)
{
    long long int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}

int permutations(long long int *A, long long int *B, int fIndex, int lIndex)
{
   if (fIndex == lIndex){
        if (isEqual(A, B)){
            return 1;
        }
   } else {
       for (int i = fIndex; i <= lIndex; i++) {
            swap((A + fIndex), (A + i));
            int ans = permutations(A, B, fIndex + 1, lIndex);
            if (ans) return 1;
            swap((A + fIndex), (A + i)); 
       }
   }
   return 0;
}

int main(){
    long long int arrayA[8];
    long long int arrayB[8];

    for (int i = 0; i < 8; i++){
        scanf("%lld", &arrayA[i]);
    }
    for (int i = 0; i < 8; i++){
        scanf("%lld", &arrayB[i]);
    }

    int answer = permutations(arrayA, arrayB, 0, 7);
    if (answer) printf("yes");
    else printf("no");
}