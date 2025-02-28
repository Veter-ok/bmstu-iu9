#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char *argv[]) {
    if(argc < 4) {
        printf("Usage: frame <height> <width> <text>");
        return 0;
    }
    unsigned int height = atoi(argv[1]), width = atoi(argv[2]);
    unsigned int size = strlen(argv[3]), space = (width - size) / 2;
    unsigned int primaryRow = ((height / 2) - 1 == 0 || (height % 2 == 1)) ? (height / 2) : (height / 2) - 1;
    char text[size];
    strcpy(text, argv[3]);
    if (width - 2 < size){
        printf("Error");
        return 0;
    }
    for (int i = 0; i < height; i++){
        printf("*");
        for (int j = 1; j < width - 1; j++){
            if (i == primaryRow){
                if (j < space|| j >= space + size){
                    printf(" ");
                }else{
                    printf("%c", text[j - space]);
                }
            }else if (i == 0 || i == height - 1){
                printf("*");
            }else{
                printf(" ");
            }
        } 
        printf("*\n");
    }

    return 0;
}