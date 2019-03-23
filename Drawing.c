#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <conio.h>
#include <windows.h>

#define xsize 110 
#define ysize 27 

typedef struct{   
char name[8];
int x ;
int y ;
}object;  //角色的参数. x,y为其坐标	

void Endless_Stripe(void);
void Bilibili_TV(void);
void Controllable_Move(void);
void C_M_paint(char s[],int x,int y);
void C_M_getmove(int *x ,int *y);
void Random_Move(void);
void R_M_reorder(object a[],int num);
void R_M_repaint(const object a[],const int num);
void R_M_random_move(int *x ,int *y);
void R_M_update_location(object a[],int num);

int main()
{
	srand(time(0));
	int a = 0;
	printf("1无限循环条纹, 2小电视, 3可控制移动的小人儿, 4一群随机运动的怪物\n");
	scanf("%d",&a);
	
	switch(a)
	{
		case 1:
			Endless_Stripe();break;
		case 2:
			Bilibili_TV();break;
		case 3:
			Controllable_Move();break;
		case 4:
			Random_Move();break;
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

void Random_Move(void)
{
	int num;//怪物数量 
	int i;//遍历用的 
	char name[8];
	
	printf("\n请输入怪物名称(越短越好): ");	//生成怪物
	scanf("%s",&name);
	printf("\n请输入怪物数量: ");
	scanf("%d",&num) ;
	object objects[num];
	
	for(i=0;i<num;++i){//初始化 
		strcpy(objects[i].name,name);//字符串是数组不能直接赋值 
		objects[i].x=rand()%xsize;
		objects[i].y=rand()%ysize;
	}
	
	system("cls"); 
	
	while(1){
		R_M_repaint(objects,num);	
		R_M_update_location(objects,num);
		Sleep(100);
		system("cls");	
	}
}

void R_M_reorder(object a[],int num){//从小到大排序，y优先 
	int i = 0;
	int j = 0;
	int min;
	object t;	

	for(i=0;i<num;++i){      					//比较排序 
		min=i;
		for(j=i;j<num;++j){
			if(a[min].y>a[j].y){
				min=j;
			}
			if(a[min].y==a[j].y&&a[min].x>a[j].x){
				min=j;
			}
		}
		if(min!=i){				//相当于换元素了，不过名字都一样没必要换 
			t.x=a[min].x;
			t.y=a[min].y;
			a[min].x=a[i].x;
			a[min].y=a[i].y;
			a[i].x=t.x;
			a[i].y=t.y;
		}
	}
}

void R_M_repaint(const object a[],const int num){
	int t = 0;
	int i = 0;
	int j = 0;
	
	for(i=0;i<a[0].y;++i){//打印a[0] 
		printf("\n");
	}
	for(j=0;j<a[0].x;++j){
		printf(" ");
	}
	printf("%s",a[0].name);
	
	for(t=1;t<num;++t){//根据坐标打印剩下的 //因为要判断i-1所以遍历从1开始 
		if(a[t-1].y<a[t].y){
			for(;i<a[t].y;++i){
				printf("\n");
			}
			for(j=0;j<a[t].x;++j){//不在同一行，从0开始打空格 
				printf(" ");
			}
			printf("%s",a[t].name);
		}else{			
			for(;i<a[t].y;++i){
				printf("\n");
			}
			for(;j<a[t].x;++j){//在同一行不从0打印空格，避免闪现,但还是因为字符串本身的长度而有些闪现 
				printf(" ");
			}
			printf("%s",a[t].name);
		}
	}
}

void R_M_random_move(int *x ,int *y){//随机向一个方向运动 
	int a =rand()%4;
	switch(a)
	{
		case 0 : ++(*x);break;
		case 1 : --(*x);break;
		case 2 : ++(*y);break;
		case 3 : --(*y);break;
	}
	
	if((*x)<0)(*x)=0;//防止越界 
	if((*y)<0)(*y)=0;
	if((*x)>xsize)(*x)=xsize;
	if((*y)>ysize)(*y)=ysize;
}

void R_M_update_location(object a[],int num){
	int i;
	for(i=0;i<num;++i){//全部随机运动一次 
		R_M_random_move(&a[i].x,&a[i].y);
	}
	
	R_M_reorder(a ,num );//重新排序 
}

