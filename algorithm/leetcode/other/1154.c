#include <stdio.h>
#include <string.h>


/*
    1. 拆分字符串，得到年、月、日
    2. 根据年份判断平年还是闰年，进而决定2月有几天
    3. 根据月份判断这个月有多少天
    4. 月日相加得到最终结果
*/
int main(){
    char* date = "2019-10-01";
    char* ymd = strtok(date,'-');
    printf("ndate=%s \n",ymd);

    // long ndate;
    // long temp;
    // long ten=10;
    // for (int i=0;i<strlen(date);i++){
    //     if (date[i]!='-'){
    //         temp = (date[i]-'0')*ten;
    //         ndate=ndate*ten+temp;
    //     }
    // }

    // printf("ndate=%lu \n",ndate);
}