#include <stdio.h>

unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i))
{
    if (nel == 1){
		return compare(0) == 0 ? 0 : 1;
	}
    long long int left = 0;
    long long int right = nel - 1;
    long long int mid = (left + right) / 2;

    while (left <= right){
        int isTarget = compare(mid);
        if (isTarget == 0){
            return mid;
        } else if (isTarget == 1){
            left = mid + 1;
        }else{
            right = mid - 1;
        }
        mid = (left + right) / 2;
    }
    return nel;
}