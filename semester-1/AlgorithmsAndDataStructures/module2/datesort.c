#include <stdio.h>

struct Date {
    int Day, Month, Year;
};


void sortDate(struct Date *dates, int length, struct Date *result){
    int countDays[31] = {0};
    int countMonths[12] = {0};
    int countYears[61] = {0};
    struct Date pre_res[length];
    for (int i = 0; i < length; i++){
        countDays[dates[i].Day - 1]++;
    }
    for (int i = 1; i < 31; i++){
        countDays[i] += countDays[i - 1];
    }
    int k = 0, i = 0;
    for (int j = length - 1; j >= 0; j--){
        k = dates[j].Day - 1;
        i = countDays[k] - 1;
        countDays[k] = i;
        result[i] = dates[j];
    }

    for (int i = 0; i < length; i++){
        countMonths[result[i].Month - 1]++;
    }
    for (int i = 1; i < 12; i++){
        countMonths[i] += countMonths[i - 1];
    }
    for (int j = length - 1; j >= 0; j--){
        k = result[j].Month - 1;
        i = countMonths[k] - 1;
        countMonths[k] = i;
        pre_res[i] = result[j];
    }

    for (int i = 0; i < length; i++){
        countYears[pre_res[i].Year - 1970]++;
    }
    for (int i = 1; i < 61; i++){
        countYears[i] += countYears[i - 1];
    }
    for (int j = length - 1; j >= 0; j--){
        k = pre_res[j].Year - 1970;
        i = countYears[k] - 1;
        countYears[k] = i;
        result[i] = pre_res[j];
    }
}

int main(){
    int n;
    scanf("%d", &n);
    struct Date dates[n];
    struct Date result[n];
    for (int i = 0; i < n; i++){
        scanf("%d %d %d", &dates[i].Year, &dates[i].Month, &dates[i].Day);
    }
    sortDate(dates, n, result);
    for (int i = 0; i < n; i++){
        printf("%04d %02d %02d\n", result[i].Year, result[i].Month, result[i].Day);
    }
    return 0;
}