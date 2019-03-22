#include<stdio.h>
#include<windows.h>
#include<stdlib.h>
#include<time.h>
#include <conio.h>

#define xsize 112
#define ysize 28 

void Endless_Stripe(void);
void Bilibili_TV(void);
void Controllable_Move(void);
void C_M_paint(char s[],int x,int y);
void C_M_getmove(int *x ,int *y);

int main()
{
	srand(time(0));
	int a = 0;
	printf("1无限循环条纹, 2小电视, 3可控制移动的小人儿\n");
	scanf("%d",&a);
	
	switch(a)
	{
		case 1:
			Endless_Stripe();break;
		case 2:
			Bilibili_TV();break;
		case 3:
			Controllable_Move();break;
	}
	return 0 ;
}

void Endless_Stripe(void)
{
	while(1){

		printf("\t\t|\t        *       \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t       ***      \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t      *****     \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t     *******    \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t    *********   \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t   ***********  \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t  ************* \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t ***************\t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
						//the middle
		printf("\t\t \t  ************* \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t   ***********  \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t    *********   \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t     *******    \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t      *****     \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t       ***      \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);

	}
}

void Bilibili_TV(void)
{
	printf("\n");
		Sleep(70);
	printf("              *                   *\n");
		Sleep(70);
	printf("               *                 *\n");
		Sleep(70);
	printf("                *               *\n");
		Sleep(70);
	printf("          * * * * * * * * * * * * * * *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          *      *            *       *\n");
		Sleep(70);
	printf("          *    *               *      *\n");
		Sleep(70);
	printf("          *  *                  *     *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          *      *    **    *         *\n");
		Sleep(70);
	printf("          *       *  *  *  *          *\n");
		Sleep(70);
	printf("          *        **    **           *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          * * * * * * * * * * * * * * *\n");
		Sleep(70);
	printf("      \n  ");
		Sleep(70);	
}

void Controllable_Move(void)
{
	typedef struct{   
		char name[7];
		int x ;
		int y ;
	}object;  //角色的参数. x,y为其坐标
	
	object player;  //生成角色 
	printf("wasd控制移动\n");
	printf("请输入角色名称: ");
	scanf("%s",&player.name);
	player.x=xsize/2-3;
	player.y=ysize/3*2;
	
	system("cls"); 
	while(1)	//无限清屏重画
	{
		C_M_paint(player.name,player.x,player.y);
		C_M_getmove(&player.x,&player.y);
		Sleep(10);
		system("cls");
	}
	
}

void C_M_paint(char s[],int x,int y)
{			
	int i=0;
	for(i=0;i<y;++i){
		printf("\n");
	}
	for(i=0;i<x;++i){
		printf(" ");
	}
	printf("%s",s);
}

void C_M_getmove(int *x ,int *y)
{
	char a = getch();//wasd控制坐标变化 从而控制方向 
	
	if(a=='w') --(*y);
	if(a=='s') ++(*y);
	if(a=='a') (*x)-=2; //横向密集小幅度运动不明显 
	if(a=='d') (*x)+=2;
	
	if((*x)<0)(*x)=0;  //防止越界 
	if((*y)<0)(*y)=0;
	if((*x)>xsize)(*x)=xsize;
	if((*y)>ysize)(*y)=ysize;
}
