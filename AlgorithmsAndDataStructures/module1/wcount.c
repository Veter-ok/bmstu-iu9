#include <stdio.h>

int wcount(char *s){
    int answer = 0;
    size_t i = 0;
    while(s[i] != '\0'){
        if (s[i] != ' ' && (s[i+1] == ' ' || s[i+1] == '\0')){
            answer++;
        }
        i++;
    }
    return answer;
}

int main(){
    char string[1000]; 
    gets(string);
    int a = wcount(string);
    printf("%d", a);
}