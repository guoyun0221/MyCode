#include<stdio.h>
#include<stdlib.h>
#include<string.h>

void Shutdown(void)
{
	printf("你的电脑将于3分钟后关机\n");
	printf("Your system will shut down in 2 minutes\n\n");
	system("shutdown -s -t 120");
	
	printf("输入'我是傻逼'可以阻止关机\n");
	printf("Input the words 'I-am-a-idiot' to stop this\n\n");
	char s[20] ;
	scanf("%s",&s);
	
	if((!strcmp(s,"我是傻逼"))||(!strcmp(s,"I-am-a-idiot"))){
		system("shutdown -a");
	}else{
		printf("\n享受这两分钟吧\n"); 
		printf("Enjoy your last two minutes\n\n");
		system("pause\n");
	}
}

int main()
{
	Shutdown();
	return 0 ;
}
