#include <stdlib.h>
#include "elem.h"

struct Elem  *searchlist(struct Elem  *list, int k)
{
    if (list == NULL || (list->tag == INTEGER && list->value.i == k)){
        return list;
    }
    return searchlist(list->tail, k);
}